package main

import "testing"

func TestGood(t *testing.T) {
	confirmationNumber, err := DepositCheck("12345", 100)
	assertNoError(t, err)
	assertConfirmationNumber(t, confirmationNumber, "CONF-12345")
}

func assertConfirmationNumber(t *testing.T, confirmationGot string, confirmationWant string) {
	if confirmationGot != confirmationWant {
		t.Errorf("Expected confirmation number to be [%s], but got [%s]", confirmationWant, confirmationGot)
	}
}

func TestBadInput_NoCheckNumber(t *testing.T) {
	_, err := DepositCheck("", 100)
	assertError(t, err)
}

func TestBadInput_ZeroAmount(t *testing.T) {
	_, err := PerformDeposit("12345", 0)
	assertError(t, err)
}

func TestBadInput_NegativeAmount(t *testing.T) {
	_, err := PerformDeposit("12345", -100)
	assertError(t, err)
}

func TestSpecialistReview(t *testing.T) {
	confirmationNumber, err := PerformDeposit("X12345", 100001)
	assertNoError(t, err)
	assertConfirmationNumber(t, confirmationNumber, "FAKECONF-X12345")
	assertMapContains(t, specialistReviews, "X12345")
}

func TestFBIReview(t *testing.T) {
	confirmationNumber, err := PerformDeposit("Y12345", 1000001)
	assertNoError(t, err)
	assertConfirmationNumber(t, confirmationNumber, "FAKECONF-Y12345")
	assertMapContains(t, fbiReviews, "Y12345")
}

func assertMapContains(t *testing.T, m map[string]float64, key string) {
	if _, ok := m[key]; !ok {
		t.Error("missing checknumber in review")
	}
}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Error("Expected an error")
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("Did not expect an error")
	}
}
