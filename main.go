package main

import "fmt"

func main() {
	const UsdToEur = 0.84
	const UsdToRub = 84.02
	const EurToRub = UsdToRub / UsdToEur

	usd := userInput()
	fmt.Println(usd)
}

func userInput() float64 {
	var usd float64
	fmt.Print("Введите сумму USD: ")
	fmt.Scanln(&usd)

	return usd
}

func calculate(usd float64, currency1 string, currency2 string) {}
