package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{4, 5, 6, 7, 8, 0, 1, 2, 3}
	f := sort.Search(len(arr), func(i int) bool { //1)闭包arr
		return arr[i] <= arr[0]
	})
	fmt.Println("arr:", arr)
	fmt.Println("found:", f, "arr[f]:", arr[f])
}
