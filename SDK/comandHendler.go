package comandHandler

import (
	"cli_calc/SDK/model"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func PrintHelp() {
	fmt.Println(`Калькутор cli_calc`)
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
