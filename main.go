package main

import (
	"decomposition/core"
	"fmt"
)

func main() {

	first := core.Get_node(0, 0)
	second := core.Get_node(1, 1)
	fmt.Println(core.Get_distance(first, second))
}
