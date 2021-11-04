package money

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
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

//MarshalBSON marshalees bson
func (c *CAD) MarshalBSON() ([]byte, error) {
	type m CAD
	return bson.Marshal(c.cents)
}

// UnmarshalBSON unmarhal bson
func (c *CAD) UnmarshalBSON(data []byte) error {
	var ReallyBigAlias int64
	err := bson.Unmarshal(data, &ReallyBigAlias)
	if err != nil {
		return err
	}

	*c = Cents(ReallyBigAlias)
	return nil
}
