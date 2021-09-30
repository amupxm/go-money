package main

import (
	"encoding/json"
	"fmt"

	"github.com/amupxm/go-money"
)

func main() {

	moneyOne, _ := money.ParseCAD("-$100.90")
	moneyTwo, _ := money.ParseCAD(".90$")

	result := moneyOne.Sub(*moneyTwo)
	fmt.Printf("money is : %v\n", result) // will prints money is:CAD$-101.80

	type jsonType struct {
		User  string    `json:"user"`
		Money money.CAD `json:"money"`
	}
	var someStruct jsonType
	someJSONString := `{"user":"amupxm","money":"-78.23"}`
	json.Unmarshal([]byte(someJSONString), &someStruct)
	fmt.Printf("Unmarshaled money as cents : %v\n", someStruct.Money.AsCent()) // Unmarshaled money as cents : -7823

	data := struct {
		Type  string    `json:"type"`
		Money money.CAD `json:"money"`
	}{
		Type:  "Offer",
		Money: *money.NewMoney(99, 99),
	}
	b, _ := json.Marshal(data)
	fmt.Printf("Marshaled money is : %v\n", string(b)) // Marshaled money is : {"type":"Offer","money":"CAD$99.99"}
}
