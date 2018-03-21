package main

import (
	"strings"
	"fmt"
	"github.com/go-gl/gl/v4.5-core/gl"
	"io/ioutil"
)

type ShaderProgram struct {
	programID, vertexID, fragmentID uint32
}

func createShaderProgram(vertexFile, fragmentFile string) ShaderProgram {
	vertexID, err := loadShader(vertexFile, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentID, err := loadShader(fragmentFile, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	programID := gl.CreateProgram()
	gl.AttachShader(programID, vertexID)
	gl.AttachShader(programID, fragmentID)
	gl.LinkProgram(programID)

	gl.DeleteShader(vertexID)
	gl.DeleteShader(fragmentID)

	return ShaderProgram{programID, vertexID, fragmentID}
}

func (s ShaderProgram) start() {
	gl.UseProgram(s.programID)
}

func (s ShaderProgram) stop() {
	gl.UseProgram(0)
}

func loadShader(filename string, shaderType uint32) (uint32, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(string(file) + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0,fmt.Errorf("failed to compile %v: %v", file, log)
	}

	return shader, nil
}