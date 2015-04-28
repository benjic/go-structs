// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import "testing"

func TestJohnsonAllShortestPath(t *testing.T) {

	g := New()
	w := g.AddVertex("W")
	x := g.AddVertex("X")
	y := g.AddVertex("Y")
	z := g.AddVertex("Z")

	g.AddEdge(w, z, 2)
	g.AddEdge(x, y, 3)
	g.AddEdge(x, w, 6)
	g.AddEdge(y, w, 4)
	g.AddEdge(y, z, 5)
	g.AddEdge(z, x, -7)
	g.AddEdge(z, y, -3)

	_, err := g.JohnsonsAllShortestPath()

	if err != nil {
		t.Error(err)
	}

}

func TestJohnsonsNegativeWeightCycle(t *testing.T) {

	g := New()
	a := g.AddVertex("A")
	b := g.AddVertex("B")
	c := g.AddVertex("C")
	d := g.AddVertex("D")
	e := g.AddVertex("E")

	g.AddEdge(a, b, -1)
	g.AddEdge(a, c, 4)
	g.AddEdge(b, c, 3)
	g.AddEdge(b, d, -2)
	g.AddEdge(b, e, 2)
	g.AddEdge(d, b, -1)
	g.AddEdge(d, c, 5)
	g.AddEdge(e, d, -3)

	_, err := g.JohnsonsAllShortestPath()

	if err == nil {
		t.Error("Expected error for negative weight cycle in graph, failed")
	}
}
