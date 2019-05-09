package lib

import (
	"bytes"
	"fmt"
)

// Order .
type Order struct {
	b bytes.Buffer
}

// Append .
func (o *Order) Append(addO Order) {
	addO.b.WriteTo(&o.b)
}

// WriteString .
func (o *Order) WriteString(str string) {
	o.b.WriteString(str)
}

// Printf .
func (o *Order) Printf(format string, a ...interface{}) {
	fmt.Fprintf(&o.b, format, a...)
}

// String .
func (o Order) String() string {
	return o.b.String()
}
