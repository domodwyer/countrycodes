package countrycodes_test

import (
	"fmt"
	"testing"

	"github.com/domodwyer/country-codes"
)

func TestToName(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"GB", "United Kingdom"},
		{"TD", "Chad"},
		{"DJ", "Djibouti"},
		{"KP", "Korea, Democratic People's Republic Of"},
		{"WS", "Samoa"},
		{"TO", "Tonga"},
	}

	for _, test := range tests {
		actual, err := countrycodes.ToName(test.in)
		if err != nil {
			t.Error("Unexpected error")
		}

		if test.expected != actual {
			t.Error("Failed to match expected: ", test.expected, " got: ", actual)
		}
	}
}

func TestToName_Error(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"AA", "AA"},
	}

	for _, test := range tests {
		actual, err := countrycodes.ToName(test.in)
		if test.expected != actual {
			t.Error("Failed to match expected: ", test.expected, " got: ", actual)
		}

		if err == nil {
			t.Error("Expected error wasn't returned")
		}
	}
}

func TestToCode(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"United Kingdom", "GB"},
		{"Chad", "TD"},
		{"Djibouti", "DJ"},
		{"Korea, Democratic People's Republic Of", "KP"},
		{"Samoa", "WS"},
		{"Tonga", "TO"},
	}

	for _, test := range tests {
		actual, err := countrycodes.ToCode(test.in)
		if err != nil {
			t.Error("Unexpected error")
		}

		if test.expected != actual {
			t.Error("Failed to match expected: ", test.expected, " got: ", actual)
		}
	}
}

func TestToCode_Error(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"AA", "AA"},
	}

	for _, test := range tests {
		actual, err := countrycodes.ToCode(test.in)
		if test.expected != actual {
			t.Error("Failed to match expected: ", test.expected, " got: ", actual)
		}

		if err == nil {
			t.Error("Expected error wasn't returned")
		}
	}
}

func ExampleToName() {
	name, err := countrycodes.ToName("GB")
	if err != nil {
		panic("Country code not found :( ")
	}

	fmt.Printf("GB is the country code for %s", name)
	// Output: GB is the country code for United Kingdom
}

func ExampleToCode() {
	code, err := countrycodes.ToName("United Kingdom")
	if err != nil {
		panic("Country name not found :( ")
	}

	fmt.Printf("United Kingdom's country code is %s", code)
	// Output: United Kingdom's country code is GB
}
