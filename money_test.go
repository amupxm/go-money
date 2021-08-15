package money_test

import (
	"testing"

	"github.com/amupxm/go-money"
)

func TestCents(t *testing.T) {
	testTable := []struct {
		Int      int64
		Expected money.CAD
	}{
		{0, *money.NewMoney(0, 0)},
		{-0, *money.NewMoney(0, 0)},
		{105, *money.NewMoney(1, 5)},
		{1050, *money.NewMoney(10, 50)},
		{-105, *money.NewMoney(-1, -5)},
		{-1, *money.NewMoney(0, -1)},
		{-500, *money.NewMoney(-5, 0)},
		{-100000005, *money.NewMoney(-1000000, -05)},
		{-100000005, *money.NewMoney(-1000000, -05)},
	}
	for testCase, test := range testTable {
		if got := money.Cents(test.Int); got != test.Expected {
			t.Errorf("#%d : Cents(%d) = %v; want %v", testCase, test.Int, got, test.Expected)
			continue
		}
	}
}

func Test_ParseCAD_Success(t *testing.T) {
	testTable := []struct {
		S      string
		Result money.CAD
	}{
		{"CAD-$1,234.56", *money.NewMoney(-1234, -56)},
		{"$-.09", *money.NewMoney(0, -9)},
		{"CAD$-.9", *money.NewMoney(0, -90)},
		{"0.02", *money.NewMoney(0, 2)},
		{"-5", *money.NewMoney(-5, 0)},
		{"CAD$-.09", *money.NewMoney(0, -9)},
		{"9¢", *money.NewMoney(0, 9)},
		{"-9¢", *money.NewMoney(0, -9)},
		{"-$1235.56", *money.NewMoney(-1235, -56)},
		{"-123456¢", *money.NewMoney(-1234, -56)},
		{"123456¢", *money.NewMoney(1234, 56)},
		{"$.000123456", *money.NewMoney(0, 0)},
		{"-$1234.56", *money.NewMoney(-1234, -56)},
		{"-$1,234.56", *money.NewMoney(-1234, -56)},
		{"$-1,234.56", *money.NewMoney(-1234, -56)},
		{"CAD -$1234.56", *money.NewMoney(-1234, -56)},
		{"CAD $-1234.56", *money.NewMoney(-1234, -56)},
		{"CAD-$1,234.56", *money.NewMoney(-1234, -56)},
		{"CAD$-1,234.56", *money.NewMoney(-1234, -56)},
		{"$1234.56", *money.NewMoney(1234, 56)},
		{"$1,234.56", *money.NewMoney(1234, 56)},
		{"CAD $1234.56", *money.NewMoney(1234, 56)},
		{"CAD $1,234.56", *money.NewMoney(1234, 56)},
		{"CAD$1234.56", *money.NewMoney(1234, 56)},
		{"CAD$1,234.56", *money.NewMoney(1234, 56)},
		{"$0.09", *money.NewMoney(0, 9)},
		{"$.09", *money.NewMoney(0, 9)},
		{"-$0.09", *money.NewMoney(0, -9)},
		{"-$.09", *money.NewMoney(0, -9)},
		{" $-0.09", *money.NewMoney(0, -9)},
		{" $-.09", *money.NewMoney(0, -9)},
		{"CAD $0.09", *money.NewMoney(0, 9)},
		{"CAD $.09", *money.NewMoney(0, 9)},
		{"CAD -$0.09", *money.NewMoney(0, -9)},
		{"CAD -$.09", *money.NewMoney(0, -9)},
		{"CAD $-0.09", *money.NewMoney(0, -9)},
		{"CAD $-.09", *money.NewMoney(0, -9)},
		{"CAD$0.09", *money.NewMoney(0, 9)},
		{"CAD$.09", *money.NewMoney(0, 9)},
		{"CAD-$0.09", *money.NewMoney(0, -9)},
		{"CAD-$.09", *money.NewMoney(0, -9)},
		{" CAD$-0.09", *money.NewMoney(0, -9)},
		{"CAD$-.09", *money.NewMoney(0, -9)},
		{"9¢", *money.NewMoney(0, 9)},
		{" -9¢", *money.NewMoney(0, -9)},
		{"123456¢", *money.NewMoney(1234, 56)},
		{"-123456¢", *money.NewMoney(-1234, -56)},
	}
	for testCase, test := range testTable {
		got, err := money.ParseCAD(test.S)
		if err != nil {
			t.Errorf("#%d : we get an error with value %s adn error %v", testCase, test.S, err)
			continue
		}
		if *got != test.Result {
			t.Errorf("#%d : we got %v from values %v but expected %v ", testCase, got.AsCent(), test.Result.AsCent(), test.S)
			continue
		}
	}
}

func Test_ParseCAD_Failure(t *testing.T) {
	testTable := []struct {
		S string
	}{
		{"CAD-$1¢,234.56"},
		{"$-.0.9"},
		{"$"},
		{"$-¢."},
		{"$1-2.-1"},
		{"$.123456"},
	}
	for testCase, test := range testTable {
		got, err := money.ParseCAD(test.S)

		if err == nil {
			t.Errorf("#%d : ParseCAD(%s) = %v; but expected an error", testCase, test.S, got)
			continue
		}
	}
}
