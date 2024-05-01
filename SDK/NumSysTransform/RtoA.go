package NumSysTransform

import (
	"cli_calc/SDK/NumSysTransform/model"
	"errors"
	"strings"
)

func RtoA(romanNum string) (int, error) {
	result := 0

	romanNum = strings.ToUpper(romanNum)
	result, ok := model.RomanGlyphs[romanNum]
	if !ok {
		return -1, errors.New("неизвестное число")
	}

	return result, nil
}
