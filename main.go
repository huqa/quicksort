package main

import (
	"fmt"
	"math/rand"
)

// stretching my legs
func main() {
	list := rand.Perm(1000)
	fmt.Println("unsorted slice of 1000", list)
	quicksort(&list, 0, len(list)-1)
	fmt.Println("sorted slice of 1000", list)

	list = rand.Perm(1000)
	fmt.Println("unsorted slice of 1000", list)
	quicksortHoare(&list, 0, len(list)-1)
	fmt.Println("hoare sorted slice of 1000", list)
}

// quicksort from Introduction to Algorithms
func quicksort(A *[]int, p int, r int) {
	if p < r {
		q := partition(A, p, r)
		quicksort(A, p, q-1)
		quicksort(A, q+1, r)
	}
}

// partition rearranges the subarray A in place
func partition(A *[]int, p int, r int) int {
	pivot := (*A)[r]
	i := p
	for j := p; j < r; j++ {
		if (*A)[j] <= pivot {
			(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
			i = i + 1
		}
	}
	(*A)[i], (*A)[r] = (*A)[r], (*A)[i]
	return i
}

// quicksort Hoare partition version
func quicksortHoare(A *[]int, p int, r int) {
	if p < r {
		q := hoarePartition(A, p, r)
		quicksortHoare(A, p, q)
		quicksortHoare(A, q+1, r)
	}
}

// hoare partition
func hoarePartition(A *[]int, p int, r int) int {
	pivot := (*A)[p]
	i := p - 1
	j := r + 1
	for {
		for ok := true; ok; ok = (*A)[i] < pivot {
			i = i + 1
		}
		for ok := true; ok; ok = (*A)[j] > pivot {
			j = j - 1
		}
		if i < j {
			(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		} else {
			return j
		}
	}
}
