/*
3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количеству
сотен, десятков и единиц в этом числе.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num, ed, des, sot int
	var userInput string

	fmt.Println("Программа вычисляет значения разрядов трехзначного десятичного числа, введенного пользователем.")
	fmt.Printf("введите трехзначное число : ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в число
	num, err := strconv.Atoi(userInput)

	// проверка на наличие ошибки
	if err != nil {
		fmt.Printf("введенные данные не удается перевести в число.")		
	} else {
		fmt.Printf("число %d содержит:\n", num)

		// проверка знака введенного числа
		if num < 0 {
			num *= (-1)
		}
	
		// вычисление единиц, десятков и сотен 
		ed  = num % 10
		num = num / 10
		des = num % 10
		sot = num / 10
	
		// вывод результата
		fmt.Printf("- сотен    %d\n", sot)
		fmt.Printf("- десятков %d\n", des)
		fmt.Printf("- единиц   %d\n", ed)

	}

}
