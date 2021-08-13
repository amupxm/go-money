
<p align="right">
  <a href="https://actions-badge.atrox.dev/amupxm/go-money/goto?ref=main"><img alt="Build Status" src="https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Famupxm%2Fgo-money%2Fbadge%3Fref%3Dmain&style=flat-square" /></a>
  <a href="ttps://codecov.io/gh/amupxm/go-money"><img alt="github-status-action status" src="https://codecov.io/gh/amupxm/go-money/branch/main/graph/badge.svg?token=SPO9OYHIHE"></a>
</p>


# Go Money

Golang pure money calculation library provides safety and ease of use.


## Usage 

go-money can parse money string to money type but you can use numbers as input too.


```go

	moneyOne, _ := money.ParseCAD("-$100.90")
	moneyTwo, _ := money.ParseCAD(".90$")

	result := moneyOne.Sub(*moneyTwo)
	fmt.Printf("money is : %v\n", result) // will prints money is : CAD$-101.80
```

it also provides inner umarshaler,unmarshaler,Scanner and valuer which make it super easy to use 

### unmarshal
```go 
	type jsonType struct {
		User  string    `json:"user"`
		Money money.CAD `json:"money"`
	}
	var someStruct jsonType
	someJsonString := `{"user":"amupxm","money":"-78.23"}`
	json.Unmarshal([]byte(someJsonString), &someStruct)
	fmt.Printf("Unmarshaled money as cents : %v\n", someStruct.Money.AsCent()) // Unmarshaled money as cents : -7823


```


### marshal
```go 
	data := struct {
		Type  string    `json:"type"`
		Money money.CAD `json:"money"`
	}{
		Type:  "Offer",
		Money: *money.NewMoney(99, 99),
	}
	b, _ := json.Marshal(data)
	fmt.Printf("Marshaled money is : %v\n", string(b)) // Marshaled money is : {"type":"Offer","money":"CAD$99.99"}

```

And you can scan,value to sql/database without any troubles =)


```go 
    var theMoney money.CAD
	data := struct {
		UserName  string    `json:"user_name"`
		Money money.CAD `json:"money"`
	}{
		UserName:  "amupxm",
		Money: *money.NewMoney(99, 99),
	}
	err := pg.db.QueryRow(`INSERT INTO transactions (username, money)
	 VALUES ($1, $2) RETURNING money`, user.UserName,user.money).Scan(&theMoney)
    fmt.Println("Money is %v\n" ,theMoney )  // Money is : CAD$99.99
```
Give star if interested  :star2:
