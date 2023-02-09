package graph

import "testing"

func TestGraph_AddVertex(t *testing.T) {
	// given an empty graph
	graph := NewGraph()
	baseUrl := "https://example.com"

	// add URL to graph
	graph.AddVertex(baseUrl)

	t.Run("adding vertex that exist in graph", func(t *testing.T) {
		// when
		isAdded := graph.AddVertex(baseUrl)

		// then should return false
		if isAdded {
			t.Error("Expected false, got ", isAdded)
		}
	})

	t.Run("adding vertex that does not exist in graph", func(t *testing.T) {
		// when
		newBaseUrl := "https://joshuatest.com"
		isAdded := graph.AddVertex(newBaseUrl)

		// then should return true
		if !isAdded {
			t.Error("Expected true, got ", isAdded)
		}
	})
}

func TestGraph_AddEdge(t *testing.T) {
	// given an empty graph
	graph := NewGraph()
	baseUrl := "https://example.com"
	node1 := "https://example.com/test"

	// add URL to graph
	graph.AddVertex(baseUrl)

	t.Run("adding node to vertex that does not exist", func(t *testing.T) {
		// when
		newBaseUrl := "https://test.com"
		isAdded := graph.AddEdge(newBaseUrl, node1)

		// then should return false
		if isAdded {
			t.Error("Expected false, got ", isAdded)
		}
	})

	t.Run("adding node to vertex in graph", func(t *testing.T) {
		// when
		isAdded := graph.AddEdge(baseUrl, node1)

		// then should return true
		if !isAdded {
			t.Error("Expected true, got ", isAdded)
		}
	})

	t.Run("adding node that already belongs to vertex", func(t *testing.T) {
		// when
		isAdded := graph.AddEdge(baseUrl, node1)

		// then should return false
		if isAdded {
			t.Error("Expected false, got ", isAdded)
		}
	})
}
