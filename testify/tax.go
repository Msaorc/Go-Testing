package tax

import "errors"

func CalculateTax(amout float64) (float64, error) {
	if amout <= 0 {
		return 0, errors.New("Amout must be greater than 0")
	}
	if amout >= 1000 && amout < 20000 {
		return 10.0, nil
	}
	if amout >= 20000 {
		return 20.00, nil
	}
	return 5.0, nil
}
