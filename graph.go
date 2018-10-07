package main

type Graph struct {
	nodes map[int]bool
	edges map[int][]*GraphEdge
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]bool),
		edges: make(map[int][]*GraphEdge),
	}
}

func (g *Graph) AddNode(i int) {
	if g.HasNode(i) {
		return
	}
	g.nodes[i] = true

	ns := GetNeighbours(i)
	for _, n := range ns {
		if g.HasNode(n) {
			g.addEdge(i, n)
			g.addEdge(n, i)
		}
	}
}

func (g *Graph) HasNode(i int) bool {
    _, ok := g.nodes[i]
    return ok
}

func (g *Graph) NodeCount() int {
	return len(g.nodes)
}

func (g *Graph) addEdge(f, t int) {
	es, ok := g.edges[f]
	if !ok {
		g.edges[f] = []*GraphEdge{ NewGraphEdge(f, t) }
	} else {
		for _, e := range es {
			if e.From == f && e.To == t {
				return
			}
		}
		g.edges[f] = append(es, NewGraphEdge(f, t))
	}
}

func (g *Graph) Edges() map[int][]*GraphEdge {
	return g.edges
}
