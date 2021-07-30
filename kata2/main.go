package main

import "fmt"

func main() {
	fmt.Println(Multiple3And5(4))
}

func Multiple3And5(number int) int {
	m := make(map[int]int)

	for i := 1; i < (number * 5); i++ {
		key := i * 3
		if (key) < number {
			m[key] = 1
		} else {
			break
		}
	}

	for i := 1; i < (number * 5); i++ {
		key := i * 5
		if (key) < number {
			m[key] = 1
		} else {
			break
		}
	}

	sum := 0
	for k, _ := range m {
		sum += k
	}
	return sum
}
