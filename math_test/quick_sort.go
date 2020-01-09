package main

import (
	"fmt"
)

const MAX = 10

var sortArray = []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}

func main() {

	fmt.Println("before sort：")
	show(false)
	fmt.Println()

	quickSort(sortArray, 0, MAX-1)

	fmt.Println("after sort:")
	show(false)

}

// quickSort
func quickSort(sortArray []int, left, right int) {

	if left < right {
		key := sortArray[left]
		i := left
		j := right

		for {
			for i+1 < MAX {
				i++
				if key <= sortArray[i] {
					break
				}
			}

			for j-1 >= 0 {
				if key >= sortArray[j] {
					break
				}
				j--
			}

			if i >= j {
				break
			}
			swap(i, j)
			show(true)
		}

		sortArray[left] = sortArray[j]
		sortArray[j] = key
		show(false)

		quickSort(sortArray, left, j-1)
		quickSort(sortArray, j+1, right)
	}

}

// Swap the position of a and b
func swap(a, b int) {
	sortArray[a], sortArray[b] = sortArray[b], sortArray[a]
}

// foreach
func show(isMiddle bool) {
	if isMiddle {
		fmt.Print("mid:")
	}
	for _, value := range sortArray {
		fmt.Printf("%d\t", value)
	}
	fmt.Println()
}
