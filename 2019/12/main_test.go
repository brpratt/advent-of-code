package day12

import "testing"

func TestSolvePart01(t *testing.T) {
	tests := []struct {
		moons    []moon
		steps    int
		expected int
	}{
		{
			[]moon{
				{posX: -1, posY: 0, posZ: 2},
				{posX: 2, posY: -10, posZ: -7},
				{posX: 4, posY: -8, posZ: 8},
				{posX: 3, posY: 5, posZ: -1},
			},
			10,
			179,
		},
		{
			[]moon{
				{posX: -8, posY: -10, posZ: 0},
				{posX: 5, posY: 5, posZ: 10},
				{posX: 2, posY: -7, posZ: 3},
				{posX: 9, posY: -8, posZ: -3},
			},
			100,
			1940,
		},
	}

	for _, test := range tests {
		value := SolvePart01(test.moons, test.steps)
		if value != test.expected {
			t.Errorf("expected %d, got %d", test.expected, value)
		}
	}
}

func TestSolvePart02(t *testing.T) {
	tests := []struct {
		moons    []moon
		expected int
	}{
		{
			[]moon{
				{posX: -1, posY: 0, posZ: 2},
				{posX: 2, posY: -10, posZ: -7},
				{posX: 4, posY: -8, posZ: 8},
				{posX: 3, posY: 5, posZ: -1},
			},
			2772,
		},
		{
			[]moon{
				{posX: -8, posY: -10, posZ: 0},
				{posX: 5, posY: 5, posZ: 10},
				{posX: 2, posY: -7, posZ: 3},
				{posX: 9, posY: -8, posZ: -3},
			},
			4686774924,
		},
	}

	for _, test := range tests {
		value := SolvePart02(test.moons)
		if value != test.expected {
			t.Errorf("expected %d, got %d", test.expected, value)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 8310},
		{2, 319290382980408},
	}

	for _, test := range tests {
		got := Solve(test.part)
		if got != test.expected {
			t.Errorf("failed part %d: expected %d, got %d", test.part, test.expected, got)
		}
	}
}
