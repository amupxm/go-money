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
		// values from database will be  string or []byte or int64
		{Value: int64(-5), Expected: money.Cents(-500)},
		{Value: int64(-4), Expected: money.Cents(-400)},
		{Value: int64(-3), Expected: money.Cents(-300)},
		{Value: int64(-2), Expected: money.Cents(-200)},
		{Value: int64(-1), Expected: money.Cents(-100)},
		{Value: int64(0), Expected: money.Cents(0)},
		{Value: int64(1), Expected: money.Cents(100)},
		{Value: int64(2), Expected: money.Cents(200)},
		{Value: byte(3), Expected: money.Cents(300)},
		{Value: "4", Expected: money.Cents(400)},
		{Value: int64(5), Expected: money.Cents(500)},
		{Value: "0.01", Expected: money.Cents(1)},
		{Value: "0.02", Expected: money.Cents(2)},
		{Value: "0.09", Expected: money.Cents(9)},
		{Value: "0.90", Expected: money.Cents(90)},
		{Value: "9.00", Expected: money.Cents(900)},
		{Value: "90.00", Expected: money.Cents(9000)},
		{Value: "-0.09", Expected: money.Cents(-9)},
		{Value: "-0.90", Expected: money.Cents(-90)},
		{Value: "-9.00", Expected: money.Cents(-900)},
		{Value: "-90.00", Expected: money.Cents(-9000)},
		{Value: "-98765432.10", Expected: money.Cents(-9876543210)},
		{Value: "98765432.10", Expected: money.Cents(9876543210)},
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
		{Value: money.Cents(10001), Expected: "100.01"},
		{Value: money.Cents(-500), Expected: "-5.0"},
		{Value: money.Cents(-400), Expected: "-4.0"},
		{Value: money.Cents(-300), Expected: "-3.0"},
		{Value: money.Cents(-200), Expected: "-2.0"},
		{Value: money.Cents(-100), Expected: "-1.0"},
		{Value: money.Cents(100), Expected: "1.0"},
		//==========================================
		{Value: money.Cents(-11), Expected: "-0.11"},
		{Value: money.Cents(-10), Expected: "-0.10"},
		{Value: money.Cents(-9), Expected: "-0.09"},
		{Value: money.Cents(-8), Expected: "-0.08"},
		{Value: money.Cents(-7), Expected: "-0.07"},
		{Value: money.Cents(-6), Expected: "-0.06"},
		{Value: money.Cents(-5), Expected: "-0.05"},
		{Value: money.Cents(-4), Expected: "-0.04"},
		{Value: money.Cents(-3), Expected: "-0.03"},
		{Value: money.Cents(-2), Expected: "-0.02"},
		{Value: money.Cents(-1), Expected: "-0.01"},
		{Value: money.Cents(0), Expected: "0.0"},
		{Value: money.Cents(1), Expected: "0.01"},
		{Value: money.Cents(2), Expected: "0.02"},
		{Value: money.Cents(3), Expected: "0.03"},
		{Value: money.Cents(4), Expected: "0.04"},
		{Value: money.Cents(5), Expected: "0.05"},
		{Value: money.Cents(6), Expected: "0.06"},
		{Value: money.Cents(7), Expected: "0.07"},
		{Value: money.Cents(8), Expected: "0.08"},
		{Value: money.Cents(9), Expected: "0.09"},
		{Value: money.Cents(10), Expected: "0.10"},
		{Value: money.Cents(11), Expected: "0.11"},
		{Value: money.Cents(-113), Expected: "-1.13"},
		{Value: money.Cents(-112), Expected: "-1.12"},
		{Value: money.Cents(-111), Expected: "-1.11"},
		{Value: money.Cents(-110), Expected: "-1.10"},
		{Value: money.Cents(-109), Expected: "-1.09"},
		{Value: money.Cents(-108), Expected: "-1.08"},
		{Value: money.Cents(-107), Expected: "-1.07"},
		{Value: money.Cents(-106), Expected: "-1.06"},
		{Value: money.Cents(-105), Expected: "-1.05"},
		{Value: money.Cents(-104), Expected: "-1.04"},
		{Value: money.Cents(-103), Expected: "-1.03"},
		{Value: money.Cents(-102), Expected: "-1.02"},
		{Value: money.Cents(-101), Expected: "-1.01"},
		{Value: money.Cents(-100), Expected: "-1.0"},
		{Value: money.Cents(-99), Expected: "-0.99"},
		{Value: money.Cents(-98), Expected: "-0.98"},
		{Value: money.Cents(98), Expected: "0.98"},
		{Value: money.Cents(99), Expected: "0.99"},
		{Value: money.Cents(100), Expected: "1.0"},
		{Value: money.Cents(101), Expected: "1.01"},
		{Value: money.Cents(102), Expected: "1.02"},
		{Value: money.Cents(103), Expected: "1.03"},
		{Value: money.Cents(104), Expected: "1.04"},
		{Value: money.Cents(105), Expected: "1.05"},
		{Value: money.Cents(106), Expected: "1.06"},
		{Value: money.Cents(107), Expected: "1.07"},
		{Value: money.Cents(108), Expected: "1.08"},
		{Value: money.Cents(109), Expected: "1.09"},
		{Value: money.Cents(110), Expected: "1.10"},
		{Value: money.Cents(111), Expected: "1.11"},
		{Value: money.Cents(112), Expected: "1.12"},
		{Value: money.Cents(113), Expected: "1.13"},
		{Value: money.Cents(-1013), Expected: "-10.13"},
		{Value: money.Cents(-1012), Expected: "-10.12"},
		{Value: money.Cents(-1011), Expected: "-10.11"},
		{Value: money.Cents(-1010), Expected: "-10.10"},
		{Value: money.Cents(-1009), Expected: "-10.09"},
		{Value: money.Cents(-1008), Expected: "-10.08"},
		{Value: money.Cents(-1007), Expected: "-10.07"},
		{Value: money.Cents(-1006), Expected: "-10.06"},
		{Value: money.Cents(-1005), Expected: "-10.05"},
		{Value: money.Cents(-1004), Expected: "-10.04"},
		{Value: money.Cents(-1003), Expected: "-10.03"},
		{Value: money.Cents(-1002), Expected: "-10.02"},
		{Value: money.Cents(-1001), Expected: "-10.01"},
		{Value: money.Cents(-1000), Expected: "-10.0"},
		{Value: money.Cents(-999), Expected: "-9.99"},
		{Value: money.Cents(-998), Expected: "-9.98"},
		{Value: money.Cents(998), Expected: "9.98"},
		{Value: money.Cents(999), Expected: "9.99"},
		{Value: money.Cents(1000), Expected: "10.0"},
		{Value: money.Cents(1001), Expected: "10.01"},
		{Value: money.Cents(1002), Expected: "10.02"},
		{Value: money.Cents(1003), Expected: "10.03"},
		{Value: money.Cents(1004), Expected: "10.04"},
		{Value: money.Cents(1005), Expected: "10.05"},
		{Value: money.Cents(1006), Expected: "10.06"},
		{Value: money.Cents(1007), Expected: "10.07"},
		{Value: money.Cents(1008), Expected: "10.08"},
		{Value: money.Cents(1009), Expected: "10.09"},
		{Value: money.Cents(1010), Expected: "10.10"},
		{Value: money.Cents(1011), Expected: "10.11"},
		{Value: money.Cents(1012), Expected: "10.12"},
		{Value: money.Cents(1013), Expected: "10.13"},
		{Value: money.Cents(-123456), Expected: "-1234.56"},
		{Value: money.Cents(123456), Expected: "1234.56"},
		{Value: money.Cents(-9876543210), Expected: "-98765432.10"},
		{Value: money.Cents(9876543210), Expected: "98765432.10"},
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
