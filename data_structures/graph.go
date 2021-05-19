package datastructures

import "fmt"

// From Algorithms Course - Graph Theory Tutorial from a Google Engineer (https://www.youtube.com/watch?v=09_LlHjoEiY)

// An adjacency matrix is a very simple way to represent a graph.
// The idea is that the cell m[i][j] represents the edge weight
// of going from node i to node j.

// Pros
// space efficicient for representing dense graphs
// Edge weight lookup is O(1)
// Simplest graph representation

// Cons
// Requires O(V^2) space
// Iterating over all edges takes O(V^2) time

// Main alternative to the matrix is adjacency list
// Pros
// Space efficient for representing sparse graphs
// Iterating over all edges is efficient

// Cons
// Less space efficient for denser graphs
// Edge weight lookup is O(E)
// Slightly more complex graph representation

// Adjacency list based Graph implementation
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graph if it's not already included
func (g *Graph) AddVertex(k int) {
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds an edge to the directed graph
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v --> %v", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)

		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
	fmt.Println()
}

// Edge list is a way to represent a graph simple as an unordered list of edges.
// Assume the notation for any triplet (u, v, w) means:
// "the cost from node u to node v is w"
// Seldomly used as it lacks structure

// e.g. [(C,A,4), (A,C,1), (B, C, 6), (A, B, 4), ...]

// Ask yourself these question:
// 1) Is the graph directed or undirected?
// 2) Are the edges of the graph weighted?
// 3) Is the graph I will encounter likely to be sparse or dense with edges?
// 4) Should I use an adjanceny matrix, adjacency list, edge list, or other structure
// 	  to represent the graph efficiently?

// Shortest Path problem:
// given a weighted graph, find the shortest path of edges from node A to node B
// algorithms: BFS (unweighted graph, Dijkstra's, Bellman-Ford, Floyd-Warshall, etc)
