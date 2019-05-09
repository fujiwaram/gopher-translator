package translator

import (
	"github.com/fujiwaram/gopherTranslator/consts"
	"github.com/fujiwaram/gopherTranslator/translator/lib"
)

type point lib.Point

func parsePoint(data interface{}) point {
	dataPoint := data.([]interface{})
	if len(dataPoint) < 2 {
		return point{}
	}
	return point{X: dataPoint[0].(float64), Y: dataPoint[1].(float64)}
}

func (p point) adjust(margin point) point {
	return point{X: p.X + margin.X, Y: p.Y + margin.Y}
}

func (p point) angle(dst point) lib.Angle {
	a := lib.AnglePointToPoint(lib.Point(p), lib.Point(dst))
	return lib.Angle(a)
}

func (p point) length(dst point) lib.Length {
	l := lib.LengthPointToPoint(lib.Point(p), lib.Point(dst))
	return lib.Length(l)
}

func (p point) shrink() point {
	return point{X: p.X / consts.Shrink, Y: p.Y / consts.Shrink}
}
