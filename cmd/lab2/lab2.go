package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var fbiReviews = make(map[string]float64, 0)
var specialistReviews = make(map[string]float64, 0)

func main() {
	_, _ = PerformDeposit("12345678", 0)
	_, _ = PerformDeposit("12345678", 123.45)
	_, _ = PerformDeposit("12345678", 200000)
	_, _ = PerformDeposit("12345678", 2000000)
}

func PerformDeposit(checkNumber string, amount float64) (string, error) {
	fmt.Printf("Trying to Deposit Check [%s] of amount [%f]\n", checkNumber, amount)
	confirmationNumber, err := DepositCheck(checkNumber, amount)

	if err != nil {
		if strings.Contains(err.Error(), "bad amount") {
			return "", errors.Wrap(err, "fix the amount")
		}

		if strings.Contains(err.Error(), "specialist review") {
			SubmitForSpecialistReview(checkNumber, amount)
			return "FAKECONF-" + checkNumber, nil
		}

		if strings.Contains(err.Error(), "FBI") {
			SubmitToFBI(checkNumber, amount)
			return "FAKECONF-" + checkNumber, nil
		}

		return "", errors.Wrap(err, "unexpected error")
	}

	fmt.Printf("Success. Got Confirmation Number [%s]\n", confirmationNumber)
	return confirmationNumber, nil
}

func SubmitForSpecialistReview(checkNumber string, amount float64) {
	fmt.Printf("Submitting Check [%s] of Amount [%f] for Specialist Review\n", checkNumber, amount)
	specialistReviews[checkNumber] = amount
}

func SubmitToFBI(checkNumber string, amount float64) {
	fmt.Printf("Submitting Check [%s] of Amount [%f] to FBI\n", checkNumber, amount)
	fbiReviews[checkNumber] = amount
}

func DepositCheck(checkNumber string, amount float64) (string, error) {
	if checkNumber == "" {
		return "", errors.New("bad input: empty check number")
	}

	if amount < 0 {
		return "", errors.New("bad input: amount cannot be negative")
	}

	if amount == 0 {
		return "", errors.New("bad input: amount cannot be zero")
	}

	if amount > 1000000 {
		return "", errors.New("suspicious activity, call FBI")
	}

	if amount > 100000 {
		return "", errors.New("suspicious activity, need specialist review")
	}

	var confirmationNumber = "CONF-" + checkNumber
	fmt.Printf("Deposited Check [%s] of amount [%f]. Confirmation Number [%s]\n", checkNumber, amount, confirmationNumber)
	return confirmationNumber, nil
}
