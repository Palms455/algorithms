package EulerFunction

import "testing"

func TestEuler(t *testing.T) {
	tests := []struct {
		name         string
		nums, eulers int
	}{
		{
			name:   "Euler 9",
			nums:   9,
			eulers: 6,
		},
		{
			name:   "Euler 36",
			nums:   36,
			eulers: 12,
		},
		{
			name:   "Euler 84",
			nums:   84,
			eulers: 24,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualSlow := slowEuler(test.nums)
			actualFast := fastEuler(test.nums)
			if actualSlow != test.eulers {
				t.Errorf("test slow %s failed\n", test.name)
				t.Errorf("expected: %d, actual: %d", test.eulers, actualSlow)
			}
			if actualFast != test.eulers {
				t.Errorf("test fast %s failed\n", test.name)
				t.Errorf("expected: %d, actual: %d", test.eulers, actualSlow)
			}
		})
	}

}
