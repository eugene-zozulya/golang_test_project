package validator

type ValidatorInterface interface {
	Validate(value interface{}) bool
}
