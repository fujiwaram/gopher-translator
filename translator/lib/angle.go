package lib

// Angle .
type Angle float64

// RightDiff .
func (src Angle) RightDiff(dst Angle) Angle {
	if src >= dst {
		return src - dst
	}
	// src < dst
	return 360 + (src - dst)
}
