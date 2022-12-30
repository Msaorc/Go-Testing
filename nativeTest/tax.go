package tax

import "time"

func CalculateTax(amout float64) float64 {
	if amout <= 0 {
		return 0
	}
	if amout >= 1000 && amout < 20000 {
		return 10.0
	}
	if amout >= 20000 {
		return 20.00
	}
	return 5.0
}

func CalculateTaxMicrosecond(amout float64) float64 {
	time.Sleep(time.Microsecond)
	if amout == 0 {
		return 0
	}
	if amout >= 1000 {
		return 10.0
	}
	return 5.0
}
