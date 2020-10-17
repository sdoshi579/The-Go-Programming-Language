package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func main() {
	p := point{
		x: 1,
		y: 2,
	}
	q := point{4, 6}
	fmt.Printf("Distance %f \n", distance(p, q))
	fmt.Printf("Distance method %f \n", p.distance(q))
}

func distance(p, q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p point) distance(q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
