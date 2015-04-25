// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

// Graph abstraction and associated algorithms
package graph

type Graph struct {
	Vertices map[string]*Vertex
	Edges    []*Edge
}

type Edge struct {
	Tail, Head *Vertex
	Weight     int
}

type Vertex struct {
	Label    string
	incoming []*Edge
	outgoing []*Edge
}

// Instantiates a new graph primitive
func New() Graph {
	g := Graph{
		make(map[string]*Vertex, 0),
		make([]*Edge, 0),
	}

	return g
}

func (g *Graph) AddVertex(label string) {
	vert := &Vertex{label, make([]*Edge, 0), make([]*Edge, 0)}
	if _, ok := g.Vertices[vert.Label]; !ok {
		g.Vertices[vert.Label] = vert
	}

	return g.Vertices[label]
}

func (g *Graph) AddEdge(tail, head *Vertex, weight int) {
	e := &Edge{tail, head, weight}
	tail.outgoing = append(tail.outgoing, e)
	head.incoming = append(head.incoming, e)

	g.Edges = append(g.Edges, e)
}
