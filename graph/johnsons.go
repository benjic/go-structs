// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

// Johnsons algorithm is an all shortest path algorithm which provides weights
// of shortest paths from all vertices within the graph. The function returns a
// map vertex labels as keys and their sink maps keyed by sink labels. If the
// graph contains a negative cycle the function will return a nil result with an
// appropriate error.
func (g *Graph) JohnsonsAllShortestPath() (map[string]map[string]int, error) {

	paths := make(map[string]map[string]int, len(g.Vertices))
	q := g.AddVertex("q")

	for _, vert := range g.Vertices {
		paths[vert.Label] = make(map[string]int, len(g.Vertices))
		g.AddEdge(q, vert, 0)
	}

	// Execute Bellman-Ford with q as source
	d, err := g.BellmanFordShortestPath(q)

	delete(d, "q")
	delete(g.Vertices, "q")
	delete(paths, "q")
	g.Edges = g.Edges[0 : len(g.Edges)-len(g.Vertices)-1]
	// TODO: Leaves edges from q in all vertices incoming arc maps

	// If bellman-ford
	if err != nil {
		return nil, err
	}

	for _, edge := range g.Edges {
		edge.Weight = edge.Weight + d[edge.Tail.Label] - d[edge.Head.Label]
	}

	for _, source := range g.Vertices {
		for sink, weight := range g.DijkstrasShortestPath(source) {
			paths[source.Label][sink] = weight - d[source.Label] + d[sink]
		}
	}

	return paths, nil
}
