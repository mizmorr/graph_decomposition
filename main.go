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
func testing(a []int) {
	var wg sync.WaitGroup
	wg.Add(3)
	m := make(map[int64]string, 0)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		core.Test2(a[:len(a)/2], []int{}, 0, 2, &m)

	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		core.Test2(a[len(a)/2:], []int{}, 0, 2, &m)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		first_half := a[:len(a)/2]
		second_half := a[len(a)/2:]
		for _, elem := range first_half {
			for _, e := range second_half {
				appender(&m, fmt.Sprint(elem, e))

			}
		}
	}(&wg)
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	// first := core.Get_node(0, 0)
	// second := core.Get_node(1, 1)
	// third := core.Get_node(2, 2)
	// nodes := []*core.Geometric_node{first, second, third}
	// ggraph := core.GG_create(nodes, 2)
	// ggraph.Print()
	// testing()
	// t := time.Now()
	// testing(a)
	// res_conc := time.Since(t).Seconds()
	// m := make(map[int64]string, 0)
	// t = time.Now()
	// core.Test2(a, []int{}, 0, 2, &m)
	// res_linear := time.Since(t).Seconds()
	// fmt.Println("res_conc", res_conc, "res_linear")
	another_test(a)

}
