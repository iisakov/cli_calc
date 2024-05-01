package NumSysTransform

import (
	"cli_calc/SDK/NumSysTransform/model"
	"sort"
)

func AtoR(arabianNum int) (result string, err error) {
	keys := make([]int, 0)
	for k, _ := range model.ArabianGlyphs {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for {
		for _, decimalVal := range keys {
			if arabianNum >= decimalVal {
				result += model.ArabianGlyphs[decimalVal]
				arabianNum -= decimalVal
				break
			}
		}
		if arabianNum <= 0 {
			break
		}
	}
	return
}
