package datastructures

// Singly linked list node
type ListNodeS struct {
	Val  int
	Next *ListNodeS
}

type ListNodeD struct {
	Val  int
	Prev *ListNodeD
	Next *ListNodeD
}
