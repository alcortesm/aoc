package memory_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day03b/memory"
)

var posCoordTable = []struct {
	pos   int
	coord memory.Coord
}{
	{pos: 0, coord: memory.Coord{X: 0, Y: 0}},
	{pos: 1, coord: memory.Coord{X: 1, Y: 0}},
	{pos: 2, coord: memory.Coord{X: 1, Y: 1}},
	{pos: 3, coord: memory.Coord{X: 0, Y: 1}},
	{pos: 4, coord: memory.Coord{X: -1, Y: 1}},
	{pos: 5, coord: memory.Coord{X: -1, Y: 0}},
	{pos: 6, coord: memory.Coord{X: -1, Y: -1}},
	{pos: 7, coord: memory.Coord{X: 0, Y: -1}},
	{pos: 8, coord: memory.Coord{X: 1, Y: -1}},
	{pos: 9, coord: memory.Coord{X: 2, Y: -1}},
	{pos: 10, coord: memory.Coord{X: 2, Y: 0}},
	{pos: 11, coord: memory.Coord{X: 2, Y: 1}},
	{pos: 12, coord: memory.Coord{X: 2, Y: 2}},
	{pos: 13, coord: memory.Coord{X: 1, Y: 2}},
	{pos: 14, coord: memory.Coord{X: 0, Y: 2}},
	{pos: 15, coord: memory.Coord{X: -1, Y: 2}},
	{pos: 16, coord: memory.Coord{X: -2, Y: 2}},
	{pos: 17, coord: memory.Coord{X: -2, Y: 1}},
	{pos: 18, coord: memory.Coord{X: -2, Y: 0}},
	{pos: 19, coord: memory.Coord{X: -2, Y: -1}},
	{pos: 20, coord: memory.Coord{X: -2, Y: -2}},
	{pos: 21, coord: memory.Coord{X: -1, Y: -2}},
	{pos: 22, coord: memory.Coord{X: 0, Y: -2}},
	{pos: 23, coord: memory.Coord{X: 1, Y: -2}},
	{pos: 24, coord: memory.Coord{X: 2, Y: -2}},
	{pos: 25, coord: memory.Coord{X: 3, Y: -2}},
	{pos: 26, coord: memory.Coord{X: 3, Y: -1}},
	{pos: 27, coord: memory.Coord{X: 3, Y: 0}},
	{pos: 28, coord: memory.Coord{X: 3, Y: 1}},
	{pos: 29, coord: memory.Coord{X: 3, Y: 2}},
	{pos: 30, coord: memory.Coord{X: 3, Y: 3}},
	{pos: 31, coord: memory.Coord{X: 2, Y: 3}},
	{pos: 32, coord: memory.Coord{X: 1, Y: 3}},
	{pos: 33, coord: memory.Coord{X: 0, Y: 3}},
	{pos: 34, coord: memory.Coord{X: -1, Y: 3}},
	{pos: 35, coord: memory.Coord{X: -2, Y: 3}},
	{pos: 36, coord: memory.Coord{X: -3, Y: 3}},
	{pos: 37, coord: memory.Coord{X: -3, Y: 2}},
	{pos: 38, coord: memory.Coord{X: -3, Y: 1}},
	{pos: 39, coord: memory.Coord{X: -3, Y: 0}},
	{pos: 40, coord: memory.Coord{X: -3, Y: -1}},
	{pos: 41, coord: memory.Coord{X: -3, Y: -2}},
	{pos: 42, coord: memory.Coord{X: -3, Y: -3}},
	{pos: 43, coord: memory.Coord{X: -2, Y: -3}},
	{pos: 44, coord: memory.Coord{X: -1, Y: -3}},
	{pos: 45, coord: memory.Coord{X: 0, Y: -3}},
	{pos: 46, coord: memory.Coord{X: 1, Y: -3}},
	{pos: 47, coord: memory.Coord{X: 2, Y: -3}},
	{pos: 48, coord: memory.Coord{X: 3, Y: -3}},
	{pos: 49, coord: memory.Coord{X: 4, Y: -3}},
}

func TestPosToCoordOK(t *testing.T) {
	t.Parallel()
	for _, tt := range posCoordTable {
		tt := tt
		description := fmt.Sprint(tt.pos)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := memory.PosToCoord(tt.pos)
			if err != nil {
				t.Errorf("unexpected error (%v)", err)
			}
			if tt.coord != got {
				t.Errorf("want=%d, got=%d", tt.coord, got)
			}
		})
	}
}

