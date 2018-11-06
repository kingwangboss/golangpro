package main

//
import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	var sum float64
	for i := 0; i < path.LenPoints()-1; i++ {
		sum = [i]path + [i+1]path
	}
	return 0
}

func (path Path) LenPoints() int {
	return len(path)
}

func main() {
	path := Path{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(path.Distance())
}
