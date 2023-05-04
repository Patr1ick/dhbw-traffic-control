package model

import "testing"

type Case struct {
	name   string
	pos    Coordinate
	expect bool
}

func TestPositionValid(t *testing.T) {
	settings := Settings{
		Width:           1000,
		Height:          1000,
		Depth:           2,
		DatabaseAddress: nil,
	}

	cases := []Case{
		{
			name: "Outside Boundaries Left",
			pos: Coordinate{
				X: -10,
				Y: 10,
			},
			expect: false,
		},
		{
			name: "Outside Boundaries Top",
			pos: Coordinate{
				X: 10,
				Y: -10,
			},
			expect: false,
		},
		{
			name: "Outside Boundaries Right",
			pos: Coordinate{
				X: 1010,
				Y: 10,
			},
			expect: false,
		},
		{
			name: "Outside Boundaries Bottom",
			pos: Coordinate{
				X: 10,
				Y: -1010,
			},
			expect: false,
		},
		{
			name: "Inside Boundaries",
			pos: Coordinate{
				X: 10,
				Y: 10,
			},
			expect: true,
		},
	}

	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				value := settings.Valid(c.pos)
				if value != c.expect {
					t.Errorf("pos: %v, expected: %v, got: %v", c.pos, c.expect, value)
				}
			},
		)
	}

}
