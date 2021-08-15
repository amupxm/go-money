
<p align="right">
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
You can pass all valid money string like <b><i>$CAD-9.0</i> , <i>CAD$79</i> , <i>23.65</i> , <i>-0.9</i> , <i>.09</i> , <i>78.98CAD</i> , <i>98.07$</i></b> and ... 

It also provides inner umarshaler,unmarshaler,Scanner and valuer which make it super easy to use 

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
Money type in compatible with postgres monet type.


Give star if interested  :star2:
