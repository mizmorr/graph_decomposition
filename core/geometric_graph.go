package core

import "math"

type Geometric_node struct {
	x float64
	y float64
}

func Get_distance(source, destination *Geometric_node) float64 {
	return math.Sqrt(math.Pow(destination.x-source.x, 2) + math.Pow(destination.y-source.y, 2))
}
func Get_node(x, y float64) *Geometric_node {
	return &Geometric_node{x: x, y: y}
}
