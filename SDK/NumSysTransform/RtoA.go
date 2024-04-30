package NumSysTransform

import (
	"cli_calc/SDK/NumSysTransform/model"
	"errors"
)

func RtoA(romanNum string) (int, error) {
	result := 0
	countCh := len(romanNum)
	var vals []int

	if countCh == 1 {
		ch, ok := model.RomanGlyphs[romanNum]
		if !ok {
			return -1, errors.New("неизвестный символ в числе")
		}
		result += ch
	} else if countCh > 1 {
		for _, ch := range romanNum {
			val, ok := model.RomanGlyphs[string(ch)]
			if !ok {
				return -1, errors.New("неизвестный символ в числе")
			}
			vals = append(vals, val)
		}
		result += vals[len(vals)-1]
	}

	if countCh == 2 {
		result = addNum(result, vals[0], vals[1])
	} else if countCh > 2 {
		for i := countCh - 1; i > 0; i-- {
			result = addNum(result, vals[i-1], vals[i])
		}
	}

	return result, nil
}

func addNum(result, a, b int) int {
	if a >= b {
		result += a
	} else {
		result -= a
	}
	return result
}
