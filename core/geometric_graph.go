package core

import (
	"decomposition/maps"
	"fmt"
	"math"
	"reflect"
	"sort"
	"sync"
)

var (
	someMapMutex  = sync.RWMutex{}
	someMapMutex2 = sync.RWMutex{}
)

type Geometric_node struct {
	x            float64
	y            float64
	ID           int64
	edges_number int64
	core_number  int
}
type Edge struct {
	source          *Geometric_node
	destination     *Geometric_node
	triangle_number int64
	triangles       map[int64]*Corner
}
type Corner struct {
	middle      *Edge
	destination *Edge
}
type Geometric_graph struct {
	nodes []*Geometric_node
	edges map[int64]*Edge
}

func GG_create(nodes []*Geometric_node, radius float64) *Geometric_graph {
	edges := get_edges(nodes, radius)

	graph := &Geometric_graph{
		nodes: nodes,
		edges: edges,
	}
	graph.get_triangles()
	graph.mark_nodes()
	return graph
}
func (g *Geometric_graph) Test3() {
	fmt.Println(g.edges)
}
func (g *Geometric_graph) Print() {
	// for i, edge := range g.edges {
	// 	fmt.Println(i, edge.source.ID, edge.destination.ID)
	// }
	for i := 0; i < len(g.edges); i++ {
		fmt.Println(g.edges[int64(i)].destination.ID, g.edges[int64(i)].source.ID, g.edges[int64(i)].triangle_number)
		for _, t := range g.edges[int64(i)].triangles {
			fmt.Println(t.destination.destination.ID, t.destination.source.ID, t.middle.destination.ID, t.middle.source.ID)
		}
		fmt.Println()
	}
	fmt.Println()
	// for _, node := range g.nodes {
	// 	fmt.Println(node.ID, node.edges_number)
	// }
}

func (e *Edge) add_triangle(middle, destination *Edge) {
	if len(e.triangles) == 0 {
		e.triangles[0] = &Corner{middle: middle, destination: destination}
		e.triangle_number++
	} else {
		corner := &Corner{middle: middle, destination: destination}
		if !e.is_triangle_exist(corner) {
			e.triangles[int64(len(e.triangles))] = corner
			e.triangle_number++
		}
	}
}
func (g *Geometric_graph) get_triangles() {
	for i := int64(0); i < int64(len(g.edges)-1); i++ {
		for j := i + 1; j < int64(len(g.edges)); j++ {
			if ok, node_1, node_2 := are_connected(g.edges[i], g.edges[j]); ok {
				if ok, num := g.Search(node_1, node_2); ok {
					g.edges[i].add_triangle(g.edges[j], g.edges[num])
					g.edges[j].add_triangle(g.edges[i], g.edges[num])
					g.edges[num].add_triangle(g.edges[i], g.edges[j])
				}
			}
		}
	}

}
func (g *Geometric_graph) remove_edge(id int64) {
	delete(g.edges, id)
}
func (g *Geometric_graph) triangle_k_core() {

}
func Test4(g *Geometric_graph) {
	ma := maps.Values2(g.edges)
	for i := int64(0); i < int64(len(ma)); i++ {
		fmt.Println(ma[i].destination.ID, ma[i].source.ID)
	}
	fmt.Println()
	for i := int64(0); i < int64(len(g.edges)); i++ {
		fmt.Println(g.edges[i].destination.ID, g.edges[i].source.ID)
	}
}
func (t *Corner) equal(other *Corner) bool {

	f_ids := []int64{t.destination.destination.ID, t.destination.source.ID, t.middle.destination.ID, t.middle.source.ID}
	s_ids := []int64{other.destination.destination.ID, other.destination.source.ID, other.middle.destination.ID, other.middle.source.ID}
	sort.Slice(f_ids, func(i, j int) bool {
		return f_ids[i] < f_ids[j]
	})
	sort.Slice(s_ids, func(i, j int) bool {
		return s_ids[i] < s_ids[j]
	})
	return reflect.DeepEqual(f_ids, s_ids)
}
func (e *Edge) is_triangle_exist(t *Corner) bool {

	for _, elem := range e.triangles {
		if elem.equal(t) {
			return true
		}
	}
	return false
}
func are_connected(first, second *Edge) (bool, *Geometric_node, *Geometric_node) {
	switch {
	case first.source == second.source:
		return true, first.destination, second.destination
	case first.source == second.destination:
		return true, first.destination, second.source
	case first.destination == second.source:
		return true, first.source, second.destination
	case first.destination == second.destination:
		return true, first.source, second.source
	default:
		return false, nil, nil
	}
}
func (g *Geometric_graph) Search(s, d *Geometric_node) (bool, int64) {
	var wg sync.WaitGroup
	edges := maps.Values2(g.edges)
	is_searched := false
	var search_num int64 = -1
	wg.Add(4)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[:len(g.edges)/4]
		for i := 0; i < len(part); i++ {
			if part[i].source == s && part[i].destination == d || part[i].source == d && part[i].destination == s {
				is_searched = true
				search_num = int64(i)
				return
			}
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[len(edges)/4 : 2*len(edges)/4]
		for i := 0; i < len(part); i++ {
			if part[i].source == s && part[i].destination == d || part[i].source == d && part[i].destination == s {
				is_searched = true
				search_num = int64(i + len(edges)/4)

				return
			}
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[2*len(edges)/4 : 3*len(edges)/4]
		for i := 0; i < len(part); i++ {
			if part[i].source == s && part[i].destination == d || part[i].source == d && part[i].destination == s {
				is_searched = true
				search_num = int64(i + 2*len(edges)/4)
				return
			}
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		part := edges[3*len(edges)/4:]
		for i := 0; i < len(part); i++ {
			if part[i].source == s && part[i].destination == d || part[i].source == d && part[i].destination == s {
				is_searched = true
				search_num = int64(i + 3*len(edges)/4)
				return
			}
		}
	}(&wg)
	wg.Wait()
	return is_searched, search_num
}
func (g *Geometric_graph) mark_nodes() {
	for _, edge := range g.edges {
		g.nodes[edge.source.ID].edges_number++
		g.nodes[edge.destination.ID].edges_number++
	}
}
func map_write(nodes []*Geometric_node, tmp []*Geometric_node, left, k int, radius float64, edges *map[int64]*Edge) {
	if k == 0 {
		if get_distance(tmp[0], tmp[1]) <= radius {
			someMapMutex.Lock()

			if len(*edges) == 0 {

				(*edges)[0] = &Edge{source: tmp[0], destination: tmp[1], triangles: map[int64]*Corner{}}

			} else {
				(*edges)[int64(len(*edges))] = &Edge{source: tmp[0], destination: tmp[1], triangles: map[int64]*Corner{}}
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
						m[0] = &Edge{source: node, destination: node2, triangles: map[int64]*Corner{}}

					} else {

						m[int64(len(m))] = &Edge{source: node, destination: node2, triangles: map[int64]*Corner{}}
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
