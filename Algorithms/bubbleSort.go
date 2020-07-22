package main

import "fmt"

var arr = []int{1, 3, 4, 8, 6, 7, 2, 0}

func bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				var temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	fmt.Println(array)
}

func main() {
	bubbleSort(arr)
}
