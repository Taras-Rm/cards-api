package helpers

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var (
	ErrInvalidCardNumberLength = errors.New("invalid card number length")
	ErrInvalidCardNumberSymbol = errors.New("invalid card number symbol")
	ErrInvalidCardNumber       = errors.New("invalid card number")
	ErrInvalidCardExpMonth     = errors.New("invalid card expiration month")
	ErrInvalidCardExpYear      = errors.New("invalid card expiration year")
	ErrExpiredCard             = errors.New("card expired")
)

const (
	minMonth        = 1
	maxMonth        = 12
	maxNumberLength = 19
	minNumberLength = 13
)

func ValidateCardExpirationDate(month, year string) error {
	// check month length and if all digits
	monthReg := regexp.MustCompile("^[0-9]{2}$")
	if !monthReg.MatchString(month) {
		return ErrInvalidCardExpMonth
	}

	cardMonth, err := strconv.Atoi(month)
	if err != nil {
		return err
	}

	// check month value
	if cardMonth < minMonth || cardMonth > maxMonth {
		return ErrInvalidCardExpMonth
	}

	// check year length and if all digits
	yearReg := regexp.MustCompile("^[0-9]{4}$")
	if !yearReg.MatchString(year) {
		return ErrInvalidCardExpYear
	}

	cardYear, err := strconv.Atoi(year)
	if err != nil {
		return err
	}

	currentYear, currentMonth, _ := time.Now().Date()

	// check if card expired
	if cardYear < currentYear || (cardYear == currentYear && cardMonth < int(currentMonth)) {
		return ErrExpiredCard
	}

	return nil
}

func ValidateCardNumber(number string) error {
	// check length
	if len(number) < minNumberLength || len(number) > maxNumberLength {
		return ErrInvalidCardNumberLength
	}

	// chack if all digits
	symbReg := regexp.MustCompile("^[0-9]+$")
	if !symbReg.MatchString(number) {
		return ErrInvalidCardNumberSymbol
	}

	// check number according to luna algorithm
	if !lunaCheck(number) {
		return ErrInvalidCardNumber
	}

	return nil
}

func lunaCheck(number string) bool {
	sum := 0
	isCheckingDigit := false

	for i := len(number) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(number[i]))

		if isCheckingDigit {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		isCheckingDigit = !isCheckingDigit
	}

	return sum%10 == 0
}
