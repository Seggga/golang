package main

/*
1. Напишите программу для вычисления площади прямоугольника.
Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
*/

import (
	"fmt"
	"strconv"
)

func main() {

	var a, b float64
	var userInput string

	fmt.Println("This application calculates the rectangle's area. Rectangle's dimentions are to be set by the user.")
	fmt.Printf("please, enter the length of side A (number > 0) : ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в числовой формат
	a, err := strconv.ParseFloat(userInput, 64)
	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Printf("...entered data cannot be recognized as a nubmer.\n")
		fmt.Printf("%v", err)
		return
	}
	
	// проверка на валидность введенных данных
	if a <= 0 {
		// невалидные введенные данные 
		fmt.Printf("...entered data is not valid ( <=0 ).\n")
		return
	}

	// ошибок нет, идет запрос второго числа
	fmt.Printf("please, enter the length of side B  (number > 0): ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в числовой формат
	b, err = strconv.ParseFloat(userInput, 64)
	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Printf("...entered data cannot be recognized as a nubmer.\n")
		fmt.Printf("%v", err)
		return
	}
	
	// проверка на валидность введенных данных
	if b <= 0 {
		// невалидные введенные данные 
		fmt.Printf("...entered data is not valid ( <=0 ).\n")
		return
	}

	// ввод данных не вызвал ошибок, выводится результат
	fmt.Printf("You have entered rectangle with sides %.2fx%.2f. Tht area is %.2f.", a, b, a*b)
}
