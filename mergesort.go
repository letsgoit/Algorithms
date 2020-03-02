package main

import "fmt"

func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}
	for i < len(a) {
		r[i+j] = a[i]

	}
	for j < len(b) {
		r[i+j] = b[j]

	}
	return r
}
func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	var middle = len(items) / 2
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	return merge(a, b)
}
func main() {

	fmt.Print(MergeSort([]int{4, 5, 6, 2, 9, 11, 60}))
	fmt.Println("vim-go")
}
