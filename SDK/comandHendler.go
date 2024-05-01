package comandHandler

import (
	"cli_calc/SDK/model"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func PrintHelp() {
	fmt.Printf(`
Что умеет калькулятор:
		
Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b
Данные передаются в одну строку.
Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.	
Калькулятор принимает принимать на вход числа от 1 до 10 включительно, не более.
На выходе числа не ограничиваются по величине и могут быть любыми.
Калькулятор умеет работать только с целыми числами.
Калькулятор умеет работать только с арабскими или римскими цифрами одновременно.
При вводе римских чисел ответ выводится римскими цифрами, при вводе арабских — арабскими.
Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль.
Результатом работы калькулятора с римскими числами могут быть только положительные числа.

Также программа понимает следующие команды:
help, -h: вывести подсказку (данное сообщение)
exit, -q: завершить работу программы
version, -v: вывести версию программы
`)
}

func PrintMessage(messages ...string) {
	response := "\n"
	for _, message := range messages {
		response += message + "\n"
	}
	fmt.Println(response)
}

func PrintError(messages ...string) {
	response := "\n"
	for _, message := range messages {
		response += message + "\n"
	}
	response += "Для ознакомления с функционалом cli_calc используйте команду 'help'."
	fmt.Println(response)
}

func PrintErrorAndExit(messages ...string) {
	response := "\n"
	for _, message := range messages {
		response += message + "\n"
	}
	response += "Утилита закрыта."

	fmt.Println(response)
	os.Exit(1)
}

func Exit(messages ...string) {
	response := "\n"
	for _, message := range messages {
		response += message + "\n"
	}
	response += "Спасибо, что воспользовались cli_calc от [by_artisan]"

	fmt.Println(response)
	os.Exit(0)
}

func CheckOperator(s string) (bool, error) {
	result, err := regexp.MatchString(`^[\+\-\:\*\/]$`, s)
	if err != nil {
		return false, err
	}
	return result, nil
}

func Calculate(nums [2]model.Num, operator string) (model.Num, error) {
	var result model.Num
	var err error

	if nums[0].NumType == nums[1].NumType {
		result.NumType = nums[0].NumType
	} else {
		err = errors.New("в выражении используются разные типы чисел")
	}

	switch operator {
	case "+":
		result.NumVal = nums[0].NumVal + nums[1].NumVal
	case "-":
		result.NumVal = nums[0].NumVal - nums[1].NumVal
		if result.NumVal < 1 && result.NumType == "roman" {
			err = errors.New("результат вычисления римских чисел не может быть меньше единицы")
		}
	case "*":
		result.NumVal = nums[0].NumVal * nums[1].NumVal
	case ":":
		fallthrough
	case "/":
		result.NumVal = nums[0].NumVal / nums[1].NumVal
	default:
		err = errors.New("при вычислении выражения что-то пошло не так")
	}

	return result, err
}
