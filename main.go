package main

import (
	"bufio"
	comandHandler "cli_calc/SDK"
	"cli_calc/SDK/NumSysTransform"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Println("Введите выражение:")

	for {
		command, err := in.ReadString('\n')
		if err != nil {
			comandHandler.PrintErrorAndExit()
		}
		comandSlice := strings.Split(command, " ")

		if len(comandSlice) == 1 {

			switch command {

			case "help\n":
				comandHandler.PrintHelp()

			case "exit\n":
			case "-q\n":
				comandHandler.Exit()

			default:
				comandHandler.PrintError("Неизвестная команда.")
			}

		} else if len(comandSlice) == 3 {
			result, err := NumSysTransform.RtoA(comandSlice[0])
			if err != nil {
				fmt.Println(err.Error())
				comandHandler.PrintErrorAndExit(err.Error())
			}
			fmt.Println(result)
		}

	}
}
