package main

import "fmt"

func duplicateInArray(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] && i != j {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 2, 2}))
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3}))
}
