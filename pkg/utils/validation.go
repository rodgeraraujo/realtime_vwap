package utils

import (
	"regexp"
)

var (
	tradingRegexValidation = `^([A-Z]{3}\-[A-Z]{3},)*([A-Z]{3}\-[A-Z]{3})$`
)

func IsValidTradingPairs(tradingPairs string) bool {
	r, _ := regexp.Compile(tradingRegexValidation)
	return r.MatchString(tradingPairs)
}
