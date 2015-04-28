// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import "testing"

func TestDijkstrasShortestPath(t *testing.T) {

	g := New()
	a := g.AddVertex("A")
	b := g.AddVertex("B")
	c := g.AddVertex("C")
	d := g.AddVertex("D")
	e := g.AddVertex("E")

	g.AddEdge(a, b, 1)
	g.AddEdge(a, c, 4)
	g.AddEdge(b, c, 3)
	g.AddEdge(b, d, 2)
	g.AddEdge(b, e, 6)
	g.AddEdge(d, b, 1)
	g.AddEdge(d, c, 5)
	g.AddEdge(e, d, 3)

	cases := []struct {
		sink *Vertex
		want int
	}{
		{a, 0},
		{b, 1},
		{c, 4},
		{d, 3},
		{e, 7},
	}

	got := g.DijkstrasShortestPath(a)
	for _, c := range cases {
		if got[c.sink.Label] != c.want {
			t.Errorf("Expected path between a and %s to have weight %d, got %d", c.sink.Label, c.want, got[c.sink.Label])
		}
	}

}
