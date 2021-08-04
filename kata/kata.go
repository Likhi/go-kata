package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func FindOutlier(integers []int) int {
	return 0
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

func MultiTable(number int) string {
	result := ""
	for i := 1; i < 11; i++ {
		result += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(number*i) + "\n"
	}
	return result[:len(result)-1]
}

// https://www.codewars.com/kata/589273272fab865136000108
func BlackOrWhiteKey(keyPressCount int) string {
	w, b := "white", "black"
	keys := map[int]string{
		1:  w,
		2:  b,
		3:  w,
		4:  w,
		5:  b,
		6:  w,
		7:  b,
		8:  w,
		9:  w,
		10: b,
		11: w,
		12: b,
	}
	for keyPressCount > 88 {
		keyPressCount -= 88
	}
	for keyPressCount > 12 {
		keyPressCount -= 12
	}
	k := keys[keyPressCount]
	return k
}

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
