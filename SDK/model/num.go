package model

import (
	"cli_calc/SDK/NumSysTransform"
	"strconv"
)

type Num struct {
	NumVal  int
	NumType string
}

func (n *Num) Creat(s string) (err error) {
	if n.NumVal, err = strconv.Atoi(s); err == nil {
		n.NumType = "digit"
	} else {
		n.NumVal, err = NumSysTransform.RtoA(s)
		n.NumType = "roman"
	}
	return
}
