package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.5-core/gl"
	"log"
	"fmt"
)


type Display struct {
	*glfw.Window
	width, height int32
}


func createDisplay(title string, width, height int32) Display {

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(resize)
	//window.SetCursorPosCallback(mouseCallback)
	//window.SetScrollCallback(scrollCallback)

	//window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	gl.Viewport(0, 0, width, height)

	return Display{window, width, height}
}

func resize(window *glfw.Window, width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func(d *Display) clear() {
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (d *Display) update() {
	d.SwapBuffers()
	glfw.PollEvents()
}

func (d *Display) end() {
	d.Destroy()
	glfw.Terminate()
}