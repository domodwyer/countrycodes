package countrycodes

import "testing"

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
		actual, err := ToName(test.in)
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
		actual, err := ToName(test.in)
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
		actual, err := ToCode(test.in)
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
		actual, err := ToCode(test.in)
		if test.expected != actual {
			t.Error("Failed to match expected: ", test.expected, " got: ", actual)
		}

		if err == nil {
			t.Error("Expected error wasn't returned")
		}
	}
}
