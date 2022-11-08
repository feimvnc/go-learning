// quick sort, like divide and conquer methon

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piece := createPieces(15)
	fmt.Println(piece)
	quickSort(piece)
	fmt.Println(piece)
}

func createPieces(size int) []int {
	piece := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		piece[i] = rand.Intn(999) - rand.Intn(999)
	}
	return piece
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	center := rand.Int() % len(a)
	a[center], a[right] = a[right], a[center]
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println(quickSortStart([]int{5, 6, 7, 2, 1.0}))
// }

// func partition(arr []int, low, high int) ([]int, int) {
// 	pivot := arr[high]
// 	i := low
// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
// 	arr[i], arr[high] = arr[high], arr[i]
// 	return arr, i
// }

// func quickSort(arr []int, low, high int) []int {
// 	if low < high {
// 		var p int
// 		arr, p = partition(arr, low, high)
// 		arr = quickSort(arr, low, p-1)
// 		arr = quickSort(arr, p+1, high)
// 	}
// 	return arr
// }

// func quickSortStart(arr []int) []int {
// 	return quickSort(arr, 0, len(arr)-1)
// }
