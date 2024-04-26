package utils

import (
	"fmt"

	"github.com/zuanet/zuad/domain/consensus/utils/constants"
)

// FormatZua takes the amount of sompis as uint64, and returns amount of ZUA with 8  decimal places
func FormatZua(amount uint64) string {
	res := "                   "
	if amount > 0 {
		res = fmt.Sprintf("%19.8f", float64(amount)/constants.SompiPerZua)
	}
	return res
}
