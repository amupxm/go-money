package money_test

import (
	"testing"

	"github.com/amupxm/go-money"
)

func TestScan_Success(t *testing.T) {
	testTable := []struct {
		Value    interface{}
		Expected money.CAD
	}{
		{Value: -5, Expected: money.Cents(-500)},
		{Value: -4, Expected: money.Cents(-400)},
		{Value: -3, Expected: money.Cents(-300)},
		{Value: -2, Expected: money.Cents(-200)},
		{Value: -1, Expected: money.Cents(-100)},
		{Value: 0, Expected: money.Cents(0)},
		{Value: 1, Expected: money.Cents(100)},
		{Value: 2, Expected: money.Cents(200)},
		{Value: byte(3), Expected: money.Cents(300)},
		{Value: "4", Expected: money.Cents(400)},
		{Value: 5, Expected: money.Cents(500)},
		{Value: "0.01", Expected: money.Cents(1)},
		{Value: "0.02", Expected: money.Cents(2)},
	}

	for testNumber, test := range testTable {
		var m money.CAD
		clone := test.Value
		err := m.Scan(&clone)

		if err != nil {
			t.Errorf("number %d :: Error on converting to money %v with error %s", testNumber+1, test.Value, err.Error())
			continue
		}
		if m.AsCent() != test.Expected.AsCent() {
			t.Errorf("#%d : Expected %v got %v", testNumber, test.Expected, m)
			continue
		}

	}
}

func TestScan_failure(t *testing.T) {
	testTable := []struct {
		Value interface{}
	}{
		{Value: []byte("b")},        // as invalid byte arr
		{Value: "string"},           // as invalid string
		{Value: []int{1, 3}},        // as slice of int
		{Value: []string{"1", "3"}}, // as slice of string ( it will cause unhandled error )
		{Value: -32.235},            // as float

	}
	for testNumber, test := range testTable {
		var m money.CAD
		clone := test.Value
		err := m.Scan(&clone)
		if err == nil {
			t.Errorf("#%d : Expected error but value %v returned nil which returns %v as value", testNumber, test.Value, m)
			continue
		}

	}
}
func TestValue(t *testing.T) {
	testTable := []struct {
		Value    money.CAD
		Expected string
	}{
		{Value: money.Cents(10001), Expected: "100.1"},
		{Value: money.Cents(-500), Expected: "-5.0"},
		{Value: money.Cents(-400), Expected: "-4.0"},
		{Value: money.Cents(-300), Expected: "-3.0"},
		{Value: money.Cents(-200), Expected: "-2.0"},
		{Value: money.Cents(-100), Expected: "-1.0"},
		{Value: money.Cents(0), Expected: "0.0"},
		{Value: money.Cents(100), Expected: "1.0"},
	}
	for testNumber, test := range testTable {
		v, err := test.Value.Value()
		if err != nil {
			t.Errorf("#%d : Error on converting to string %v with error %s", testNumber, test.Value, err.Error())
			continue
		}
		if v != test.Expected {
			t.Errorf("#%d : Expected %s got %s", testNumber, test.Expected, v)
			continue
		}
	}
}
