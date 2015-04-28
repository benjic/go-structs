// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import (
	"errors"
	"math"
)

func (g *Graph) FloydWarshallAllShortestPath() (map[string]map[string]int, error) {

	dist := make(map[string]map[string]int, len(g.Vertices))

	for _, vert := range g.Vertices {
		dist[vert.Label] = make(map[string]int, len(g.Vertices))

		for _, v := range g.Vertices {
			dist[vert.Label][v.Label] = int(math.MaxInt32)
		}

		dist[vert.Label][vert.Label] = 0
		for _, edge := range vert.outgoing {
			dist[vert.Label][edge.Head.Label] = edge.Weight
		}
	}

	for k := range g.Vertices {
		for j := range g.Vertices {
			for i := range g.Vertices {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	for vert := range g.Vertices {
		if dist[vert][vert] != 0 {
			return nil, errors.New("Graph contians a negitve weight cycle")
		}
	}

	return dist, nil
}
