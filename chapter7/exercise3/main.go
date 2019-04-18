package shapes

import "math"

type Circle struct {
	x, y, r float64
}

func (c *Circle) perimiter() float64 {
	return 2 * math.Pi * c.r
}

// Too lazy to do Rectangle.
