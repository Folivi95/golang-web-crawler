package graph

type Graph struct {
	Adjacency map[string][]string
}

// NewGraph creates new graph data structure
func NewGraph() *Graph {
	return &Graph{
		Adjacency: make(map[string][]string),
	}
}

// AddVertex add vertex to the graph
func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.Adjacency[vertex]; ok {
		return false
	}
	g.Adjacency[vertex] = []string{}
	return true
}

// AddEdge adds node to an existing vertex
// or creates a vertex using the node if vertex exists
// but node does not exist in graph
func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.Adjacency[vertex]; !ok {
		return false
	}
	if ok := contains(g.Adjacency[vertex], node); ok {
		return false
	}

	if _, ok := g.Adjacency[node]; !ok {
		g.AddVertex(node)
	}

	g.Adjacency[vertex] = append(g.Adjacency[vertex], node)
	return true
}

func contains(edges []string, node string) bool {
	set := make(map[string]struct{}, len(edges))
	for _, n := range edges {
		set[n] = struct{}{}
	}
	_, ok := set[node]
	return ok
}
