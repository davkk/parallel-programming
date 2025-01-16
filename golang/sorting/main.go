package main

import (
	"flag"
	"math/rand"
	"time"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

// https://www.slingacademy.com/article/writing-concurrent-sorting-algorithms-in-go/
// https://csit.am/2015/proceedings/PDC/PDCp1.pdf
func parallelMergeSort(arr []int, depth int) []int {
	if len(arr) <= 1 {
		return arr
	}

	if depth <= 0 {
		return mergeSort(arr)
	}

	mid := len(arr) / 2
	leftCh := make(chan []int)
	rightCh := make(chan []int)

	go func() { leftCh <- parallelMergeSort(arr[:mid], depth-1) }()
	go func() { rightCh <- parallelMergeSort(arr[mid:], depth-1) }()

	left := <-leftCh
	right := <-rightCh
	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	return merged
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size)
	}
	return arr
}

var sizeFlag = flag.Int("n", 1000, "size of the array")
var methodFlag = flag.String("m", "", "sorting method: 'sequential' or 'parallel'")

func main() {
	flag.Parse()
	size := *sizeFlag
	method := *methodFlag

	arr := generateRandomArray(size)

	var start time.Time

	switch method {
	case "sequential":
		start = time.Now()
		mergeSort(arr)
	case "parallel":
		start = time.Now()
		parallelMergeSort(arr, 5)
	default:
		panic("unknown sorting method")
	}

	elapsed := time.Since(start)
	println(size, elapsed.Nanoseconds())
}
