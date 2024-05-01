package comandHandler

import (
	"fmt"
	"os"
	"regexp"
)

func PrintHelp() {
	fmt.Println(`Калькутор cli_calc`)
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
