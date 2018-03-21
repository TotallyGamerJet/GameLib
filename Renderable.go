package main

type Renderable struct {
	vertices []float32
	indices []int32
	textureCoords []float32
}

func createRenderable(vertices []float32, indices []int32, textureCoords []float32) Renderable {
	return Renderable{vertices, indices, textureCoords}
}
