package main

import (
	"fmt"
	"image/color"
	"math"
)

type point struct {
	x, y float64
}

type coloredPoint struct {
	point
	color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := coloredPoint{point{1, 1}, red}
	q := coloredPoint{point{5, 4}, blue}
	fmt.Println(p.distance(q.point))
	p.scaleBy(2)
	q.scaleBy(2)
	fmt.Println(p)
	fmt.Println(p.distance(q.point))
}

func (p point) distance(q point) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func (p *point) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}
