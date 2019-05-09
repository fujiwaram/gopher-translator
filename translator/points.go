package translator

import "github.com/fujiwaram/gopherTranslator/translator/lib"

type points []point

func parsePoints(data interface{}) points {
	dataPoints := data.([]interface{})
	resultPoints := make(points, 0, len(dataPoints))
	for _, dataPoint := range dataPoints {
		resultPoints = append(resultPoints, parsePoint(dataPoint))
	}
	return resultPoints
}

func (ps points) adjust(margin point) points {
	newPoints := make(points, len(ps))
	for i, p := range ps {
		newPoints[i] = p.shrink()
		newPoints[i] = newPoints[i].adjust(margin)
	}
	return newPoints
}

func (ps points) getMaxX() float64 {
	p := ps.getMaxXY()
	return p.X
}

func (ps points) getMaxXY() point {
	if len(ps) == 0 {
		return point{}
	}
	max := ps[0]
	for _, point := range ps {
		if max.X < point.X {
			max.X = point.X
		}
		if max.Y < point.Y {
			max.Y = point.Y
		}
	}
	return max
}

func (ps points) getMinXY() point {
	if len(ps) == 0 {
		return point{}
	}
	min := ps[0]
	for _, point := range ps {
		if min.X > point.X {
			min.X = point.X
		}
		if min.Y > point.Y {
			min.Y = point.Y
		}
	}
	return min
}

func (ps points) getStartPoint() point {
	l := len(ps)
	if l == 0 {
		return point{}
	}
	return ps[0]
}

func (ps points) getEndPoint() point {
	l := len(ps)
	if l == 0 {
		return point{}
	}
	return ps[l-1]
}

func (ps points) getEndAngle() lib.Angle {
	if len(ps) < 2 {
		return 0
	}
	l := len(ps)
	a := ps[l-2].angle(ps[l-1])
	return a
}

func (ps points) hasPoint() bool {
	return len(ps) >= 1
}
