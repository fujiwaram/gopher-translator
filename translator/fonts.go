package translator

import (
	"github.com/fujiwaram/gopher-translator/consts"
	"github.com/fujiwaram/gopher-translator/translator/lib"
)

// Data structure
//
// - FontOutlines
//   - fontOutline
//     - actionPoint
//       - action
//       - points
//         - point

// FontOutlines .
type FontOutlines []fontOutline

// ParseFontOutlines .
func ParseFontOutlines(data interface{}) FontOutlines {
	dataOutlines := data.([]interface{})
	fonts := make(FontOutlines, 0, len(dataOutlines))

	margin := point{}
	for _, dataOutline := range dataOutlines {
		f := parseFontOutline(dataOutline)
		adjustedFont := f.adjust(margin)
		fonts = append(fonts, adjustedFont)

		margin = adjustedFont.getMargin()
	}
	return fonts
}

// Translate .
func (fonts FontOutlines) Translate(message string) (string, error) {
	var o lib.Order

	o.WriteString("draw mode\n")
	// first
	prevFont := makeDummyFirstFont(fonts.getFirstPoint())
	// string
	for _, f := range fonts {
		subO, err := f.translate(prevFont)
		if err != nil {
			return "", err
		}
		if err := o.Append(*subO); err != nil {
			return "", err
		}
		prevFont = f
	}
	// last
	f := makeDummyLastFont(fonts.getEndPoint())
	subO, err := f.translate(prevFont)
	if err != nil {
		return "", err
	}
	if err := o.Append(*subO); err != nil {
		return "", err
	}

	// message
	msgO := newMessage(message).translate()
	if err := o.Append(*msgO); err != nil {
		return "", err
	}

	return o.String(), nil
}

func (fonts FontOutlines) getMaxXY() point {
	if len(fonts) == 0 {
		return point{}
	}
	max := fonts[0].getMaxXY()
	for _, f := range fonts {
		p := f.getMaxXY()
		if max.X < p.X {
			max.X = p.X
		}
		if max.Y < p.Y {
			max.Y = p.Y
		}
	}
	return max
}

func (fonts FontOutlines) getMinXY() point {
	if len(fonts) == 0 {
		return point{}
	}
	min := fonts[0].getMinXY()
	for _, f := range fonts {
		p := f.getMinXY()
		if min.X > p.X {
			min.X = p.X
		}
		if min.Y > p.Y {
			min.Y = p.Y
		}
	}
	return min
}

func (fonts FontOutlines) getRectangle() (x, y lib.Length) {
	min := fonts.getMinXY()
	max := fonts.getMaxXY()

	x = lib.Length(max.X - min.X)
	y = lib.Length(max.Y - min.Y)
	return x, y
}

func (fonts FontOutlines) getFirstPoint() point {
	x, _ := fonts.getRectangle()
	centerStartX := float64(x) / 2
	p := point{X: centerStartX, Y: 0}
	// TODO: consider font height

	return p
}

func (fonts FontOutlines) getEndPoint() point {
	max := fonts.getMaxXY()
	p := point{X: max.X + consts.LastMarging, Y: 0}
	return p
}
