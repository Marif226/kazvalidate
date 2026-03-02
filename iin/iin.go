// Package iin provides validation for Kazakhstan Individual Identification Numbers (IIN).
package iin

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Validate reports whether iin is a valid 12-digit Kazakhstan IIN.
//
// It validates the length and numeric format, checks whether the encoded birth
// date is valid, and verifies the checksum according to the IIN algorithm.
func Validate(iin string) (bool, error) {
	// Check IIN length.
	if len(iin) != 12 {
		return false, errors.New("iin must contain 12 digits")
	}

	// Check that all characters are digits.
	for _, char := range iin {
		if char < '0' || char > '9' {
			return false, errors.New("iin must contain only digits")
		}
	}

	// Parse birth date from the first 6 digits.
	year, err := strconv.Atoi(iin[0:2])
	if err != nil {
		return false, errors.New("invalid year format")
	}
	month, err := strconv.Atoi(iin[2:4])
	if err != nil {
		return false, errors.New("invalid month format")
	}
	day, err := strconv.Atoi(iin[4:6])
	if err != nil {
		return false, errors.New("invalid day format")
	}

	// Determine century from the 7th digit.
	century, err := strconv.Atoi(string(iin[6]))
	if err != nil || century < 1 || century > 6 {
		return false, errors.New("invalid format of 7th digit (century)")
	}

	switch century {
	case 1, 2:
		year += 1900
	case 3, 4:
		year += 2000
	}

	// Validate the parsed date.
	if _, err := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-%02d", year, month, day)); err != nil {
		return false, errors.New("invalid date")
	}

	// Validate checksum digit.
	return controlDigit(iin)
}

// controlDigit validates the IIN checksum digit.
func controlDigit(iin string) (bool, error) {
	sum := 0

	// Compute weighted sum using the first coefficient sequence.
	for i := range 11 {
		digit, err := strconv.Atoi(string(iin[i]))
		if err != nil {
			return false, err
		}
		sum += digit * (i + 1)
	}

	// Calculate checksum digit.
	controlDigit := sum % 11

	// If checksum is 10, recalculate using the second coefficient sequence.
	if controlDigit == 10 {
		sum = 0
		for i := range 11 {
			digit, err := strconv.Atoi(string(iin[i]))
			if err != nil {
				return false, err
			}
			t := (i + 3) % 11
			if t == 0 {
				t = 11
			}
			sum += digit * t
		}
		controlDigit = sum % 11

		// If checksum is still 10, the IIN is invalid.
		if controlDigit == 10 {
			return false, errors.New("invalid control digit")
		}
	}

	// Compare calculated checksum with the last IIN digit.
	lastDigit, err := strconv.Atoi(string(iin[11]))
	if err != nil {
		return false, err
	}
	if controlDigit == lastDigit {
		return true, nil
	} else {
		return false, errors.New("invalid control digit")
	}
}
