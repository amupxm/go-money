package money

import (
	"encoding/json"
)

// MarshalJSON marshals the money to JSON.
func (c CAD) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON unmarshals the money from JSON.
func (c *CAD) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	parsedCad, err := ParseCAD(str)
	if err != nil {
		return err
	}

	*c = *NewMoney(parsedCad.dollar, parsedCad.cents)
	return nil
}
