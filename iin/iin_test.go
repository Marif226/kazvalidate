package iin

import "testing"

type iinTest struct {
	iin      string
	expected bool
	wantErr  bool
}

func TestIIN(t *testing.T) {
	tests := []iinTest{
		// Valid IINs
		{"901209300017", true, false},
		{"890801400014", true, false},
		{"041011500012", true, false}, // Leap year
		{"110204500014", true, false},
		{"880526400018", true, false},

		// Invalid IINs
		{"99010130012", false, true},   // Less than 12 digits
		{"9901013001234", false, true}, // More than 12 digits
		{"990131300123", false, true},  // Invalid date (31 February)
		{"990101700123", false, true},  // Invalid 7th digit
		{"99010130012A", false, true},  // Contains non-digit character
		{"990101300124", false, true},  // Invalid control digit
		{"000230500034", false, true},  // Invalid date (30 February)
		{"991231400056", false, true},  // Incorrect date or control digit
	}

	for _, test := range tests {
		t.Run(test.iin, func(t *testing.T) {
			result, err := Validate(test.iin)
			if result != test.expected {
				t.Errorf("Validate(%s) = %v; expected %v; err: %v", test.iin, result, test.expected, err)
			}
			if (err != nil) != test.wantErr {
				t.Errorf("Validate(%s) error = %v; wantErr %v", test.iin, err, test.wantErr)
			}
		})
	}
}
