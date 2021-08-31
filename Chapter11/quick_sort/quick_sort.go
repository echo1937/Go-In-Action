package main

import "fmt"

var a = [8]int{5, 2, 4, 7, 1, 3, 2, 6}

func partition(start int, end int) int {
	pivot := start
	index := pivot + 1

	for i := index; i <= end; i++ {
		if a[i] < a[pivot] {
			a[i], a[index] = a[index], a[i]
			index++
		}
	}
	a[pivot], a[index-1] = a[index-1], a[pivot]
	return index - 1
}

func quickSort(start int, end int) {
	if start < end {
		mid := partition(start, end)
		fmt.Printf("partition (%d-%d, mid=%d) %d %d %d %d %d %d %d %d\n",
			start, end, mid,
			a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
		quickSort(start, mid-1)
		quickSort(mid+1, end)
		fmt.Printf("quickSort (%d-%d, %d-%d) %d %d %d %d %d %d %d %d\n",
			start, mid-1, mid+1, end,
			a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
	}
}

func main() {
	quickSort(0, len(a)-1)
}
