package main

import (
	"fmt"
	"strconv"
)

func largeNumberInDream(number1 string, number2 string) string {
	reverse1 := reverseString(number1)
	reverse2 := reverseString(number2)
	ans := ""
	carry := 0
	flag := 0
	for reverse1 != "" || reverse2 != "" || carry != 0 {
		op := 0

		if reverse1 == "" && reverse2 != "" {
			tmp, _ := strconv.ParseInt(string(reverse2[0]), 10, 32)
			op = int(tmp) + carry
			reverse2 = reverse2[1:]
		} else if reverse2 == "" && reverse1 != "" {
			tmp, _ := strconv.ParseInt(string(reverse1[0]), 10, 32)
			op = int(tmp) + carry
			reverse1 = reverse1[1:]

		} else if reverse2 != "" || reverse1 != "" {
			tmp1, _ := strconv.ParseInt(string(reverse1[0]), 10, 32)
			tmp2, _ := strconv.ParseInt(string(reverse2[0]), 10, 32)
			op = int(tmp1) + int(tmp2) + carry
			reverse1 = reverse1[1:]
			reverse2 = reverse2[1:]

		} else {
			ans = strconv.Itoa(carry) + ans
			return ans
		}

		if op >= 10 {
			carry = 1
			op = op % 10
		} else {
			carry = 0
		}
		ans = strconv.Itoa(op) + ans
		flag++
	}

	return ans
}

func reverseString(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result

	}
	return result
}

func main() {
	var str1 string
	var str2 string

	fmt.Println("number 1")
	fmt.Scanf("%s", &str1)
	fmt.Println("number 2")
	fmt.Scanf("%s", &str2)

	fmt.Println("ans", largeNumberInDream(str1, str2))
}
