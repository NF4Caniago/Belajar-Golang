package main

import (
	"fmt"
	"strings"
)

func main() {
	// const mytext = "konstan ini coy"
	// 
	// var text1 string
	// text1 = "Hello Afif"
	// var text2 string = "Hellowww"
	// text3 := "helllloooowwww"
	// text3 = "End"

	// var positifNumber uint8 = 90
	// var negatifNumber = -1243423644
	// var decimalNumber = 3.14
	// var exist bool = true
	// var text = "Hello World 111"
	// var textNo_escape = `"Hello World" adalah habit dari programmer`

	// fmt.Println("Bilangan Positif : %d\n", positifNumber)
	// fmt.Println("Bilangan Negatif : %d\n", negatifNumber)
	// fmt.Println("Bilangan Decimal : %f\n", decimalNumber)
	// fmt.Println("Bilangan Decimal : %4.f\n", decimalNumber)
	// fmt.Println("Exist? %t\n", exist)
	// fmt.Println(text)
	// fmt.Println(textNo_escape)

	var index = strings.Index("Afif Ilham Caniago", "Ca")
	fmt.Println(index)
}