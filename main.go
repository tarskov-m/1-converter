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

	// Инициализация map
	validConversion := map[string][]string{
		"":    {"USD", "EUR", "RUB"},
		"USD": {"EUR", "RUB"},
		"EUR": {"USD", "RUB"},
		"RUB": {"USD", "EUR"},
	}

	exchangeRates := map[string]float64{
		"USD-EUR": UsdToEur,
		"USD-RUB": UsdToRub,
		"EUR-RUB": EurToRub,
		"EUR-USD": 1 / UsdToEur,
		"RUB-USD": 1 / UsdToRub,
		"RUB-EUR": 1 / EurToRub,
	}

	for {
		currency1 = userInputCurrency("", &validConversion)
		if currency1 == "" {
			fmt.Println("Программа завершена")
			return
		}
		amount = userInputAmount()

		currency2 = userInputCurrency(currency1, &validConversion)
		if currency2 == "" {
			fmt.Println("Программа завершена")
			return
		}

		calculate(amount, currency1, currency2, &exchangeRates)

		userInput = stopProgram()
		if !userInput {
			fmt.Println("Программа завершена")
			return
		}
	}
}

func userInputCurrency(currency string, validConversion *map[string][]string) string {
	var number int

	options := (*validConversion)[currency]
	// Создаем копию, чтобы не изменять оригинальный слайс
	optionsCopy := make([]string, len(options))
	copy(optionsCopy, options)
	optionsCopy = append(optionsCopy, "Выход")

	for {
		fmt.Println("Введите валюту:")
		for i, option := range optionsCopy {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		if _, err := fmt.Scanln(&number); err != nil {
			fmt.Println("Введите корректное значение")
			continue
		}

		if number > len(optionsCopy) || number < 1 {
			fmt.Println("Введите корректное значение")
			continue
		}

		if number == len(optionsCopy) {
			return ""
		}

		return optionsCopy[number-1]
	}
}

func userInputAmount() float64 {
	var amount float64

	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scanln(&amount)
		if err != nil {
			fmt.Println("Введите корректное значение")
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
			continue
		} else if amount <= 0 {
			fmt.Println("Введите значение больше 0")
			continue
		} else {
			return amount
		}
	}
}

func calculate(amount float64, currency1 string, currency2 string, exchangeRates *map[string]float64) {
	key := currency1 + "-" + currency2

	if rate, ok := (*exchangeRates)[key]; ok {
		fmt.Printf("Результат: %.2f %s\n", amount*rate, currency2)
	} else {
		fmt.Println("Неизвестная валютная пара")
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
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
			continue
		} else if number < 1 || number > 2 {
			fmt.Println("Введите значение от 1 до 2")
			continue
		} else {
			return number == 1
		}
	}
}
