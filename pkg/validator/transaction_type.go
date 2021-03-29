package validator

type TransactionTypeValidator struct {
	ValidatorInterface
}

func (v *TransactionTypeValidator) Validate(value interface{}) bool {

	available_value := []interface{}{"deposit", "withdrowal"}

	for _, v := range available_value {
		if v == value {
			return true
		}
	}

	return false
}
