package main

import (
	"fmt"
	"math"
)

func main() {
		var userHeight = 1.8
		var userWeight float64 = 100
		var IMT = userWeight / math.Pow(userHeight, 2)
		fmt.Print(IMT)
}
