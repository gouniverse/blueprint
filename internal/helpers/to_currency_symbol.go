package helpers

import (
	"strconv"
	"strings"
)

func ToCurrencySymbol(currency string) string {
	if strings.ToUpper(currency) == "GBP" {
		return "&pound;"
	}
	if strings.ToUpper(currency) == "EUR" {
		return "&euro;"
	}
	if strings.ToUpper(currency) == "USD" {
		return "$"
	}
	return currency
}

func StrToPrice(price string, currency string) string {
	priceFloat, errPrice := strconv.ParseFloat(price, 64)

	if errPrice != nil {
		return "n/a"
	}

	price = strconv.FormatFloat(priceFloat, 'f', 2, 64)
	return ToCurrencySymbol(currency) + price
}

func FloatToPrice(priceFloat float64, currency string) string {
	price := strconv.FormatFloat(priceFloat, 'f', 2, 64)
	return ToCurrencySymbol(currency) + price
}

func StringToPrice(priceString string, currency string) (string, error) {
	priceFloat, errPrice := strconv.ParseFloat(priceString, 64)
	if errPrice != nil {
		return "", errPrice
	}
	return FloatToPrice(priceFloat, currency), nil
}
