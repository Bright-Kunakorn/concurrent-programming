package main

import (
	"fmt"
	"strconv"
)

func sumEven(x []string) int {
	sumEven := 0
	for i := 0; i < len(x); i++ {
		if _, err := strconv.Atoi(x[i]); err != nil {
			if x[i] == ">" && i+1 < len(x) {
				x[i] = x[i+1]
			} else if x[i] == "<" && i-1 >= 0 {
				x[i] = x[i-1]
			}
		}
		if _, err := strconv.Atoi(x[i]); err != nil {
			if x[i] == ">" && i+2 < len(x) {
				x[i] = x[i+2]
			} else if x[i] == "<" && i-2 >= 0 {
				x[i] = x[i-2]
			}
		}
		if num, err := strconv.Atoi(x[i]); err == nil && num%2 == 0 {
			sumEven += num
		}
	}
	return sumEven
}

func sumOdd(x []string) int {
	sumOdd := 0
	for i := 0; i < len(x); i++ {
		if _, err := strconv.Atoi(x[i]); err != nil {
			if x[i] == ">" && i+1 < len(x) {
				x[i] = x[i+1]
			} else if x[i] == "<" && i-1 >= 0 {
				x[i] = x[i-1]
			}
		}
		if _, err := strconv.Atoi(x[i]); err != nil {
			if x[i] == ">" && i+2 < len(x) {
				x[i] = x[i+2]
			} else if x[i] == "<" && i-2 >= 0 {
				x[i] = x[i-2]
			}
		}
		if num, err := strconv.Atoi(x[i]); err == nil && num%2 != 0 {
			sumOdd += num
		}
	}
	return sumOdd
}

func main() {
	x2 := []string{"1", "2", "<", "<", "<", "<", "6", ">", ">", "7"}
	fmt.Println(sumEven(x2))
	fmt.Println(sumOdd(x2))
}
