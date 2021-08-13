package money_test

import (
	"testing"

	"github.com/amupxm/go-money"
)

func TestMarshalJSON(t *testing.T) {
	testTable := []struct {
		M *money.CAD
		S string
	}{
		{M: money.NewMoney(1, 0), S: `"CAD$1.00"`},
		{M: money.NewMoney(1, 01), S: `"CAD$1.01"`},
		{M: money.NewMoney(1, 10), S: `"CAD$1.10"`},
		{M: money.NewMoney(-11, 0), S: `"CAD$-11.00"`},
	}
	for testIndex, test := range testTable {
		b, err := test.M.MarshalJSON()
		if err != nil {
			t.Errorf("#%d: with error %v", testIndex, err)
			continue
		}
		if string(b) != test.S {
			t.Errorf("#%d: expected %v, got %v", testIndex, test.S, string(b))
			continue
		}
	}

}

func TestUnmarshalJSON_Success(t *testing.T) {
	testTable := []struct {
		M     *money.CAD
		Bytes []byte
	}{
		{M: money.NewMoney(1, 0), Bytes: []byte(`"CAD$1.00"`)},
		{M: money.NewMoney(-11, 0), Bytes: []byte(`"CAD$-11.00"`)},
		{M: money.NewMoney(1, 01), Bytes: []byte(`"CAD$1.01"`)},
		{M: money.NewMoney(1, 10), Bytes: []byte(`"CAD$1.10"`)},
	}
	for testIndex, test := range testTable {
		var m money.CAD
		err := m.UnmarshalJSON(test.Bytes)
		if err != nil {
			t.Errorf("test %d: with error %v", testIndex, err)
		}
		if m.AsCent() != test.M.AsCent() {
			t.Errorf("test %d: expected %v, got %v", testIndex, test.M, m)
		}
	}
}
