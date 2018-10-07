package main

import (
	"errors"
)

var ErrNoPath = errors.New("no path")

// Astar algorithm
type Astar struct {
	graph  *Graph
	source int
	target int
	hFn    func(nd1, nd2 int) int // heuristic

	frontier map[int]*GraphEdge // search frontier
	gcost    map[int]int        // cost to some node
	fcost    map[int]int        // cost to target. fcost = gcost + hcost (heuristic)
	spt      map[int]*GraphEdge // shortest path tree
}

func NewAstar(g *Graph, s, t int) *Astar {
	return &Astar{
		graph:    g,
		source:   s,
		target:   t,
		hFn:      ManhattanDistance,
		frontier: make(map[int]*GraphEdge),
		fcost:    make(map[int]int),
		gcost:    make(map[int]int),
		spt:      make(map[int]*GraphEdge),
	}
}

// Search trys to find the shortest path from source to target.
func (d *Astar) Search() error {
	src, dst := d.source, d.target
	d.frontier[src] = NewGraphEdge(src, src)
	d.gcost[src] = 0
	d.fcost[src] = 0
	pq := NewIndexedPriorityQueueMinWithNWayAndSize(d.fcost, WayCount(), d.graph.NodeCount())
	pq.Insert(src)
	fn := func(e *GraphEdge, g int, t int) {
		d.frontier[t] = e
		d.gcost[t] = g
		d.fcost[t] = g + d.hFn(t, dst)
	}

	for !pq.IsEmpty() {
		idx, err := pq.Pop()
		if err != nil {
			break
		}

		edge := d.frontier[idx]
		d.spt[idx] = edge
		i := edge.To
		//edge.Show() // for debug

		if i == dst {
			return nil
		}

		edges := d.graph.Edges()
		for _, e := range edges[i] {
			t := e.To
			g := d.gcost[i] + 1
			if _, ok := d.frontier[t]; !ok {
				fn(e, g, t)
				pq.Insert(t)
			} else if g < d.gcost[t] {
				if _, ok := d.spt[t]; !ok {
					fn(e, g, t)
					pq.ChangePriority(t)
				}
			}
		}
	}

	return ErrNoPath
}
