package main

import (
	"os"
	"fmt"
	"errors"
	"strconv"
)

func main() {
	fmt.Println(os.Args[1])
	if len (os.Args) == 1 {
		fmt.Println("Floats please")
		os.Exit(1)
	}
	arguments := os.Args
	var err = errors.New("An error")
	k := 1
	var n float64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("No floats")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n

	for i := 2; i < len(arguments); i++ {
		n, err = strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
