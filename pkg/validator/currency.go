package validator

type CurrencyValidator struct {
	ValidatorInterface
}

func (v *CurrencyValidator) Validate(value interface{}) bool {

	available_value := []interface{}{"USD", "EUR"}

	for _, v := range available_value {
		if v == value {
			return true
		}
	}

	return false
}
