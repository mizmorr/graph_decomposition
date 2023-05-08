package main

import (
	"bufio"
	"decomposition/maps"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

type Graph interface {
	add_Edge()
	k_core_label()
	append_node()
	insert_node()
	remove_node(id int64)
	get_new_id() int64
}
type Node struct {
	ID          int64
	adj_list    map[int64]int64
	edge_number int64
}
type Undirected_Graph struct {
	nodes map[int64]*Node
}
type Core_set struct {
	numbers map[int64]int64
}

func (g *Undirected_Graph) add_Edge(source, dest int64) {
	g.nodes[source].adj_list[source] = dest
	g.nodes[dest].adj_list[dest] = source
}
func pop_front(s []*Node) (int64, []*Node, error) {
	if len(s) == 0 {
		return -1, nil, fmt.Errorf("bad")
	}
	return s[0].ID, s[1:], nil
}
func pop(s []int64) (int64, []int64, error) {
	if len(s) == 0 {
		return -1, nil, fmt.Errorf("bad")
	}
	return s[0], s[1:], nil
}
func (g *Undirected_Graph) get_new_id() int64 {
	if len(g.nodes) == 0 {
		return 0
	}
	return g.nodes[int64(len(g.nodes)-1)].ID + 1
}
func p_pring(ar []*Node) {
	for _, n := range ar {
		fmt.Printf("%v ", n.ID)
	}
	fmt.Println()
}
func test() {
	g := first_sample()
	var (
		unprocessed = maps.Keys_ordered(g.nodes)
		v           int64
		err         error
	)
	for len(g.nodes) > 0 {
		v, unprocessed, err = pop(unprocessed)
		g.print()
		fmt.Println(g.nodes[v].adj_list)
		if err == nil {
			fmt.Println(v)
			g.remove_node(v)
		}
		fmt.Println()
	}
}
func (g *Undirected_Graph) k_core_label() *Core_set {
	core_set := Core_set{numbers: make(map[int64]int64, len(g.nodes))}
	var (
		k, v int64
		err  error
	)
	for len(g.nodes) > 0 {
		k++
		unprocessed := maps.Keys_ordered(g.nodes)
		for len(unprocessed) > 0 {
			v, unprocessed, err = pop(unprocessed)
			if err == nil && g.nodes[v] != nil {
				if g.nodes[v].edge_number < k {
					for _, node := range g.nodes[v].adj_list {
						if g.nodes[node] != nil {
							g.nodes[node].edge_number--
							unprocessed = append(unprocessed, g.nodes[node].ID)
						}
					}
					core_set.numbers[v] = k - 1
					g.remove_node(v)
				}
			}
		}

	}
	return &core_set
}

func first_sample() *Undirected_Graph {
	var result []int64
	g := NewUndirectedGraph()
	f, err := ioutil.ReadFile("samples/graph1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		for _, j := range s {
			postj, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}
			result = append(result, int64(postj))
		}
	}
	adjes := map[int64]map[int64]int64{}
	for _, node := range result {
		if _, value := adjes[node]; !value {
			adjes[node] = make(map[int64]int64, 0)
		}
	}
	for i := int64(0); i < int64(len(result)-1); i += 2 {
		adjes[result[i]][int64(len(adjes[result[i]]))] = result[i+1]
		adjes[result[i+1]][int64(len(adjes[result[i+1]]))] = result[i]
	}
	for i := int64(0); i < int64(len(adjes)); i++ {
		g.append_node(adjes[i])
	}
	for _, node := range g.nodes {
		node.edge_number = int64(len(node.adj_list))
	}
	// for i := 0; i < len(k); i++ {
	// }
	return g
}

func (g *Undirected_Graph) remove_node(id int64) {
	// vals := maps.Values2(g.nodes)
	// node_id := interpolation_search(vals, 0, int64(len(vals)-1), id)

	// if g.nodes[node_id] == nil {
	// 	return
	// }
	for _, from := range g.nodes[id].adj_list {
		if g.nodes[from] != nil {
			adj := maps.Values2(g.nodes[from].adj_list)
			id_adj := interpolation_search(adj, 0, int64(len(adj)-1), id)
			delete(g.nodes[from].adj_list, id_adj)
		}
	}
	delete(g.nodes, id)

}

// func (g *Undirected_Graph) remove_node(id int64) {

//		g.nodes = a[len(a)-1] // Copy last element to index i.
//		a[len(a)-1] = ""      // Erase last element (write zero value).
//		a = a[:len(a)-1]
//		fmt.Print64ln(a)
//	}

func (g *Undirected_Graph) append_node(ids map[int64]int64) {
	g.nodes[g.get_new_id()] = &Node{g.get_new_id(), ids, int64(len(ids))}
}

// func (g *Undirected_Graph) append_node(ids []int64) {
// 	adj := func(ar []int64) (nodes []*Node) {

//			for _, id := range ar {
//				nodes = append(nodes, g.get_node(id))
//			}
//			return
//		}(ids)
//		g.nodes = append(g.nodes, &Node{g.get_id(), 0, adj})
//	}
func NewUndirectedGraph() *Undirected_Graph {
	return &Undirected_Graph{nodes: make(map[int64]*Node, 0)}
}

