package graph

type Graph struct {
	Adjacency map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		Adjacency: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.Adjacency[vertex]; ok {
		return false
	}
	g.Adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.Adjacency[vertex]; !ok {
		// fmt.Printf("vertex %s does not exists! \n", vertex)
		return false
	}
	if ok := contains(g.Adjacency[vertex], node); ok {
		// fmt.Printf("node %s already exists! \n", node)
		return false
	}

	if _, ok := g.Adjacency[node]; !ok {
		// fmt.Printf("Node %s not found in Adjacency, creating new vertex\n", node)
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
