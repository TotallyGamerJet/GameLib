package main

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

func render(model RawModel, program ShaderProgram) {
	program.start()
	defer program.stop()
	gl.BindVertexArray(model.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.DrawElements(gl.TRIANGLES, model.vertexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)
	gl.BindVertexArray(0)
}