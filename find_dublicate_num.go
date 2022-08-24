package main

import "fmt"

func main() {
	array := []int{1, 2, 2, 2, 4, 4, 5, 6}

	j := []int{}
	for i := 0; i < len(array)-1; i++ {
		if array[i] != array[i+1] {
			j = append(j, array[i])

		}
	}
	j = append(j, array[len(array)-1])
	fmt.Println(j)
}
