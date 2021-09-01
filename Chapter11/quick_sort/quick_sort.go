package main

import "fmt"

var a = [8]int{5, 2, 4, 7, 1, 3, 2, 6}

func partitionLeft(start int, end int) int {
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

func partitionRight(start int, end int) int {
	pivot := start
	for start < end {
		for a[pivot] <= a[end] && start < end {
			end--
		}
		for a[start] <= a[pivot] && start < end {
			start++
		}
		if start < end {
			a[start], a[end] = a[end], a[start]
		}
	}
	a[pivot], a[end] = a[end], a[pivot]
	return end
}

func quickSort(start int, end int) {
	if start < end {
		mid := partitionRight(start, end)
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
