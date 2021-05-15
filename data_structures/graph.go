package datastructures

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

// Shortest Path problem
// given a weighted graph, find the shortest path of edges from node A to node B
// algorithms: BFS (unweighted graph, Dijkstra's, Bellman-Ford, Floyd-Warshall, etc)
