package money_test

import (
	"encoding/json"
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

func TestJSON_RealWorld(t *testing.T) {
	type UserStruct struct {
		Name    string
		Balance *money.CAD
	}
	testTable := []struct {
		User   *UserStruct
		Result string
	}{
		{User: &UserStruct{Name: "John Doe", Balance: money.NewMoney(1, 0)}, Result: `{"Name":"John Doe","Balance":"CAD$1.00"}`},
		{User: &UserStruct{Name: "John Doe", Balance: money.NewMoney(1, 01)}, Result: `{"Name":"John Doe","Balance":"CAD$1.01"}`},
		{User: &UserStruct{Name: "John Doe", Balance: money.NewMoney(1, 10)}, Result: `{"Name":"John Doe","Balance":"CAD$1.10"}`},
		{User: &UserStruct{Name: "John Doe", Balance: money.NewMoney(-11, 0)}, Result: `{"Name":"John Doe","Balance":"CAD$-11.00"}`},
	}
	for testCase, test := range testTable {
		b, err := json.Marshal(test.User)
		if err != nil {
			t.Errorf("test %d: with error %v when marshaling", testCase, err)
			continue
		}
		if string(b) != test.Result {
			t.Errorf("test %d: expected %v, got %v", testCase, test.Result, string(b))
			continue
		}
		var user UserStruct
		err = json.Unmarshal(b, &user)
		if err != nil {
			t.Errorf("test %d: with error %v when unmarshaling", testCase, err)
			continue
		}
		if user.Name != test.User.Name || user.Balance.AsCent() != test.User.Balance.AsCent() {
			t.Errorf("test %d: expected %v, got %v", testCase, test.User, user)
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

func TestUnmarshalJSON_Failure(t *testing.T) {
	testTable := []struct {
		Bytes []byte
	}{
		{Bytes: []byte(`"CAD$1.1.00"`)},
		{Bytes: []byte(`"CAD$1.1.00`)},
		{Bytes: []byte(`"CAD$-"`)},
	}
	for testIndex, test := range testTable {
		var m money.CAD
		err := m.UnmarshalJSON(test.Bytes)
		if err == nil {
			t.Errorf("#%d : expected an error happen with values %v but it do not happened", testIndex, err)
		}

	}
}
