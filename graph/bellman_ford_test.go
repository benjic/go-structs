// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import "testing"

// Compare result of bellman ford with expected result
func TestBellmanFordShortestPath(t *testing.T) {

	g := New()
	a := g.AddVertex("A")
	b := g.AddVertex("B")
	c := g.AddVertex("C")
	d := g.AddVertex("D")
	e := g.AddVertex("E")

	g.AddEdge(a, b, -1)
	g.AddEdge(a, c, 4)
	g.AddEdge(b, c, 3)
	g.AddEdge(b, d, 2)
	g.AddEdge(b, e, 2)
	g.AddEdge(d, b, 1)
	g.AddEdge(d, c, 5)
	g.AddEdge(e, d, -3)

	want := map[string]int{"A": 0, "B": -1, "C": 2, "D": -2, "E": 1}
	got, err := g.BellmanFordShortestPath(a)

	if err != nil {
		t.Error(err)
	}

	for vert, dist := range got {
		if want[vert] != dist {
			t.Errorf("Expected minimum distance between \"A\" and %s to be %d, got %d", vert, want[vert], dist)
		}
	}
}

// Confirms the return of an error of a negative cycle when given a graph with a
// negative cycle
func TestBellmanFordNegativeWeightCycle(t *testing.T) {

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

	_, err := g.BellmanFordShortestPath(a)

	if err == nil {
		t.Error("Expected error for negative weight cycle in graph, failed")
	}
}
