package main

import (
	"decomposition/core"
	"fmt"
	"sync"
)

func another_test(a []int) {
	var someMapMutex = sync.RWMutex{}
	m := make(map[int64]int, 0)
	var wg sync.WaitGroup
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *map[int64]int) {
		defer wg.Done()
		part := a[:len(a)/2]
		for _, v := range part {
			someMapMutex.Lock()
			if len(*m) == 0 {
				(*m)[0] = v
			} else {
				(*m)[int64(len(*m))] = v
			}
			someMapMutex.Unlock()
		}

	}(&wg, &m)
	go func(wg *sync.WaitGroup, m *map[int64]int) {
		defer wg.Done()
		part := a[len(a)/2:]
		for _, v := range part {
			someMapMutex.Lock()
			if len(*m) == 0 {
				(*m)[0] = v
			} else {
				(*m)[int64(len(*m))] = v
			}
			someMapMutex.Unlock()
		}

	}(&wg, &m)
	go func(wg *sync.WaitGroup, m *map[int64]int) {
		defer wg.Done()
		part := a
		for _, v := range part {
			someMapMutex.Lock()
			if len(*m) == 0 {
				(*m)[0] = v
			} else {
				(*m)[int64(len(*m))] = v
			}
			someMapMutex.Unlock()
		}

	}(&wg, &m)
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()

}
func appender(m *map[int64]string, str string) {
	var someMapMutex = sync.RWMutex{}
	someMapMutex.Lock()
	if len(*m) == 0 {
		(*m)[0] = str
	} else {
		(*m)[int64(len(*m))] = str
	}
	someMapMutex.Unlock()

}

func main() {
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	first := core.Get_node(0, 0, 0)
	second := core.Get_node(1, 1, 1)
	third := core.Get_node(2, 2, 2)
	fourth := core.Get_node(3, 3, 3)
	fifth := core.Get_node(4, 4, 4)
	sixth := core.Get_node(5, 5, 5)
	seventh := core.Get_node(6, 6, 6)
	eighth := core.Get_node(7, 7, 7)
	nodes := []*core.Geometric_node{first, second, third, fourth, fifth, sixth, seventh, eighth}

	graph := core.GG_create(nodes, 3)
	core.Test7(graph)
	// fmt.Println(ggraph.Search(0, 1))

	// testing()
	// t := time.Now()
	// testing(a)
	// res_conc := time.Since(t).Seconds()
	// m := make(map[int64]string, 0)
	// t = time.Now()
	// core.Test2(a, []int{}, 0, 2, &m)
	// res_linear := time.Since(t).Seconds()
	// fmt.Println("res_conc", res_conc, "res_linear")
	// another_test(a)

}
