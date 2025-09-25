package main

import "fmt"

const UsdToEur = 0.84
const UsdToRub = 84.02
const EurToRub = UsdToRub / UsdToEur

func main() {
	var currency1 string
	var currency2 string
	var amount float64
	var userInput bool

	fmt.Println("___ Конвертер валют ___")

	for {
		currency1 = userInputCurrency("")
		if currency1 == "" {
			fmt.Println("Программа завершена")
			return
		}
		amount = userInputAmount()

		currency2 = userInputCurrency(currency1)
		if currency2 == "" {
			fmt.Println("Программа завершена")
			return
		}

		calculate(amount, currency1, currency2)

		userInput = stopProgram()
		if !userInput {
			fmt.Println("Программа завершена")
			return
		}
	}
}

func userInputCurrency(currency string) string {
	var number int
	var value string

	for {
		if currency == "" {
			fmt.Println("Введите валюту:")
			fmt.Println("1. USD")
			fmt.Println("2. EUR")
			fmt.Println("3. RUB")
			fmt.Println("4. Выход")
			_, err := fmt.Scanln(&number)
			if err != nil {
				fmt.Println("Введите корректное значение")
				continue
			}
			if number == 4 {
				return value
			} else if number > 4 {
				fmt.Println("Введите корректное значение")
				continue
			} else {
				switch number {
				case 1:
					value = "USD"
				case 2:
					value = "EUR"
				case 3:
					value = "RUB"
				}
				return value
			}
		} else {
			if currency == "USD" {
				fmt.Println("Введите валюту:")
				fmt.Println("1. EUR")
				fmt.Println("2. RUB")
				fmt.Println("3. Выход")
				_, err := fmt.Scanln(&number)
				if err != nil {
					fmt.Println("Введите корректное значение")
					continue
				}
				if number == 3 {
					return value
				} else if number > 3 {
					fmt.Println("Введите корректное значение")
					continue
				} else {
					switch number {
					case 1:
						value = "EUR"
					case 2:
						value = "RUB"
					}
					return value
				}
			} else if currency == "EUR" {
				fmt.Println("Введите валюту:")
				fmt.Println("1. USD")
				fmt.Println("2. RUB")
				fmt.Println("3. Выход")
				_, err := fmt.Scanln(&number)
				if err != nil {
					fmt.Println("Введите корректное значение")
					continue
				}
				if number == 3 {
					return value
				} else if number > 3 {
					fmt.Println("Введите корректное значение")
					continue
				} else {
					switch number {
					case 1:
						value = "USD"
					case 2:
						value = "RUB"
					}
					return value
				}
			} else if currency == "RUB" {
				fmt.Println("Введите валюту:")
				fmt.Println("1. USD")
				fmt.Println("2. EUR")
				fmt.Println("3. Выход")
				_, err := fmt.Scanln(&number)
				if err != nil {
					fmt.Println("Введите корректное значение")
					continue
				}
				if number == 3 {
					return value
				} else if number > 3 {
					fmt.Println("Введите корректное значение")
					continue
				} else {
					switch number {
					case 1:
						value = "USD"
					case 2:
						value = "EUR"
					}
					return value
				}
			}
		}
	}
}

func userInputAmount() float64 {

	var amount float64

	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scanln(&amount)
		if err != nil {
			fmt.Println("Введите корректное значение")
			continue
		} else if amount == 0 {
			fmt.Println("Введите значение больше 0")
			continue
		} else {
			return amount
		}
	}
}

func calculate(amount float64, currency1 string, currency2 string) {
	if currency1 == "USD" && currency2 == "EUR" {
		fmt.Printf("%.2f USD = %.2f EUR\n", amount, amount*UsdToEur)
	} else if currency1 == "USD" && currency2 == "RUB" {
		fmt.Printf("%.2f USD = %.2f RUB\n", amount, amount*UsdToRub)
	} else if currency1 == "EUR" && currency2 == "USD" {
		fmt.Printf("%.2f EUR = %.2f USD\n", amount, amount/UsdToEur)
	} else if currency1 == "EUR" && currency2 == "RUB" {
		fmt.Printf("%.2f EUR = %.2f RUB\n", amount, amount*EurToRub)
	} else if currency1 == "RUB" && currency2 == "USD" {
		fmt.Printf("%.2f RUB = %.2f USD\n", amount, amount/UsdToRub)
	} else if currency1 == "RUB" && currency2 == "EUR" {
		fmt.Printf("%.2f RUB = %.2f EUR\n", amount, amount/EurToRub)
	}
}

func stopProgram() bool {
	var number int

	for {
		fmt.Println("Хотите продолжить?")
		fmt.Println("1. Да")
		fmt.Println("2. Нет")
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Введите корректное значение")
			continue
		} else if number == 0 || number > 2 {
			if number == 0 {
				fmt.Println("Введите значение больше 0")
				continue
			} else {
				fmt.Println("Введите значение меньше 3")
				continue
			}
		} else {
			if number == 1 {
				return true
			} else if number == 2 {
				return false
			}
		}
	}
}
