package main

import "fmt"

func main() {

	var userInput string
	userInput = "y"

	fmt.Println("welcome to project calculator")

	for userInput == "y" {
		fmt.Print("please enter your first number: ")

		var firstNumber float32
		fmt.Scan(&firstNumber)

		fmt.Print("please enter your operation (+, -, *, /, [R/r], %) : ")

		var operation string
		fmt.Scan(&operation)

		fmt.Print("please enter your second number: ")

		var secondNumber float32
		fmt.Scan(&secondNumber)

		var ans float32

		switch operation {

		case "/":
			if secondNumber == 0 {
				fmt.Println("sorry we can't divide a number by 0")
			} else {
				ans = firstNumber / secondNumber
				fmt.Print("your answer is: ")
				fmt.Println(ans)
			}

		case "r", "R":
			if secondNumber == 0 {
				fmt.Println("sorry we can't divide a number by 0")
			} else {
				var newInterger int
				var oldInteger int
				var ansAsInt int
				newInterger = int(firstNumber)
				oldInteger = int(secondNumber)
				ansAsInt = newInterger % oldInteger
				fmt.Print("your answer is: ")
				fmt.Println(ansAsInt)
			}
		case "+":
			ans = firstNumber + secondNumber
			fmt.Print("your answer is: ")
			fmt.Println(ans)

		case "-":
			ans = firstNumber - secondNumber
			fmt.Print("your answer is: ")
			fmt.Println(ans)

		case "*":
			ans = firstNumber * secondNumber
			fmt.Print("your answer is: ")
			fmt.Println(ans)

		case "%":
			ans = (firstNumber / secondNumber) * 100
			fmt.Print("your answer is: ")
			fmt.Println(ans, "%")

		default:
			fmt.Println("sorry invalid operation")
		}

		fmt.Print("Do you want to do any other operations ? (y/n) : ")
		fmt.Scan(&userInput)

		for userInput != "y" && userInput != "n" {
			fmt.Println("sorry what you have entered is invalid")
			fmt.Print("Do you want to do any other operations ? (y/n) : ")
			fmt.Scan(&userInput)
		}
		if userInput == "n" {
			fmt.Println("Thank you for using my calculator! You can always come back to the best calculator in the world!")
		}

	}
}
