package tax

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("Amout must be greater than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.00, nil
	}
	return 5.0, nil
}

func CalculateTaxMock(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.00
	}
	return 5.0
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTaxMock(amount)
	return repository.SaveTax(tax)
}
