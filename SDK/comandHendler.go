package comandHandler

import (
	"cli_calc/SDK/model"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func PrintHelp() {
	fmt.Println(
		`Требования:
		
Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.
Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.	
Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
Калькулятор умеет работать только с целыми числами.
Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.
При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.
При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу.
При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу.
Результатом операции деления является целое число, остаток отбрасывается.
Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.
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

	result.NumType = nums[0].NumType

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
