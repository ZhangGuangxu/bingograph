package main

import (
    "fmt"
)

type GraphEdge struct {
    From int // node index
    To int // node index
}

func NewGraphEdge(f, t int) *GraphEdge {
    return &GraphEdge{
        From: f,
        To: t,
    }
}

// for debug
func (e *GraphEdge) Show() {
    fmt.Printf("e: From=%d, To=%d\n", e.From, e.To)
}