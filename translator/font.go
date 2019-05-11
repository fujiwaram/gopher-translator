package translator

import (
	"github.com/fujiwaram/gopher-translator/consts"
	"github.com/fujiwaram/gopher-translator/translator/lib"
)

type fontOutline []actionPoint

func parseFontOutline(data interface{}) fontOutline {
	dataActionPoints := data.([]interface{})
	resultActionPoints := make(fontOutline, 0, len(dataActionPoints))
	for _, dataActionPoint := range dataActionPoints {
		ap := parseActionPoint(dataActionPoint)
		resultActionPoints = append(resultActionPoints, ap)
	}
	return resultActionPoints
}

const bitDistance = 0.001

func makeDummyFirstFont(firstPoint point) fontOutline {
	f := make(fontOutline, 2)
	// first lib.angle 90°
	f[0] = newActionPoint(none, point{X: firstPoint.X, Y: firstPoint.Y*-1 - bitDistance})
	f[1] = newActionPoint(none, point{X: firstPoint.X, Y: firstPoint.Y * -1})

	return f
}

func makeDummyLastFont(endPoint point) fontOutline {
	f := make(fontOutline, 2)
	f[0] = newActionPoint(moveTo, endPoint)
	// last lib.angle 0°
	f[1] = newActionPoint(moveTo, point{X: endPoint.X, Y: endPoint.Y - bitDistance})

	return f
}

func (outline fontOutline) adjust(margin point) fontOutline {
	newOutline := make(fontOutline, len(outline))
	for i, ap := range outline {
		if ap.action.isChunkTail() {
			head := newOutline.getChunkHead(i)
			newOutline[i] = newActionPoint(ap.action, head.getStartPoint())
			continue
		}
		newAp := ap.adjust(margin)
		newOutline[i] = newAp
	}
	return newOutline
}

func (outline fontOutline) getMargin() point {
	widthMargin := outline.getMaxX() + consts.Margin
	margin := point{X: widthMargin}
	return margin
}

func (outline fontOutline) getChunkHead(chunkTailIndex int) actionPoint {
	chunkHeadIdx := 0
	for i := chunkTailIndex - 1; i > 0; i-- {
		if outline[i].action.isChunkTail() {
			chunkHeadIdx = i + 1
			break
		}
	}
	return outline[chunkHeadIdx]
}

func (outline fontOutline) getMaxX() float64 {
	p := outline.getMaxXY()
	return p.X
}

func (outline fontOutline) getMaxXY() point {
	if len(outline) == 0 {
		return point{}
	}
	max := outline[0].getMaxXY()
	for _, actionPoint := range outline {
		p := actionPoint.getMaxXY()
		if max.X < p.X {
			max.X = p.X
		}
		if max.Y < p.Y {
			max.Y = p.Y
		}
	}
	return max
}

func (outline fontOutline) getMinXY() point {
	if len(outline) == 0 {
		return point{}
	}
	min := outline[0].getMinXY()
	for _, actionPoint := range outline {
		p := actionPoint.getMinXY()
		if min.X > p.X {
			min.X = p.X
		}
		if min.Y > p.Y {
			min.Y = p.Y
		}
	}
	return min
}

func (outline fontOutline) endIndexPoint(index int) point {
	cnt := 0
	var p point
	for i := len(outline) - 1; i >= 0; i-- {
		if !outline[i].points.hasPoint() {
			continue
		}
		if cnt != index {
			cnt++
			continue
		}
		p = outline[i].points.getEndPoint()
		break
	}

	return p
}

func (outline fontOutline) endActionPoint() actionPoint {
	return outline[len(outline)-1]
}

func (outline fontOutline) endAngle() lib.Angle {
	src := outline.endIndexPoint(1)
	dst := outline.endIndexPoint(0)

	return src.angle(dst)
}

func (outline fontOutline) translate(prevFont fontOutline) (*lib.Order, error) {
	var o lib.Order

	currentAngle := prevFont.endAngle()
	currentActionPoint := prevFont.endActionPoint()
	for _, actionPoint := range outline {
		subO, err := actionPoint.translate(currentAngle, currentActionPoint)
		if err != nil {
			return nil, err
		}
		err = o.Append(*subO)
		if err != nil {
			return nil, err
		}

		currentAngle = actionPoint.getEndAngle(currentActionPoint.getEndPoint())
		currentActionPoint = actionPoint
	}

	return &o, nil
}
