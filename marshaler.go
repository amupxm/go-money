package money

import (
	"encoding/json"
)

func (c CAD) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

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
