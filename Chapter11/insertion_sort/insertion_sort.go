package main

import "fmt"

var a = [5]int{10, 5, 2, 4, 7}

func insertionSort() {
	for j := 1; j < len(a); j++ {
		fmt.Printf("%d, %d, %d, %d, %d\n", a[0], a[1], a[2], a[3], a[4])
		key, i := a[j], j-1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	fmt.Printf("%d, %d, %d, %d, %d\n", a[0], a[1], a[2], a[3], a[4])
}

func main() {
	insertionSort()
}
