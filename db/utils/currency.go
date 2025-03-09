package utils

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP:
		return true
	}
	return false
}
