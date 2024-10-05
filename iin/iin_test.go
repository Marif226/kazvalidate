package iin

import "testing"

// Тестовая структура для ИИН
type iinTest struct {
	iin      string
	expected bool
}

// Тестовая функция для validate.IIN
func TestIIN(t *testing.T) {
	tests := []iinTest{
		// Валидные ИИН
		{"901209300017", true},
		{"890801400014", true},
		{"041011500012", true}, // Високосный год
		{"110204500014", true},
		{"880526400018", true},

		// Невалидные ИИН
		{"99010130012", false},   // Меньше 12 цифр
		{"9901013001234", false}, // Больше 12 цифр
		{"990131300123", false},  // Неверная дата (31 февраля)
		{"990101700123", false},  // Неверная 7-я цифра
		{"99010130012A", false},  // Наличие нецифрового символа
		{"990101300124", false},  // Неверная контрольная цифра
		{"000230500034", false},  // Неверная дата (30 февраля)
		{"991231400056", false},  // Некорректная дата или контрольная цифра
	}

	for _, test := range tests {
		result, err := Validate(test.iin)
		if result != test.expected {
			t.Errorf("validate.IIN(%s) = %v; expected %v; err: %v", test.iin, result, test.expected, err)
		}
	}
}
