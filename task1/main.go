package main

/*
1. Напишите программу для вычисления площади прямоугольника.
Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var a, b float64
	var userInput string

	fmt.Println("This application calculates the rectangle's area. Rectangle's dimentions are to be set by the user.")
	fmt.Printf("please, enter the length of side A: ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в числовой формат
	a, err := strconv.ParseFloat(userInput, 64)
	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Printf("...entered data cannot be recognized as a nubmer.\n")
		os.Exit(2)
	}

	// ошибок нет, идет запрос второго числа
	fmt.Printf("please, enter the length of side B: ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в числовой формат
	b, err = strconv.ParseFloat(userInput, 64)
	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Printf("...entered data cannot be recognized as a nubmer.\n")
		os.Exit(2)
	}

	// ввод данных не вызвал ошибок, выводится результат
	fmt.Printf("You have entered rectangle with sides %.2fx%.2f. Tht area is %.2f.", a, b, a*b)
}
