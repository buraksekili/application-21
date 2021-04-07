package task2

import (
	"fmt"
	"regexp"
	"strings"
)

// checkPassw checks validation of the given password 
func checkPassw(passw string) bool {
	passw = strings.TrimSpace(passw)
	if len(passw) < 8 {
		return false
	}

	digits := `[0-9]{1}`
	lowers := `[a-z]{1}`
	uppers := `[A-Z]{1}`
	specials := `[!@#$%^&*()-+]{1}`
	if m, err := regexp.MatchString(digits, passw); !m || err != nil {
		return false
	}
	if m, err := regexp.MatchString(lowers, passw); !m || err != nil {
		return false
	}
	if m, err := regexp.MatchString(uppers, passw); !m || err != nil {
		return false
	}
	if m, err := regexp.MatchString(specials, passw); !m || err != nil {
		return false
	}
	return true
}

// Task2 is the solution for the question2
func Task2() {

	testCases := []struct {
		name     string
		passw    string
		expected bool
	}{
		{
			name:     "valid password",
			passw:    "1lU!5678",
			expected: true,
		},
		{
			name:     "invalid password, empty password (len)",
			passw:    "         ",
			expected: false,
		},
		{
			name:     "invalid password (len)",
			passw:    "1lU!",
			expected: false,
		},
		{
			name:     "invalid password with empty chars (len)",
			passw:    "1lU!         ",
			expected: false,
		},
		{
			name:     "invalid password (upper)",
			passw:    "1la!5678",
			expected: false,
		},
		{
			name:     "invalid password (lower)",
			passw:    "1LL!5678",
			expected: false,
		},
		{
			name:     "invalid password (spec)",
			passw:    "1LL45678",
			expected: false,
		},
		{
			name:     "invalid password (digit)",
			passw:    "Burak!#burak",
			expected: false,
		},
		{
			name:     "invalid password (lower, digit, spec, len)",
			passw:    "BURAK",
			expected: false,
		},
		{
			name:     "invalid password (spec, len)",
			passw:    "1lU",
			expected: false,
		},
		{
			name:     "invalid password (spec, len)",
			passw:    "1lU",
			expected: false,
		},
		{
			name:     "invalid password (spec, len)",
			passw:    "1la!ac",
			expected: false,
		},
	}

	for _, tc := range testCases {
		result := checkPassw(tc.passw)
		if result != tc.expected {
			fmt.Printf("FAILED! %s: passw: %s got %v, expected %v\n", tc.name, tc.passw, result, tc.expected)
		} else {
			fmt.Printf("PASSED %s: got %v, expected %v\n", tc.name, result, tc.expected)
		}
	}
}
