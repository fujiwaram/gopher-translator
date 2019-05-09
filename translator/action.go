package translator

import (
	"github.com/fujiwaram/gopherTranslator/consts"
	"github.com/fujiwaram/gopherTranslator/translator/lib"
)

type action string

const (
	moveTo    action = "moveTo"
	lineTo    action = "lineTo"
	curveTo   action = "curveTo"
	closePath action = "closePath"

	none action = ""
)

var (
	colorOrderOff, colorOrderSet lib.Order
)

func init() {
	colorOrderOff.Printf("%s %s\n", consts.OrderColor, "off")
	colorOrderSet.Printf("%s %s\n", consts.OrderColor, consts.LineColor)

}

func parseAction(data interface{}) action {
	return action(data.(string))
}

func (a action) isChunkTail() bool {
	if a == closePath {
		return true
	}
	return false
}

func (a action) isOnePoint() bool {
	switch a {
	case moveTo, lineTo, closePath:
		return true
	}
	return false
}

func (a action) translate(prev action) lib.Order {
	switch a {
	case moveTo:
		return colorOrderOff
	case lineTo, curveTo, closePath:
		if prev == moveTo {
			return colorOrderSet
		}
	}
	return lib.Order{}
}
