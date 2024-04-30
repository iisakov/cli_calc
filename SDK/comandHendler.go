package comandHandler

import (
	"fmt"
	"os"
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
	response += "Произошла непредвиденная ошибка, утилита закрыта."
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
