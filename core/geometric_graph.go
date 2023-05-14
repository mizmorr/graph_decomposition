package core

import (
	"decomposition/maps"
	"fmt"
	"math"
	"sync"
)

var (
	someMapMutex  = sync.RWMutex{}
	someMapMutex2 = sync.RWMutex{}
)

type Geometric_node struct {
	x  float64
	y  float64
	ID int64
}
type Edge struct {
	source      *Geometric_node
	destination *Geometric_node
}
type Triangle struct {
	source      *Geometric_node
	middle      *Geometric_node
	destination *Geometric_node
}
type Geometric_graph struct {
	nodes     []*Geometric_node
	edges     map[int64]*Edge
	triangles map[int64]*Triangle
}

func GG_create(nodes []*Geometric_node, radius float64) *Geometric_graph {
	edges := get_edges(nodes, radius)
	return &Geometric_graph{
		nodes: nodes,
		edges: edges,
	}
}
func (g *Geometric_graph) Test3() {
	fmt.Println(g.edges)
}
func (g *Geometric_graph) Print() {
	for _, edge := range g.edges {
		fmt.Println(edge.source.ID, edge.destination.ID)
	}
}

func (g *Geometric_graph) get_triangles() {
	for i := int64(0); i < int64(len(g.edges)-1); i++ {
		for j := i + 1; j < int64(len(g.edges)); j++ {
			if ok, node_1, node_2, node_3 := are_connected(g.edges[i], g.edges[j]); ok {
				if g.search(node_1, node_2) {
					if len(g.triangles) == 0 {

						g.triangles[0] = &Triangle{source: node_1, middle: node_2, destination: node_3}

					} else {
						g.triangles[int64(len(g.triangles))] = &Triangle{source: node_1, middle: node_2, destination: node_3}
					}
				}
			}
		}
	}
}
func are_connected(first, second *Edge) (bool, *Geometric_node, *Geometric_node, *Geometric_node) {
	switch {
	case first.source == second.source:
		return true, first.destination, second.destination, first.source
	case first.source == second.destination:
		return true, first.destination, second.source, first.source
	case first.destination == second.source:
		return true, first.source, second.destination, first.destination
	case first.destination == second.destination:
		return true, first.source, second.source, first.destination
	default:
		return false, nil, nil, nil
	}
}
func (g *Geometric_graph) search(s, d *Geometric_node) bool {
	var wg sync.WaitGroup
	edges := maps.Values2(g.edges)
	is_searched := false
	wg.Add(4)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[:len(g.edges)/4]
		for _, edge := range part {
			if edge.source == s && edge.destination == d || edge.source == d && edge.destination == s {
				is_searched = true
				return
			}
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[len(edges)/4 : 2*len(edges)/4]
		for _, edge := range part {
			if edge.source == s && edge.destination == d || edge.source == d && edge.destination == s {
				is_searched = true
				return
			}
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[2*len(edges)/4 : 3*len(edges)/4]
		for _, edge := range part {
			if edge.source == s && edge.destination == d || edge.source == d && edge.destination == s {
				is_searched = true
				return
			}
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[3*len(edges)/4:]
		for _, edge := range part {
			if edge.source == s && edge.destination == d || edge.source == d && edge.destination == s {
				is_searched = true
				return
			}
		}
	}(&wg)
	wg.Wait()
	return is_searched
}
func map_write(nodes []*Geometric_node, tmp []*Geometric_node, left, k int, radius float64, edges *map[int64]*Edge) {
	if k == 0 {
		if get_distance(tmp[0], tmp[1]) <= radius {
			someMapMutex.Lock()

			if len(*edges) == 0 {

				(*edges)[0] = &Edge{source: tmp[0], destination: tmp[1]}

			} else {
				(*edges)[int64(len(*edges))] = &Edge{source: tmp[0], destination: tmp[1]}
			}
			someMapMutex.Unlock()
		}
		return
	}
	for i := left; i < len(nodes); i++ {
		tmp = append(tmp, nodes[i])
		map_write(nodes, tmp, i+1, k-1, radius, edges)
		maps.Pop(&tmp)
	}
}
func get_edges(nodes []*Geometric_node, radius float64) map[int64]*Edge {
	m := make(map[int64]*Edge, 0)
	var wg sync.WaitGroup
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		first_half := nodes[:len(nodes)/2]
		map_write(first_half, []*Geometric_node{}, 0, 2, radius, &m)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		second_half := nodes[len(nodes)/2:]
		map_write(second_half, []*Geometric_node{}, 0, 2, radius, &m)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		first_half := nodes[:len(nodes)/2]
		second_half := nodes[len(nodes)/2:]
		for _, node := range first_half {
			for _, node2 := range second_half {
				if get_distance(node, node2) <= radius {
					someMapMutex.Lock()
					if len(m) == 0 {
						m[0] = &Edge{source: node, destination: node2}

					} else {

						m[int64(len(m))] = &Edge{source: node, destination: node2}
					}
					someMapMutex.Unlock()
				}
			}
		}
	}(&wg)
	fmt.Println("waiting")
	wg.Wait()
	return m
}
func get_distance(source, destination *Geometric_node) float64 {
	return math.Sqrt(math.Pow(destination.x-source.x, 2) + math.Pow(destination.y-source.y, 2))
}
func Get_node(x, y float64, id int64) *Geometric_node {
	return &Geometric_node{x: x, y: y, ID: id}
}

func Test2(a []int, tmp []int, left, k int, result *map[int64]string) {
	if k == 0 {
		someMapMutex2.Lock()
		if len(*result) == 0 {
			(*result)[0] = fmt.Sprint(tmp[0], tmp[1])
		} else {
			(*result)[int64(len(*result))] = fmt.Sprint(tmp[0], tmp[1])
		}
		someMapMutex2.Unlock()
		return
	}
	for i := left; i < len(a); i++ {
		tmp = append(tmp, a[i])
		Test2(a, tmp, i+1, k-1, result)
		maps.Pop(&tmp)
	}
}

func Testing(a []int) {
	var wg sync.WaitGroup
	wg.Add(3)
	m := make(map[int64]string, 0)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		Test2(a[:len(a)/2], []int{}, 0, 2, &m)

	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		Test2(a[len(a)/2:], []int{}, 0, 2, &m)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		first_half := a[:len(a)/2]
		second_half := a[len(a)/2:]
		for _, elem := range first_half {
			for _, e := range second_half {
				someMapMutex2.Lock()
				if len(m) == 0 {
					m[0] = fmt.Sprint(e, elem)
				} else {
					m[int64(len(m))] = fmt.Sprint(e, elem)
				}
				someMapMutex2.Unlock()
			}
		}
	}(&wg)
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println(m)
}
