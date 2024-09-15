package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, CAD:
		return true
	}

	return false
}
