package main

import "fmt"

var a = [5]int{10, 5, 2, 4, 7}

func selectionSort() {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d, %d, %d, %d, %d\n", a[0], a[1], a[2], a[3], a[4])
		index, value := i, a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] < value {
				index = j
				value = a[j]
			}
		}
		a[index] = a[i]
		a[i] = value
	}
}

func main() {
	selectionSort()
}
