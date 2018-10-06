package main

type Graph struct {
    nodes []*GraphNode
    edges map[int][]*GraphEdge
}

func NewGraph() *Graph {
    return &Graph{
        nodes: make([]*GraphNode, 0),
        edges: make(map[int][]*GraphEdge),
    }
}

func (g *Graph) AddNode(i int) {
    if g.hasNode(i) {
        return
    }
    g.nodes = append(g.nodes, NewGraphNode(i))

    ns := GetNeighbours(i)
    for _, n := range ns {
        if g.hasNode(n) {
            g.addEdge(i, n)
            g.addEdge(n, i)
        }
    }
}

func (g *Graph) hasNode(i int) bool {
    for _, n := range g.nodes {
        if n.Index == i {
            return true
        }
    }
    return false
}

func (g *Graph) NodeCount() int {
    return len(g.nodes)
}

func (g *Graph) addEdge(f, t int) {
    es, ok := g.edges[f]
    if !ok {
        g.edges[f] = append(make([]*GraphEdge, 0, 1), NewGraphEdge(f, t))
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