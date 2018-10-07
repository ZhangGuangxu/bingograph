package main

import "fmt"

func main() {
	a := make([][]int, 25)
	fmt.Println(len(a))
	fmt.Println(len(a[1]))

	//test1()
	//test2()
	//test3()
	test4()
}

func test4() {
	g := NewGraph()
	g.AddNode(0)
	g.AddNode(23)
	g.AddNode(24)

	g.AddNode(5)
	g.AddNode(10)
	g.AddNode(15)
	g.AddNode(20)
	g.AddNode(21)
	g.AddNode(22)

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(9)
	g.AddNode(14)

	g.AddNode(7)
	g.AddNode(12)
	g.AddNode(17)
	g.AddNode(18)

	var s PathSearcher = NewAstar(g, 0, 24)
	fmt.Println(s.Search())

	if g.HasNode(19) {
		s = NewAstar(g, 0, 24)
		if s.Search() == nil {
			fmt.Println("path found 1")
			return
		}
	}
	if g.HasNode(18) || g.HasNode(22) {
		s = NewAstar(g, 0, 23)
		if s.Search() == nil {
			fmt.Println("path found 2")
			return
		}
	}
	fmt.Println("no path")
}

func test3() {
	g := NewGraph()
	g.AddNode(0)
	g.AddNode(24)

	g.AddNode(5)
	g.AddNode(10)
	g.AddNode(15)
	g.AddNode(20)
	g.AddNode(21)
	g.AddNode(22)
	g.AddNode(23)

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(9)
	g.AddNode(14)

	var s PathSearcher = NewAstar(g, 0, 24)
	fmt.Println(s.Search())
}

func test2() {
	g := NewGraph()
	g.AddNode(0)
	g.AddNode(23)
	g.AddNode(24)
	var s PathSearcher = NewAstar(g, 0, 24)
	fmt.Println(s.Search())
	g.AddNode(6)
	g.AddNode(12)
	g.AddNode(18)
	s = NewAstar(g, 0, 24)
	fmt.Println(s.Search())
	s = NewAstar(g, 0, 23)
	fmt.Println(s.Search())
}

func test1() {
	g := NewGraph()
	g.AddNode(0)
	g.AddNode(23)
	g.AddNode(24)
	g.AddNode(24)
	var s PathSearcher = NewAstar(g, 0, 24)
	fmt.Println(s.Search())
	g.AddNode(5)
	g.AddNode(10)
	g.AddNode(15)
	g.AddNode(20)
	g.AddNode(21)
	g.AddNode(22)
	s = NewAstar(g, 0, 23)
	fmt.Println(s.Search())
	s = NewAstar(g, 0, 24)
	fmt.Println(s.Search())
}
