package main

/*
2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
Площадь круга должна вводиться пользователем с клавиатуры.
*/

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var s, c, d float64
	var userInput string

	fmt.Println("This application calculates the depth and the circumference using just circle's area.")
	fmt.Printf("please, enter the circle's area ( S > 0 ): ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в числовой формат
	s, err := strconv.ParseFloat(userInput, 64)

	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Println("...entered data cannot be recognized as a nubmer.")
		fmt.Println(err)
		return
	} 
	
	//проверка на валидность: > 0
	if s <=0 {
		// невалидные введенные данные 
		fmt.Println("...entered data is not valid ( <=0 ).")
		return		
	}
	
	// ошибок нет, производятся вычисления и выводится результата
	d = math.Sqrt(4 * s / math.Pi)
	c = math.Pi * d

	fmt.Printf("The depth is %.2f, circumference - %.2f", d, c)
}
