package main

import "testing"

func TestPart01(t *testing.T) {
	t.Run("examples", func(t *testing.T) {
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
			value := part01(test.moons, test.steps)
			if value != test.expected {
				t.Errorf("expected %d, got %d", test.expected, value)
			}
		}
	})

	t.Run("actual", func(t *testing.T) {
		moons := []moon{
			{posX: 1, posY: 3, posZ: -11},
			{posX: 17, posY: -10, posZ: -8},
			{posX: -1, posY: -15, posZ: 2},
			{posX: 12, posY: -4, posZ: -4},
		}

		expected := 8310
		got := part01(moons, 1000)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("examples", func(t *testing.T) {

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
			value := part02(test.moons)
			if value != test.expected {
				t.Errorf("expected %d, got %d", test.expected, value)
			}
		}
	})

	t.Run("actual", func(t *testing.T) {
		moons := []moon{
			{posX: 1, posY: 3, posZ: -11},
			{posX: 17, posY: -10, posZ: -8},
			{posX: -1, posY: -15, posZ: 2},
			{posX: 12, posY: -4, posZ: -4},
		}

		expected := 319290382980408
		got := part02(moons)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
