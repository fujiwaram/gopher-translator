package translator

import (
	"github.com/fujiwaram/gopherTranslator/consts"
	"github.com/fujiwaram/gopherTranslator/translator/lib"
)

type actionPoint struct {
	action action
	points points
}

func parseActionPoint(data interface{}) actionPoint {
	dataActionPoint := data.([]interface{})
	resultAction := parseAction(dataActionPoint[0])
	resultPoints := parsePoints(dataActionPoint[1])
	return actionPoint{action: resultAction, points: resultPoints}
}

func newActionPoint(action action, point point) actionPoint {
	return actionPoint{action: action, points: points{point}}
}

func (ap actionPoint) adjust(margin point) actionPoint {
	newPoints := ap.points.adjust(margin)
	return actionPoint{action: ap.action, points: newPoints}
}

func (ap actionPoint) getMaxX() float64 {
	return ap.points.getMaxX()
}

func (ap actionPoint) getMaxXY() point {
	return ap.points.getMaxXY()
}

func (ap actionPoint) getMinXY() point {
	return ap.points.getMinXY()
}

func (ap actionPoint) getEndAngle(prevPoint point) lib.Angle {
	if ap.action.isOnePoint() {
		return prevPoint.angle(ap.points[0])
	}
	return ap.points.getEndAngle()
}

func (ap actionPoint) getStartPoint() point {
	return ap.points.getStartPoint()
}

func (ap actionPoint) getEndPoint() point {
	return ap.points.getEndPoint()
}

var pointsHandler = map[action]actionPointsHandler{
	moveTo:    lineToHandler{},
	lineTo:    lineToHandler{},
	curveTo:   curveToHandler{},
	closePath: lineToHandler{},
}

type actionPointsHandler interface {
	translate(prevAngle lib.Angle, prevPoint point, nextPoints points) *lib.Order
}
type lineToHandler struct{}
type curveToHandler struct{}

//
// translate
//

func (ap actionPoint) translate(prevAngle lib.Angle, prevAp actionPoint) *lib.Order {
	translatedAction := ap.action.translate(prevAp.action)
	translatedPoints := pointsHandler[ap.action].translate(prevAngle, prevAp.getEndPoint(), ap.points)
	translatedAction.Append(*translatedPoints)
	return &translatedAction
}

func lineTranslate(o *lib.Order, prevAngle lib.Angle, prevPoint, nextPoint point) {
	ang := prevPoint.angle(nextPoint)
	turn := prevAngle.RightDiff(ang)
	o.Printf("%s %f\n", consts.OrderTurn, turn)

	l := prevPoint.length(nextPoint)
	o.Printf("%s %f\n", consts.OrderMove, l)
}

func (h lineToHandler) translate(prevAngle lib.Angle, prevPoint point, nextPoints points) *lib.Order {
	var o lib.Order

	lineTranslate(&o, prevAngle, prevPoint, nextPoints[0])

	return &o
}

func (h curveToHandler) translate(prevAngle lib.Angle, prevPoint point, nextPoints points) *lib.Order {
	var o lib.Order

	currentAngle := prevAngle
	currentPoint := prevPoint
	var nextPoint point

	for i := 0; i < len(nextPoints); i++ {
		nextPoint = nextPoints[i]
		// TODO: calculate bezier curve points
		lineTranslate(&o, currentAngle, currentPoint, nextPoint)

		currentAngle = currentPoint.angle(nextPoints[i])
		currentPoint = nextPoint
	}

	return &o
}