// func (g *Undirected_Graph) RemoveNode(id int6464) {
// 	if _, ok := g.nodes[id]; !ok {
// 		return
// 	}
// 	delete(g.nodes, id)

// 	for from := range g.edges[id] {
// 		delete(g.edges[from], id)
// 	}
// 	delete(g.edges, id)

//		g.nodeIDs.Release(id)
//	}
func (g *Undirected_Graph) print() {
	for _, node := range g.nodes {
		fmt.Printf("id - %v, adj - %v, edge num - %v\n", node.ID, node.adj_list, node.edge_number)
	}
}

// func interpolation_search(arr []*Node, low, high, search int64) int64 {

// 	if low <= high && search >= arr[low].ID && search <= arr[high].ID {

// 		if arr[high].ID-arr[low].ID == 0 {
// 			switch {
// 			case arr[len(arr)-1].ID == search:
// 				return int64(len(arr) - 1)
// 			default:
// 				return -1
// 			}
// 		}
// 		pos := low + (((high - low) / (arr[high].ID - arr[low].ID)) * (search - arr[low].ID))
// 		switch {
// 		case arr[pos].ID == search:
// 			return search
// 		case arr[pos].ID < search:
// 			return interpolation_search(arr, pos+1, high, search)
// 		case arr[pos].ID > search:
// 			return interpolation_search(arr, low, pos-1, search)
// 		}
// 	}
// 	return -1
// }

func interpolation_search(arr []int64, low, high, search int64) int64 {

	if low <= high && search >= arr[low] && search <= arr[high] {

		if arr[high]-arr[low] == 0 {
			switch {
			case arr[len(arr)-1] == search:
				return int64(len(arr) - 1)
			default:
				return -1
			}
		}
		pos := low + (((high - low) / (arr[high] - arr[low])) * (search - arr[low]))
		switch {
		case arr[pos] == search:
			return pos
		case arr[pos] < search:
			return interpolation_search(arr, pos+1, high, search)
		case arr[pos] > search:
			return interpolation_search(arr, low, pos-1, search)
		}
	}
	return -1
}

// func Create_Graph(size int64) Graph {
// 	nodes := []Node{}
// 	i := 0
// 	for i < size {
// 		nodes = append(nodes, Node{fmt.Sprint64(i + 1), 0, []int64{}})
// 		i++
// 	}
// 	return Graph{nodes}
// }

// func Create_Graph_with_names(names []string) {
// 	defer wg.Done()
// 	nodes := []Node{}
// 	go func() {
// 		for _, name := range names {
// 			nodes = append(nodes, Node{name, 0, []int64{}})
// 			fmt.Print64ln(name)
// 		}
// 	}()
// 	return
// }

func main() {
	// nodes := []*Node{{0, 0, map[int64]int64{0: 1, 1: 3}, 2}, {1, 0, map[int64]int64{0: 0, 1: 3}, 2}, {3, 0, map[int64]int64{0: 0, 1: 1}, 2}}
	g := first_sample()
	fmt.Println(g.k_core_label())
	// set := g.k_core_label()
	// fmt.Println(set)
	// var elements = []int{1, 2, 3, 4}

	// fmt.Println(graph.nodes)
	// graph.remove_node(8)
	// fmt.Println(graph.nodes)

	// graph.remove_node(1)
	// graph.remove_node(2)
	// graph.print()
	// graph.remove_node(8)
	// graph.print()

	// core_set := graph.k_core_label()
	// core_set.numbers[1] = 1
	// fmt.Println(binarySearch2([]int{0, 1, 2, 3}, 1))
	// 	fmt.Printf("%v, ", graph.nodes[0].ID)
	// 	fmt.Printf("%v, ", graph.nodes[1].ID)
	// 	fmt.Printf("%v, ", graph.nodes[2].ID)
	// 	fmt.Printf("%v, \n", graph.nodes[3].ID)

	// 	for i := range graph.nodes {
	// 		fmt.Printf("%v ", i)
	// 		fmt.Printf("%v, ", graph.nodes[i].ID)
	// 	}
	// 	fmt.Println()
	// 	fmt.Println("_-------------------------------------_")
	// 	fmt.Println()
	// 	// graph.print()
	// }

	// for {
	// 	graph := NewUndirectedGraph()
	// 	graph.append_node(map[int64]int64{})
	// 	graph.append_node(map[int64]int64{0: 0})
	// 	graph.append_node(map[int64]int64{0: 0, 1: 1})
	// 	graph.append_node(map[int64]int64{0: 1, 1: 2})
	// 	vals := maps.Values2(graph.nodes)
	// 	for _, val := range vals {
	// 		fmt.Printf("%v, ", val.ID)
	// 	}
	// 	fmt.Println()
	// }

	// graph.append_node([]int64{0, 2})
	// graph.append_node([]int64{0, 2})

	// node := binarySearch(graph.nodes, 1)
	// fmt.Print64ln(node.ID)

}

//	func (g Graph) createGraph(names []string) {
//		for i := range names {
//			g.adj=append
//		}
//	}

// func (g Graph) createGEdg(edge []Edge) {

// }

// func main() {
// 	// graph := Create_Graph(4)
// 	// fmt.Print64ln(graph)
// 	wg.Add(1)
// 	Create_Graph_with_names([]string{"a", "b", "c", "d"})
// 	wg.Wait()
// }
