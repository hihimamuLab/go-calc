package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //文字入力
	scanner.Scan()
	equation := scanner.Text()            //入力した文書を変数に代入
	slice := strings.Split(equation, " ") //その変数をスペースで区切って新しい変数に代入
	var num1 *int = nil                   //nilが入った空の変数の作成
	var num2 *int = nil
	var operand string = ""     //演算子を入れる空の変数の作成
	for i, str := range slice { //slice変数から何番目の要素かと要素の内容を取得してスライスの数だけ繰り返す
		if i%2 == 0 { //要素の番目が偶数(数値)の場合の処理
			number, err := strconv.Atoi(str) //要素の内容を数値に変換するのとエラーハンドリングのためのerr変数の作成と代入
			if err != nil {                  //エラーが起きた場合の処理
				fmt.Println("数値を入力してください")
				return //処理終了
			}
			num1 = num2    //num2をnum1に代入(値の保存)
			num2 = &number //numberのポインタをnum2に代入
		} else { //要素の番目が奇数(演算子)の場合の処理
			switch str { //str(要素の内容)を引数にする
			case "+": //和だったときの処理
				operand = "+"
			case "-": //差だったときの処理
				operand = "-"
			case "*": //積だったときの処理
				operand = "*"
			case "/": //商だったときの処理
				operand = "/"
			default: //例外処理
				fmt.Println("正しい記号を入力してください")
				return //処理終了
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
