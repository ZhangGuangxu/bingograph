package main

func WayCount() int {
    return 4
}

func NodeIdx2Coord(i int) (x, y int) {
    x = i / 5
    y = i % 5
    return
}

func Coord2NodeIdx(x, y int) int {
    return x*5 + y
}

func ManhattanDistance(n1, n2 int) int {
    x1, y1 := NodeIdx2Coord(n1)
    x2, y2 := NodeIdx2Coord(n2)
    x := x1 - x2
    if x < 0 {
        x = -x
    }
    y := y1 - y2
    if y < 0 {
        y = -y
    }
    return x + y
}

func IsCoordValid(c int) bool {
    return c >= 0 && c <= 4
}

func GetNeighbours(i int) []int {
    result := make([]int, 0)
    x, y := NodeIdx2Coord(i)
    x1 := x - 1
    if IsCoordValid(x1) {
        result = append(result, Coord2NodeIdx(x1, y))
    }
    x2 := x + 1
    if IsCoordValid(x2) {
        result = append(result, Coord2NodeIdx(x2, y))
    }
    y1 := y - 1
    if IsCoordValid(y1) {
        result = append(result, Coord2NodeIdx(x, y1))
    }
    y2 := y + 1
    if IsCoordValid(y2) {
        result = append(result, Coord2NodeIdx(x, y2))
    }
    return result
}
