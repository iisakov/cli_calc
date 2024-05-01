package main

import (
	"bufio"
	comandHandler "cli_calc/SDK"
	"cli_calc/SDK/NumSysTransform"
	"cli_calc/SDK/model"
	"fmt"
	"os"
	"strconv"
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
			var nums [2]model.Num
			var operator string
			isValidOperator, err := comandHandler.CheckOperator(comandSlice[1])
			if err != nil {
				fmt.Println(err.Error())
				comandHandler.PrintErrorAndExit(err.Error())
			}
			if !isValidOperator {
				comandHandler.PrintErrorAndExit("Неверный символ действия выражения")
			} else {
				operator = comandSlice[1]

			}
			if nums[0].NumVal, err = strconv.Atoi(comandSlice[0]); err == nil {
				nums[0].NumType = "digit"
			} else {
				nums[0].NumVal, err = NumSysTransform.RtoA(comandSlice[0])
				if err != nil {
					comandHandler.PrintErrorAndExit(err.Error())
				}
				nums[0].NumType = "roman"
			}
			if nums[1].NumVal, err = strconv.Atoi(comandSlice[2][:len(comandSlice[2])-1]); err == nil {
				nums[1].NumType = "digit"
			} else {
				nums[1].NumVal, err = NumSysTransform.RtoA(comandSlice[2][:len(comandSlice[2])-1])
				if err != nil {
					comandHandler.PrintErrorAndExit(err.Error())
				}
				nums[1].NumType = "roman"
			}

			if nums[0].NumType != nums[1].NumType {
				comandHandler.PrintErrorAndExit("В выражении используются разные типы чисел")
			}

			for _, num := range nums {
				if num.NumVal > 10 {
					comandHandler.PrintErrorAndExit("Одно из чисел больше 10.")
				}
			}

			result, err := comandHandler.Calculate(nums, operator)
			if err != nil {
				comandHandler.PrintErrorAndExit(err.Error())
			}
			if result.NumType == "roman" {
				romanNumResult, err := NumSysTransform.AtoR(result.NumVal)
				if err != nil {
					comandHandler.PrintErrorAndExit(err.Error())
				}
				comandHandler.PrintMessage(romanNumResult)
			} else {
				comandHandler.PrintMessage(strconv.Itoa(result.NumVal))
			}

		} else {
			comandHandler.PrintErrorAndExit("Неизвестная команда.")
		}

	}
}
