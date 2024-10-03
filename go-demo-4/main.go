package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(arr *[4]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		temp := arr[j]
		arr[j] = arr[i]
		arr[i] = temp
	}
}
