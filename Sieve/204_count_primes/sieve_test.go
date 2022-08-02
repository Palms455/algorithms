package sieve

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		name          string
		input, output int
	}{
		{
			name:   "For 10",
			input:  10,
			output: 4,
		},
		{
			name:   "For 0",
			input:  0,
			output: 0,
		},
		{
			name:   "For 1",
			input:  1,
			output: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := countPrimes(test.input); got != test.output {
				t.Errorf("CountPrimes() = %v, want %v", got, test.output)
			}
		})
	}
}
