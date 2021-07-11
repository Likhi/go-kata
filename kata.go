package main

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/58845748bd5733f1b300001f/train/go
func RangeBitCount(a, b int) int {
	ones := 0
	for i := a; i < b+1; i++ {
		//fmt.Println(i)
		h := fmt.Sprintf("%b", i)
		ones += strings.Count(h, "1")
	}
	return ones
}

// https://www.codewars.com/kata/5264d2b162488dc400000001/train/go
func SpinWords(str string) string {
	strAsSlice := strings.Split(str, " ")
	ans := ""
	for _, ss := range strAsSlice {
		if len(ss) >= 5 {
			ans += " " + reverse(ss)
		} else {
			ans += " " + ss
		}
	}
	return ans[1:]
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
