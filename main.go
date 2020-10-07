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
