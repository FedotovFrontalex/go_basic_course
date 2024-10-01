package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	//	var userHeight, userWeight float64 = 1.8, 100
	userHeight := 1.8
	//var userWeight float64 = 100
	userWeight := 100.0
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