func TestPosToCoordError(t *testing.T) {
	t.Parallel()
	for _, pos := range []int{
		-1, -2, -1000,
	} {
		pos := pos
		t.Run(fmt.Sprint(pos), func(t *testing.T) {
			t.Parallel()
			if _, err := memory.PosToCoord(pos); err == nil {
				t.Errorf("unexpected success (pos=%d)", pos)
			}
		})
	}
}

func TestCoordToPos(t *testing.T) {
	t.Parallel()
	for _, tt := range posCoordTable {
		tt := tt
		description := fmt.Sprint(tt.coord)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got := memory.CoordToPos(tt.coord)
			if tt.pos != got {
				t.Errorf("want=%d, got=%d", tt.pos, got)
			}
		})
	}
}

func TestRingSidePosOK(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		input int
		want  int
	}{
		{input: 0, want: 1},
		{input: 1, want: 3},
		{input: 2, want: 3},
		{input: 3, want: 3},
		{input: 4, want: 3},
		{input: 5, want: 3},
		{input: 6, want: 3},
		{input: 7, want: 3},
		{input: 8, want: 3},
		{input: 9, want: 5},
		{input: 10, want: 5},
		{input: 11, want: 5},
		{input: 12, want: 5},
		{input: 13, want: 5},
		{input: 14, want: 5},
		{input: 15, want: 5},
		{input: 16, want: 5},
		{input: 17, want: 5},
		{input: 18, want: 5},
		{input: 19, want: 5},
		{input: 20, want: 5},
		{input: 21, want: 5},
		{input: 22, want: 5},
		{input: 23, want: 5},
		{input: 24, want: 5},
		{input: 25, want: 7},
		{input: 48, want: 7},
		{input: 49, want: 9},
		{input: 50, want: 9},
	} {
		tt := tt
		description := fmt.Sprint(tt.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := memory.RingSidePos(tt.input)
			if err != nil {
				t.Errorf("unexpected error (%v)", err)
			}
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}

func TestRingSidePosError(t *testing.T) {
	t.Parallel()
	for _, pos := range []int{
		-1, -2, -1000,
	} {
		pos := pos
		t.Run(fmt.Sprint(pos), func(t *testing.T) {
			t.Parallel()
			if _, err := memory.RingSidePos(pos); err == nil {
				t.Errorf("unexpected success (pos=%d)", pos)
			}
		})
	}
}

func TestPosToZoneOK(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		input int
		want  memory.Zone
	}{
		// 0 returns and indeterminate value
		{input: 1, want: memory.Right},
		// 2 returns and indeterminate value
		{input: 3, want: memory.Top},
		// 4 returns and indeterminate value
		{input: 5, want: memory.Left},
		// 6 returns and indeterminate value
		{input: 7, want: memory.Bottom},
		// 8 returns and indeterminate value
		{input: 9, want: memory.Right},
		{input: 10, want: memory.Right},
		{input: 11, want: memory.Right},
		// 12 returns and indeterminate value
		{input: 13, want: memory.Top},
		{input: 14, want: memory.Top},
		{input: 15, want: memory.Top},
		// 16 returns and indeterminate value
		{input: 17, want: memory.Left},
		{input: 18, want: memory.Left},
		{input: 19, want: memory.Left},
		// 20 returns and indeterminate value
		{input: 21, want: memory.Bottom},
		{input: 22, want: memory.Bottom},
		{input: 23, want: memory.Bottom},
		// 24 returns and indeterminate value
		{input: 25, want: memory.Right},
		{input: 49, want: memory.Right},
		{input: 50, want: memory.Right},
	} {
		tt := tt
		description := fmt.Sprint(tt.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := memory.PosToZone(tt.input)
			if err != nil {
				t.Errorf("unexpected error (%v)", err)
			}
			if tt.want != got {
				t.Errorf("want=%v, got=%v", tt.want, got)
			}
		})
	}
}

func TestPosToZoneError(t *testing.T) {
	t.Parallel()
	for _, input := range []int{
		-1, -10, -100,
	} {
		input := input
		description := fmt.Sprint(input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			_, err := memory.PosToZone(input)
			if err == nil {
				t.Error("unexpected success")
			}
		})
	}
}
