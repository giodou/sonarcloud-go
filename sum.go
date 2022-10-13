package main

import "fmt"

func main()  {
	fmt.Println(Soma(5,10))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtracao(a int, b int) int {
	return a - b
}

func Divisao(a int, b int) int {
	return a / b
}

func DivisaoPor2(a int) int {
	return a / 2
}