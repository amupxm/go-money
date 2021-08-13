package money_test

import (
	"testing"

	"github.com/amupxm/go-money"
)

func TestAbs(t *testing.T) {
	testTable := []struct {
		Money    *money.CAD
		Expected money.CAD
	}{
		{money.NewMoney(0, 0), *money.NewMoney(0, 0)},
		{money.NewMoney(1, 0), *money.NewMoney(1, 0)},
		{money.NewMoney(2, 0), *money.NewMoney(2, 0)},
		{money.NewMoney(3, 0), *money.NewMoney(3, 0)},
		{money.NewMoney(4, 0), *money.NewMoney(4, 0)},
		{money.NewMoney(5, 0), *money.NewMoney(5, 0)},
		{money.NewMoney(6, 0), *money.NewMoney(6, 0)},
		{money.NewMoney(7, 0), *money.NewMoney(7, 0)},
		{money.NewMoney(-1, 0), *money.NewMoney(1, 0)},
		{money.NewMoney(-2, 0), *money.NewMoney(2, 0)},
		{money.NewMoney(-3, 0), *money.NewMoney(3, 0)},
		{money.NewMoney(-4, 0), *money.NewMoney(4, 0)},
		{money.NewMoney(0, -1), *money.NewMoney(0, 1)},
		{money.NewMoney(0, -2), *money.NewMoney(0, 2)},
		{money.NewMoney(0, -3), *money.NewMoney(0, 3)},
		{money.NewMoney(0, -4), *money.NewMoney(0, 4)},
	}
	for testCase, test := range testTable {
		if test.Money.Abs() != test.Expected {
			t.Errorf("#%d : failed: Expected %v, got %v", testCase, test.Expected, test.Money.Abs())
			continue
		}
	}
}

func TestAdd(t *testing.T) {
	testTable := []struct {
		First    *money.CAD
		Second   money.CAD
		Expected money.CAD
	}{
		{money.NewMoney(0, 0), *money.NewMoney(0, 0), *money.NewMoney(0, 0)},
		{money.NewMoney(1, 0), *money.NewMoney(1, 0), *money.NewMoney(2, 0)},
		{money.NewMoney(-1, 0), *money.NewMoney(-1, 0), *money.NewMoney(-2, 0)},
		{money.NewMoney(-1, 0), *money.NewMoney(1, 0), *money.NewMoney(0, 0)},
		{money.NewMoney(0, -01), *money.NewMoney(0, -10), *money.NewMoney(0, -11)},
		{money.NewMoney(99999999, 99), *money.NewMoney(0, 1), *money.NewMoney(100000000, 0)},
	}
	for testCase, test := range testTable {
		if test.First.Add(test.Second) != test.Expected {
			t.Errorf("#%d : failed: Expected %v, got %v", testCase, test.Expected, test.First.Add(test.Second))
			continue
		}
	}
}

func TestMul(t *testing.T) {
	testTable := []struct {
		First    *money.CAD
		Scalar   int64
		Expected money.CAD
	}{
		{money.NewMoney(0, 0), 1, *money.NewMoney(0, 0)},
		{money.NewMoney(1, 1), 22, *money.NewMoney(22, 22)},
		{money.NewMoney(-1, 0), -1, *money.NewMoney(1, 0)},
		{money.NewMoney(1, 0), -1, *money.NewMoney(-1, 0)},
		{money.NewMoney(1, 50), 2, *money.NewMoney(3, 0)},
	}
	for testCase, test := range testTable {
		if res := test.First.Mul(test.Scalar); res != test.Expected {
			t.Errorf("#%d : failed: Expected %v, got %v", testCase, test.Expected, res)
			continue
		}
	}
}

func TestCanonicalForm(t *testing.T) {
	testTable := []struct {
		Case                         *money.CAD
		ExpectedDolar, ExpectedCents int64
	}{
		{money.NewMoney(0, 0), 0, 0},
		{money.NewMoney(1, 0), 1, 0},
		{money.NewMoney(2, 0), 2, 0},
		{money.NewMoney(30, 0), 30, 0},
		{money.NewMoney(400, 0), 400, 0},
		{money.NewMoney(-0, -0), 0, 0},
		{money.NewMoney(-100000, -10), -100000, -10},
		{money.NewMoney(-0, -10), 0, -10},
	}

	for testCase, test := range testTable {
		if dollar, cents := test.Case.CanonicalForm(); dollar != test.ExpectedDolar || cents != test.ExpectedCents {
			t.Errorf("#%d : failed: Expected %d.%0d got %d.%0d", testCase, test.ExpectedDolar, test.ExpectedCents, dollar, cents)
			continue
		}
	}
}

func TestSub(t *testing.T) {
	testTable := []struct {
		First    *money.CAD
		Second   money.CAD
		Expected money.CAD
	}{
		{money.NewMoney(0, 0), *money.NewMoney(0, 0), *money.NewMoney(0, 0)},
		{money.NewMoney(1, 0), *money.NewMoney(1, 0), *money.NewMoney(0, 0)},
		{money.NewMoney(-1, 0), *money.NewMoney(-1, 0), *money.NewMoney(0, 0)},
		{money.NewMoney(-1, 0), *money.NewMoney(1, 0), *money.NewMoney(-2, 0)},
		{money.NewMoney(0, -01), *money.NewMoney(0, -10), *money.NewMoney(0, 9)},
		{money.NewMoney(99999999, 99), *money.NewMoney(0, -99), *money.NewMoney(100000000, 98)},
	}
	for testCase, test := range testTable {
		if test.First.Sub(test.Second) != test.Expected {
			t.Errorf("#%d : failed: Expected %v, got %v", testCase, test.Expected, test.First.Sub(test.Second))
			continue
		}
	}
}
