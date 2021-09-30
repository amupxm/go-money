package money

import "fmt"

//GoString creates sting
func (c CAD) GoString() string {
	return fmt.Sprintf("money.cents(%d)", c.AsCent())
}

func (c CAD) String() string {
	z := int64(1)
	if c.dollar < 0 || c.cents < 0 {
		z = -1
	}
	r := ""
	if z == -1 {
		r = "-"
	}
	return fmt.Sprintf("CAD$%s%d.%02d", r, c.dollar*z, c.cents*z)
}
