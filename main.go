package main

import (
	"bufio"
	comandHandler "cli_calc/SDK"
	"cli_calc/SDK/model"
	"cli_calc/config"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Println("Введите выражение или команду (-h):")

	for {
		command, err := in.ReadString('\n')
		if err != nil {
			comandHandler.PrintErrorAndExit()
		}
		commandSlice := strings.Split(command, " ")

		if len(commandSlice) == 1 {

			switch command {

			case "help\n":
				fallthrough
			case "-h\n":
				comandHandler.PrintHelp()

			case "exit\n":
				fallthrough
			case "-q\n":
				comandHandler.Exit()

			case "version\n":
				fallthrough
			case "-v\n":
				comandHandler.PrintMessage("cli_calc [by_artisan] v:" + config.CalcVersion)

			default:
				comandHandler.PrintError("Неизвестная команда.")
			}

		} else if len(commandSlice) == 3 {
			var nums [2]model.Num
			var operator string

			operator, err := comandHandler.CheckOperator(commandSlice[1])
			if err != nil {
				comandHandler.PrintErrorAndExit(err.Error())
			}

			for i, commandPart := range []string{commandSlice[0], strings.TrimRight(commandSlice[2], "\n")} {
				num := new(model.Num)
				err = num.Creat(commandPart)
				if err != nil {
					comandHandler.PrintErrorAndExit(err.Error())
				}

				if num.NumVal > 10 || num.NumVal < 1 {
					comandHandler.PrintErrorAndExit("Одно из чисел больше 10 или меньше 1.")
				}

				nums[i] = *num
			}

			if nums[0].NumType != nums[1].NumType {
				comandHandler.PrintErrorAndExit("В выражении используются разные типы чисел")
			}

			result, err := comandHandler.Calculate(nums, operator)
			if err != nil {
				comandHandler.PrintErrorAndExit(err.Error())
			}

			comandHandler.PrintResult(result)

		} else {
			comandHandler.PrintError("Неизвестная команда.")
		}

	}
}
