package main

/*
3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количеству
сотен, десятков и единиц в этом числе.
*/

import (
	"fmt"
	"strconv"
)

func main() {

	var userNumber, position0, position1, position2 uint
	var userInput string

	fmt.Println("Application takes a 3-digit positive number and calculates entered digits.")
	fmt.Printf("Please enter 3-digit number : ")
	fmt.Scanln(&userInput)

	// перевод введенных данных в число
	userNumber, err := strconv.Atoi(userInput)

	// проверка на наличие ошибки
	if err != nil {
		// есть ошибка, выдается сообщение, вычисления не производятся
		fmt.Println("...entered data cannot be recognized as a nubmer.")
		fmt.Println(err)
		return
	}
	
	// проверка на валидность введенных данных
	if userNumber < 100 || userNumber > 999 {
		// невалидные введенные данные 
		fmt.Printf("...entered data is not valid ( not a 3-digit number ).\n")
		return
	}
	
	fmt.Printf("entered number %d contains:\n", userNumber)
	
	// вычисление единиц, десятков и сотен 
	position0 = num % 10
	userNumber = userNumber / 10
	position1 = userNumber % 10
	position2 = userNumber / 10
	
	// вывод результата
	fmt.Printf("- hundreds  %d\n", position2)
	fmt.Printf("- tens      %d\n", position1)
	fmt.Printf("- units     %d\n", position0)

}
