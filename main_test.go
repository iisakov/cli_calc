package main

import (
	comandHandler "cli_calc/SDK"
	"cli_calc/SDK/NumSysTransform"
	"cli_calc/SDK/model"
	"testing"
)

func TestRightRtoA(t *testing.T) {
	rightTestNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(rightTestNums); i++ {
		result, err := NumSysTransform.RtoA(rightTestNums[i])
		if err != nil {
			t.Fatalf("Ошибка в валидных данных. %s", rightTestNums[i])
		}
		if result != i+1 {
			t.Fatalf("Ошибка в валидных данных. %s (%d) не совпадает с %d", rightTestNums[i], result, i)
		}
	}
}

func TestFalseRtoA(t *testing.T) {
	failTestNums := []string{"IIII", "IIV", "XVX", "XXXX", "CV", "1", "XI", "0", "Сломата"}
	for i := 0; i < len(failTestNums); i++ {
		_, err := NumSysTransform.RtoA(failTestNums[i])
		if err == nil {
			t.Fatalf("Ошибка в сломаных данных. %s", failTestNums[i])
		}
	}
}

func TestRightAtoR(t *testing.T) {
	rightTestNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(rightTestNums); i++ {
		arabianNum := i + 1
		result, err := NumSysTransform.AtoR(arabianNum)
		if err != nil {
			t.Fatalf("Ошибка в валидных данных. %d", arabianNum)
		}
		if result != rightTestNums[i] {
			t.Fatalf("Ошибка в валидных данных. %d не совпадает с %s", arabianNum, rightTestNums[i])
		}
	}
}

func TestCheckRightOperator(t *testing.T) {
	rightOperators := []string{"+", "-", "*", "/", ":"}
	for i := 0; i < len(rightOperators); i++ {
		operator := rightOperators[i]
		_, err := comandHandler.CheckOperator(operator)
		if err != nil {
			t.Fatalf("Ошибка в валидных данных. %s", operator)
		}
	}
}

func TestCheckFalseOperator(t *testing.T) {
	rightOperators := []string{"_", "'", "Деление", "сломата", ""}
	for i := 0; i < len(rightOperators); i++ {
		operator := rightOperators[i]
		_, err := comandHandler.CheckOperator(operator)
		if err == nil {
			t.Fatalf("Ошибка в сломаных данных. %s", operator)
		}
	}
}

func TestCalculate(t *testing.T) {
	type ex struct {
		nums     [2]model.Num
		operator string
	}
	testEx := []ex{{nums: [2]model.Num{{NumVal: 1, NumType: "roman"}, {NumVal: 1, NumType: "roman"}}, operator: "/"},
		{nums: [2]model.Num{{NumVal: 2, NumType: "digit"}, {NumVal: 1, NumType: "digit"}}, operator: ":"},
		{nums: [2]model.Num{{NumVal: 2, NumType: "roman"}, {NumVal: 1, NumType: "roman"}}, operator: "+"},
		{nums: [2]model.Num{{NumVal: 5, NumType: "digit"}, {NumVal: 1, NumType: "digit"}}, operator: "-"},
		{nums: [2]model.Num{{NumVal: 10, NumType: "roman"}, {NumVal: 2, NumType: "roman"}}, operator: "/"},
		{nums: [2]model.Num{{NumVal: 3, NumType: "digit"}, {NumVal: 2, NumType: "digit"}}, operator: "*"},
		{nums: [2]model.Num{{NumVal: 1, NumType: "roman"}, {NumVal: 6, NumType: "roman"}}, operator: "+"},
		{nums: [2]model.Num{{NumVal: 7, NumType: "digit"}, {NumVal: 1, NumType: "digit"}}, operator: "+"},
		{nums: [2]model.Num{{NumVal: 3, NumType: "roman"}, {NumVal: 3, NumType: "roman"}}, operator: "*"},
		{nums: [2]model.Num{{NumVal: 10, NumType: "digit"}, {NumVal: 1, NumType: "digit"}}, operator: "/"}}

	numTypes := []string{"roman", "digit"}
	for i := 0; i < len(testEx); i++ {
		currentType := numTypes[i%2]
		currentEx := testEx[i]
		result, _ := comandHandler.Calculate(currentEx.nums, currentEx.operator)
		if result.NumVal != i+1 {
			t.Fatalf("Ошибка в валидных данных. %d не совпадает с %d", result.NumVal, i+1)
		}
		if result.NumType != currentType {
			t.Fatalf("Ошибка в валидных данных. %s не совпадает с %s", result.NumType, currentType)
		}

	}
}
