package ports

type GraphStructure interface {
	AddVertex(vertex string) bool
	AddEdge(vertex, node string) bool
}
