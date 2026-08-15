package benchmarking

import "fmt"

func AddToList() {
	val := 10000
	list := make([]int, val)
	for v := range val {
		list[v] = v
	}
	fmt.Println("LIST: ", list)
}

func AddToListx() {
	var list []int
	for i := range 10000 {
		list = append(list, i)
	}
	fmt.Println("LIST: ", list)
}
