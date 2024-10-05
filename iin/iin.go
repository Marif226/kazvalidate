package iin

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Validate проверяет ИИН на валидность
func Validate(iin string) (bool, error) {
	// Проверка длины ИИН
	if len(iin) != 12 {
		return false, errors.New("iin must contain 12 digits")
	}

	// Проверка, что все символы — цифры
	for _, char := range iin {
		if char < '0' || char > '9' {
			return false, errors.New("iin must contain only digits")
		}
	}

	// Проверка даты рождения (первые 6 цифр)
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

	// Определяем столетие по 7-й цифре
	century, err := strconv.Atoi(string(iin[6]))
	if err != nil || century < 1 || century > 6 {
		return false, errors.New("invalid format of 7th digit (century)")
	}

	if century == 1 || century == 2 {
		year += 1900
	} else if century == 3 || century == 4 {
		year += 2000
	}

	// Проверяем корректность даты
	if _, err := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-%02d", year, month, day)); err != nil {
		return false, errors.New("invalid date")
	}

	// Проверка контрольной цифры
	return controlDigit(iin)
}

// controlDigit проверяет контрольную цифру ИИН
func controlDigit(iin string) (bool, error) {
	sum := 0

	// Вычисляем сумму произведений цифр ИИН на соответствующие коэффициенты
	for i := 0; i < 11; i++ {
		digit, err := strconv.Atoi(string(iin[i]))
		if err != nil {
			return false, err
		}
		sum += digit * (i + 1)
	}

	// Контрольная цифра
	controlDigit := sum % 11

	// Если контрольная цифра больше 9, пересчитываем по другому алгоритму
	if controlDigit == 10 {
		sum = 0
		for i := 0; i < 11; i++ {
			digit, _ := strconv.Atoi(string(iin[i]))
			t := (i + 3) % 11
			if t == 0 {
				t = 11
			}
			sum += digit * t
		}
		controlDigit = sum % 11

		// Если контрольная цифра все равно больше 9, то она равна 0
		if controlDigit == 10 {
			return false, errors.New("invalid control digit")
		}
	}

	// Сравниваем контрольная цифру с последней цифрой ИИН (контрольный разряд)
	lastDigit, _ := strconv.Atoi(string(iin[11]))
	if controlDigit == lastDigit {
		return true, nil
	} else {
		return false, errors.New("invalid control digit")
	}
}
