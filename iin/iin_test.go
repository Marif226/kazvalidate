package iin

import "testing"

type iinTest struct {
	iin      string
	expected bool
}

func TestIIN(t *testing.T) {
	tests := []iinTest{
		// Valid IINs
		{"901209300017", true},
		{"890801400014", true},
		{"041011500012", true}, // Leap year
		{"110204500014", true},
		{"880526400018", true},

		// Invalid IINs
		{"99010130012", false},   // Less than 12 digits
		{"9901013001234", false}, // More than 12 digits
		{"990131300123", false},  // Invalid date (31 February)
		{"990101700123", false},  // Invalid 7th digit
		{"99010130012A", false},  // Contains non-digit character
		{"990101300124", false},  // Invalid control digit
		{"000230500034", false},  // Invalid date (30 February)
		{"991231400056", false},  // Incorrect date or control digit
	}

	for _, test := range tests {
		result, err := Validate(test.iin)
		if result != test.expected {
			t.Errorf("validate.IIN(%s) = %v; expected %v; err: %v", test.iin, result, test.expected, err)
		}
	}
}
