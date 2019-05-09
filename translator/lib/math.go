package lib

import "math"

// Point .
type Point struct {
	X, Y float64
}

// AnglePointToPoint .
func AnglePointToPoint(p1, p2 Point) float64 {
	r := math.Atan2(p2.Y-p1.Y, p2.X-p1.X)
	if r < 0 {
		r = r + 2*math.Pi
	}
	return math.Floor(r * 360 / (2 * math.Pi))
}

// LengthPointToPoint .
func LengthPointToPoint(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
