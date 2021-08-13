package money

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type (
	CAD struct {
		dollar int64
		cents  int64
	}
)

func (c CAD) AsCent() int64 {
	return c.cents + c.dollar*100
}

func ParseCAD(s string) (*CAD, error) {
	// helper function to parse a string into a int64
	convToInt := func(str string) int64 {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	// CheckIs negative?
	isNeg := int64(1)
	if strings.Contains(s, "-") {
		isNeg = -1
		s = strings.Replace(s, "-", "", 1)
	}
	// CheckIs all cents?
	iscents := strings.Contains(s, "¢") || strings.Contains(s, "cents")
	// remove spaces and commas
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", "", -1)
	// Extract only the numbers
	re := regexp.MustCompile(`([-]|\.)?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(s, -1)
	if len(submatchall) != 1 {
		return nil, errors.New("invalid money string")
	}
	s = submatchall[0]
	// if string starts with . add zero to the front
	if strings.HasPrefix(s, ".") {
		s = fmt.Sprintf("0%s", s)
	}
	// split string by dot
	sArr := strings.Split(s, ".")
	// if string has more than 2 parts, return error
	if len(sArr) > 2 {
		return nil, errors.New("invalid money string")
	}
	//if its only cents , convert ans return only cents

	if iscents && len(sArr) == 1 {
		c := convToInt(sArr[0])
		asCent := Cents(c * isNeg)
		return &asCent, nil
	}

	//if itd dollas , convert to cents and return

	if !iscents && len(sArr) == 2 {
		d := convToInt(sArr[0])
		// limit c len to 2
		scalar := 0
		// count of 0 in begining of sArr[1]
		for _, c := range sArr[1] {
			if c == '0' {
				scalar++
			} else {
				break
			}
		}
		if scalar > 2 {
			sArr[1] = "0"
		}
		if scalar == 0 {
			sArr[1] = fmt.Sprintf("%s0", sArr[1])

		}
		if len(sArr[1]) > 2 {
			sArr[1] = sArr[1][:2]
		}
		c := convToInt(sArr[1])
		dollars := CAD{dollar: d * isNeg, cents: (c * isNeg)}

		return &dollars, nil
	}
	if !iscents && len(sArr) == 1 {
		dollar := CAD{dollar: convToInt(sArr[0]) * isNeg, cents: 0}
		return &dollar, nil
	}

	return nil, errors.New(fmt.Sprint(iscents, len(sArr), s))
}

func NewMoney(dollars, cents int64) *CAD {
	return &CAD{
		dollar: dollars,
		cents:  cents,
	}
}
func Cents(i int64) CAD {
	neg := false
	if i < 0 {
		neg = true
		i = -i
	}
	c := CAD{
		dollar: i / 100,
		cents:  i % 100,
	}
	if neg {
		c.dollar = -c.dollar
		c.cents = -c.cents
	}
	return c
}
