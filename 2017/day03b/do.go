package main

import (
	"fmt"
	"math"
)

// Do solves day03b.
func Do(int) (int, error) {
	return 0, nil
}

// Coord represents a coordinate position in an spiral memory.
type Coord struct {
	X, Y int
}

// String pretty prints a coord.
func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

// PosToCoord returns the spiral coordinate of a memory position.
// Returns an error with negative positions.
func PosToCoord(p int) (Coord, error) {
	if p == 0 {
		return Coord{X: 0, Y: 0}, nil
	}
	side, err := RingSidePos(p)
	if err != nil {
		return Coord{}, fmt.Errorf(
			"calculating ring side: negative position (%d)", p)
	}
	halfSide := (side - 1) / 2
	corners := cornersPos(side)
	switch p {
	case corners[0]:
		return Coord{X: halfSide, Y: halfSide}, nil
	case corners[1]:
		return Coord{X: -halfSide, Y: halfSide}, nil
	case corners[2]:
		return Coord{X: -halfSide, Y: -halfSide}, nil
	case corners[3]:
		return Coord{X: halfSide, Y: -halfSide}, nil
	}
	zone, err := PosToZone(p)
	if err != nil {
		return Coord{}, fmt.Errorf("calculating zone: %v", err)
	}
	switch zone {
	case Top:
		dist := (corners[1] - p)
		return Coord{X: -halfSide + dist, Y: halfSide}, nil
	case Left:
		dist := (corners[2] - p)
		return Coord{X: -halfSide, Y: -halfSide + dist}, nil
	case Bottom:
		dist := (corners[3] - p)
		return Coord{X: halfSide - dist, Y: -halfSide}, nil
	case Right:
		dist := (corners[0] - p)
		return Coord{X: halfSide, Y: halfSide - dist}, nil
	default:
		panic(fmt.Sprintf("unknown zone: %d", zone))
	}
}

func cornersPos(side int) []int {
	br := (side * side) - 1
	bl := br - (side - 1)
	tl := bl - (side - 1)
	tr := tl - (side - 1)
	return []int{tr, tl, bl, br}
}

// RingSidePos returns the side of the ring containing the given position.
// Returns an error with negative positions.
func RingSidePos(p int) (int, error) {
	if p < 0 {
		return 0, fmt.Errorf("negative position (%d)", p)
	}
	if p == 0 {
		return 1, nil
	}
	sqrt := int(math.Sqrt(float64(p)))
	if sqrt%2 == 0 {
		sqrt--
	}
	return sqrt + 2, nil
}

// CoordToPos returns the memory position of an spiral coordinate.
func CoordToPos(c Coord) int {
	zero := Coord{X: 0, Y: 0}
	if c == zero {
		return 0
	}

	side := ringSideFromCoord(c)
	corners := cornersCoord(side)

	// Positions of the corners
	br := (side * side) - 1
	bl := br - (side - 1)
	tl := bl - (side - 1)
	tr := tl - (side - 1)

	// Solve if c is a corner
	switch c {
	case corners[0]:
		return tr
	case corners[1]:
		return tl
	case corners[2]:
		return bl
	case corners[3]:
		return br
	}

	// Solve if c is in a side
	zone := coordToZone(c)
	switch zone {
	case Bottom:
		d := corners[3].X - c.X
		return br - d
	case Left:
		d := c.Y - corners[2].Y
		return bl - d
	case Top:
		d := c.X - corners[1].X
		return tl - d
	case Right:
		d := corners[0].Y - c.Y
		return tr - d
	default:
		panic(fmt.Sprintf("unknown zone: %d", zone))
	}
}

// TODO
func ringSideFromCoord(c Coord) int {
	return 5
}

// TODO
func cornersCoord(side int) []Coord {
	return []Coord{
		Coord{X: 2, Y: 2},
		Coord{X: -2, Y: 2},
		Coord{X: -2, Y: -2},
		Coord{X: 2, Y: -2},
	}
}

// TODO
func coordToZone(c Coord) Zone {
	return Top
}

// Zone defines the four regions of a ring.
type Zone int

const (
	// Top is the top row.
	Top Zone = iota
	// Bottom is the bottom row.
	Bottom
	// Left is the left column.
	Left
	// Right is the right column.
	Right
)

// String pretty prints a Zone.
func (z Zone) String() string {
	switch z {
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		panic(fmt.Sprintf("unknown zone (%d)", int(z)))
	}
}

// PosToZone returns the ring's zone for a given position (p).
// Returns an error if p is negative.
// Returns an indetermined value if p is located in more than one zone, for
// instance when p=0 or when p is located in a corner of the ring.
func PosToZone(p int) (Zone, error) {
	side, err := RingSidePos(p)
	if err != nil {
		return Top, fmt.Errorf(
			"calculating ring side: negative position (%d)", p)
	}
	corners := cornersPos(side)
	switch {
	case p <= corners[0]:
		return Right, nil
	case p <= corners[1]:
		return Top, nil
	case p <= corners[2]:
		return Left, nil
	case p <= corners[3]:
		return Bottom, nil
	default:
		panic("internal error")
	}
}
