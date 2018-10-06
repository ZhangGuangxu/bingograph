package main

type GraphNode struct {
    Index int
}

func NewGraphNode(i int) *GraphNode {
    return &GraphNode{
        Index: i,
    }
}