package main

import "fmt"

var a = [8]int{5, 2, 4, 7, 1, 3, 2, 6}

func merge(start int, mid int, end int) {
	// 在sort函数中, 数组被分为[start, mid]和[mid+1, end]两个部分,
	// n1为left的length, 长度mid-start +1
	// n2为right的length, 长度为end-mid
	n1, n2 := mid-start+1, end-mid
	left, right := make([]int, n1), make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = a[start+i]
	}
	for i := 0; i < n2; i++ {
		right[i] = a[mid+1+i]
	}

	// i为left下标, j为right下标, k为原数组下标
	i, j, k := 0, 0, start
	for i < n1 && j < n2 {
		if left[i] < right[j] {
			a[k] = left[i]
			k++
			i++
		} else {
			a[k] = right[j]
			k++
			j++
		}
	}

	for i < n1 {
		a[k] = left[i]
		k++
		i++
	}
	for j < n2 {
		a[k] = right[j]
		k++
		j++
	}
}

func sort(start int, end int) {
	if start < end {
		mid := (start + end) / 2
		fmt.Printf("sort (%d-%d, %d-%d) %d %d %d %d %d %d %d %d\n",
			start, mid, mid+1, end,
			a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
		sort(start, mid)
		sort(mid+1, end)
		merge(start, mid, end)
		fmt.Printf("merge (%d-%d, %d-%d) to %d %d %d %d %d %d %d %d\n",
			start, mid, mid+1, end,
			a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
	}
}

func main() {
	sort(0, len(a)-1)
}
