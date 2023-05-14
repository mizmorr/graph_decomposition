package core

import (
	"decomposition/maps"
	"fmt"
	"math"
	"sync"
)

var (
	someMapMutex = sync.RWMutex{}
)

type Geometric_node struct {
	x float64
	y float64
}
type Edge struct {
	source      *Geometric_node
	destination *Geometric_node
}
type Geometric_graph struct {
	nodes []*Geometric_node
	edges map[int64]*Edge
}

func GG_create(nodes []*Geometric_node, radius float64) *Geometric_graph {
	edges := make(map[int64]*Edge, 0)
	get_edges(nodes, []*Geometric_node{}, 0, 2, radius, &edges)
	return &Geometric_graph{
		nodes: nodes,
		edges: edges,
	}
}
func (g *Geometric_graph) Print() {
	for _, node := range g.edges {
		fmt.Println(node)
		fmt.Println(node.source.x, node.source.y, node.destination.x, node.destination.y)
	}
}
func get_edges(nodes []*Geometric_node, tmp []*Geometric_node, left, k int, radius float64, edges *map[int64]*Edge) {
	if k == 0 {
		if get_distance(tmp[0], tmp[1]) <= radius {

			if len(*edges) == 0 {
				(*edges)[0] = &Edge{source: tmp[0], destination: tmp[1]}
			} else {
				(*edges)[int64(len(*edges))] = &Edge{source: tmp[0], destination: tmp[1]}
			}
		}
		return
	}
	for i := left; i < len(nodes); i++ {
		tmp = append(tmp, nodes[i])
		get_edges(nodes, tmp, i+1, k-1, radius, edges)
		maps.Pop(&tmp)
	}
}
func get_distance(source, destination *Geometric_node) float64 {
	return math.Sqrt(math.Pow(destination.x-source.x, 2) + math.Pow(destination.y-source.y, 2))
}
func Get_node(x, y float64) *Geometric_node {
	return &Geometric_node{x: x, y: y}
}

func Test2(a []int, tmp []int, left, k int, result *map[int64]string) {
	if k == 0 {
		someMapMutex.Lock()
		if len(*result) == 0 {
			(*result)[0] = fmt.Sprint(tmp[0], tmp[1])
		} else {
			(*result)[int64(len(*result))] = fmt.Sprint(tmp[0], tmp[1])
		}
		someMapMutex.Unlock()
		return
	}
	for i := left; i < len(a); i++ {
		tmp = append(tmp, a[i])
		Test2(a, tmp, i+1, k-1, result)
		maps.Pop(&tmp)
	}
}
