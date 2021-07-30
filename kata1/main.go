package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MultiTable(3))
}

func MultiTable(number int) string {
	result := ""
	for i := 1; i < 11; i++ {
		result += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(number*i) + "\n"
	}
	return result[:len(result)-1]
}
