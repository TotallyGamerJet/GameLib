package main

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

const(
	vertexShader = `
		#version 450 core

		layout (location = 0) in vec3 aPos;

		void main()
		{
			gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
		}
		` + "\x00"

	fragmentShader =	`
		#version 450 core
		out vec4 FragColor;

		void main()
		{
			 FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
		}
		` + "\x00"
)

var(
	vertices = []float32 {
		0.5,  0.5, 0.0,  // top right
		0.5, -0.5, 0.0,  // bottom right
		-0.5, -0.5, 0.0,  // bottom left
		-0.5,  0.5, 0.0,   // top left
	}
	indices = []int {
		0, 1, 3,
		1, 2, 3,
	}
	textureCoords = []float32 {
		1.0, 1.0,
		1.0, 0.0,
		0.0, 0.0,
		0.0, 1.0,
	}
)

func main() {
	display := createDisplay("Game Library Test", 800, 600)
	defer display.end()

	loader := createLoader()
	defer loader.cleanUp()

	obj := loader.loadToVAO(vertices, indices)

	//shader := createShaderProgram(vertexShader, fragmentShader)
	shader := createShaderProgram("./shaders/vertex.shader", "./shaders/fragment.shader")

	for !display.ShouldClose() {
		display.clear()

		//render
		render(obj, shader)

		display.update()
	}
}
