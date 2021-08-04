package main

//https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("hell_no"))
}

func ToCamelCase(s string) string {
	sep := "_"
	s = strings.ReplaceAll(s, "-", sep)
	ss := strings.Split(s, sep)
	for i, v := range ss {
		if i != 0 {
			ss[i] = strings.ToUpper(string(v[0])) + v[1:]
		}
	}
	return strings.Join(ss, "")
}
