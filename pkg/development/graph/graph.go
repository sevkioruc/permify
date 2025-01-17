package graph

import (
	"sync"
)

// Node - Structure
type Node struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Edge - Edge Structure
type Edge struct {
	Extra any   `json:"extra"`
	From  *Node `json:"from"`
	To    *Node `json:"to"`
}

// Graph - Graph Structure
type Graph struct {
	nodes []*Node
	edges []*Edge
	lock  sync.RWMutex
}

// Nodes - Return Nodes Slice
func (g *Graph) Nodes() []*Node {
	return g.nodes
}

// Edges - Return Edge Slice
func (g *Graph) Edges() []*Edge {
	return g.edges
}

// AddNodes - Add nodes to graph
func (g *Graph) AddNodes(n []*Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n...)
	g.lock.Unlock()
}

// AddNode - Add node to graph
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdges - Add edges to graph
func (g *Graph) AddEdges(e []*Edge) {
	g.lock.Lock()
	g.edges = append(g.edges, e...)
	g.lock.Unlock()
}

// AddEdge - Add edge to graph
func (g *Graph) AddEdge(from, to *Node, extra any) {
	g.lock.Lock()
	g.edges = append(g.edges, &Edge{
		Extra: extra,
		From:  from,
		To:    to,
	})
	g.lock.Unlock()
}
