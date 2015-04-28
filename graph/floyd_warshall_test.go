// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import "testing"

func TestFloydWarshallAllShortestPath(t *testing.T) {

	// The graph given on wikipedia

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

	want := map[string]map[string]int{
		"W": map[string]int{
			"W": 0,
			"X": -5,
			"Y": -2,
			"Z": 2,
		},
		"X": map[string]int{
			"W": 6,
			"X": 0,
			"Y": 3,
			"Z": 8,
		},
		"Y": map[string]int{
			"W": 4,
			"X": -2,
			"Y": 0,
			"Z": 5,
		},
		"Z": map[string]int{
			"W": -1,
			"X": -7,
			"Y": -4,
			"Z": 0,
		},
	}
	got, err := g.FloydWarshallAllShortestPath()

	if err != nil {
		t.Error(err)
	}

	for source, sinks := range want {
		for sink, weight := range sinks {
			if got[source][sink] != weight {
				t.Errorf("Expected shortest path between %s and %s to be %d, got %d", source, sink, weight, got[source][sink])
			}
		}
	}
}

func TestFloydWarshallAllShortestFlatPath(t *testing.T) {

	g := New()
	a := g.AddVertex("A")
	b := g.AddVertex("B")
	c := g.AddVertex("C")
	d := g.AddVertex("D")
	e := g.AddVertex("E")
	f := g.AddVertex("F")

	g.AddEdge(a, b, -5)
	g.AddEdge(a, f, -10)
	g.AddEdge(b, c, 1)
	g.AddEdge(c, d, 1)
	g.AddEdge(d, e, -10000)

	want := map[string]map[string]int{
		"A": map[string]int{
			"A": 0,
			"B": -5,
			"C": -4,
			"D": -3,
			"E": -10003,
			"F": -10,
		},
		"B": map[string]int{
			"B": 0,
			"C": 1,
			"D": 2,
			"E": -9998,
		},
		"C": map[string]int{
			"C": 0,
			"D": 1,
			"E": -9999,
		},
		"D": map[string]int{
			"D": 0,
			"E": -10000,
		},
		"E": map[string]int{
			"E": 0,
		},
		"F": map[string]int{
			"F": 0,
		},
	}
	got, err := g.FloydWarshallAllShortestPath()

	if err != nil {
		t.Error(err)
	}

	for source, sinks := range want {
		for sink, weight := range sinks {
			if got[source][sink] != weight {
				t.Errorf("Expected shortest path between %s and %s to be %d, got %d", source, sink, weight, got[source][sink])
			}
		}
	}
}

func TestFloydWarshallNegativeWeightCycle(t *testing.T) {

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

	_, err := g.FloydWarshallAllShortestPath()

	if err == nil {
		t.Error("Expected error for negative weight cycle in graph, failed")
	}
}
