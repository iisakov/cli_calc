package main

import (
	"bufio"
	commandHandler "cli_calc/SDK"
	"cli_calc/SDK/model"
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
			commandHandler.PrintErrorAndExit()
		}
		commandSlice := strings.Split(command, " ")

		if len(commandSlice) == 1 {
			commandHandler.Handle(command)

		} else if len(commandSlice) == 3 {
			var nums [2]model.Num
			var operator string

			operator, err := commandHandler.CheckOperator(commandSlice[1])
			if err != nil {
				commandHandler.PrintErrorAndExit(err.Error())
			}

			for i, commandPart := range []string{commandSlice[0], strings.TrimRight(commandSlice[2], "\n")} {
				num := new(model.Num)
				err = num.Creat(commandPart)
				if err != nil {
					commandHandler.PrintErrorAndExit(err.Error())
				}

				if num.NumVal > 10 || num.NumVal < 1 {
					commandHandler.PrintErrorAndExit("Одно из чисел больше 10 или меньше 1.")
				}

				nums[i] = *num
			}

			if nums[0].NumType != nums[1].NumType {
				commandHandler.PrintErrorAndExit("В выражении используются разные типы чисел")
			}

			result, err := commandHandler.Calculate(nums, operator)
			if err != nil {
				commandHandler.PrintErrorAndExit(err.Error())
			}

			commandHandler.PrintResult(result)

		} else {
			commandHandler.PrintError("Неизвестная команда.")
		}

	}
}
