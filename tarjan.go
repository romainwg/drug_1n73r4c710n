package main

type Vertex struct {
	// Tarjan variables
	index   int
	lowlink int
	onstack bool
}

type Edge struct {
	vertexIn  *Vertex
	vertexOut *Vertex
}

func tarjan() {

}
