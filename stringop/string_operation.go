package main

import (
	"fmt"
	"strconv"
)

func stringPlus(number1 string, number2 string) string {
	reverse1 := reverseString(number1)
	reverse2 := reverseString(number2)
	ans := ""
	carry := 0
	flag := 0
	for reverse1 != "" || reverse2 != "" {
		op := 0
		carry = 0

		if reverse1 == "" {
			tmp, _ := strconv.ParseInt(string(reverse2[0]), 10, 32)
			op = int(tmp) + carry
			reverse2 = reverse2[1:]
		} else if reverse2 == "" {
			tmp, _ := strconv.ParseInt(string(reverse1[0]), 10, 32)
			op = int(tmp) + carry
			reverse1 = reverse2[1:]

		} else {
			tmp1, _ := strconv.ParseInt(string(reverse1[0]), 10, 32)
			tmp2, _ := strconv.ParseInt(string(reverse2[0]), 10, 32)
			// println(reverse1, reverse2)

			op = int(tmp1) + int(tmp2) + carry
			reverse1 = reverse1[1:]
			reverse2 = reverse2[1:]

		}

		if op >= 10 {
			carry = 1
			op = op % 10
		}
		ans = strconv.Itoa(op) + ans
		// fmt.Println(ans, flag)
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
	fmt.Println("ans", stringPlus("123", "456"))
}
