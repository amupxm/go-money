package money_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/amupxm/go-money"
)

func TestStringer(t *testing.T) {
	testTable := []struct {
		S        string
		Expected string
	}{
		{"1.23", "CAD$1.23"},
		{"-CAD$.2", "CAD$-0.20"},
	}

	for _, tt := range testTable {
		var b bytes.Buffer
		money, _ := money.ParseCAD(tt.S)
		fmt.Fprint(
			&b,
			money,
		)
		if b.String() != tt.Expected {
			t.Error("invalid got : " + b.String() + " expected " + tt.Expected)
			continue
		}
	}
}

func TestGoStringer(t *testing.T) {
	testTable := []struct {
		S        string
		Expected string
	}{
		{"1.23", "money.cents(123)"},
		{"-CAD$.2", "money.cents(-20)"},
	}

	for _, tt := range testTable {
		m, _ := money.ParseCAD(tt.S)
		var b bytes.Buffer
		fmt.Fprintf(
			&b,
			"%#v",
			m,
		)
		if b.String() != tt.Expected {
			t.Error("invalid got : " + b.String() + " expected " + tt.Expected)
			continue
		}
	}

}
