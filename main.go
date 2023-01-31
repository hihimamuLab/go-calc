package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	equation := scanner.Text()
	slice := strings.Split(equation, " ")
	var num1 *int = nil
	var num2 *int = nil
	var operand string = ""
	for i, str := range slice {
		if i%2 == 0 {
			number, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("数値を入力してください")
				return
			}
			num1 = num2
			num2 = &number
		} else {
			switch str {
			case "+":
				operand = "+"
			case "-":
				operand = "-"
			case "*":
				operand = "*"
			case "/":
				operand = "/"
			default:
				fmt.Println("正しい記号を入力してください")
				return
			}
		}
		if num1 != nil && num2 != nil {
			temp := 0
			switch operand {
			case "+":
				temp = *num1 + *num2
			case "-":
				temp = *num1 - *num2
			case "*":
				temp = *num1 * *num2
			case "/":
				temp = *num1 / *num2
			}
			*num2 = temp
			num1 = nil
		}
	}
	fmt.Println(*num2)
}
