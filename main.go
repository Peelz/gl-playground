package main

import "strconv"

func main() {
	number1 := "123"
	number2 := "456"
	ans := "0"
	carry := 0

	for number1 != "" || number2 != "" {
		var op int
		if number1 == "" {
			op = strconv.ParseInt(number2, 10, 32) + carry
		} else if number2 == "" {
			op = strconv.ParseInt(number1, 10, 32) + carry
		} else {
			op = strconv.ParseInt(number1, 10, 32) + strconv.ParseInt(number2) + carry
		}

		if op >= 10 {
			carry = 1
			op = op % 10
		}
		ans = op + ans
	}
}
