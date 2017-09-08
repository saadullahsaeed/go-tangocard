# Tango Card RaaS API Client for v2

[![GoDoc](https://godoc.org/github.com/saadullahsaeed/go-tangocard?status.svg)](https://godoc.org/github.com/saadullahsaeed/go-tangocard/lib)

Golang client for the Tangocard RaaS API v2.

This is a Work in Progress and I've so far only implemented the methods I needed for my project.

Currently supported methods:
* Get Customers
* Get Accounts 
* Get Account Details
* Get Catalog
* Get Orders
* Create Order
* Get Order Details

### Installation

```
go get github.com/saadullahsaeed/go-tangocard
```

### Usage

##### Get Catalog
```go
client := tangocard.NewClient(
	"PlatformIDHere",
	"PlatformKeyHere",
	nil,
)
res, err := client.GetCatalog()
fmt.Println(err)
fmt.Println(res)
```

##### Create Order

```go
req := &tangocard.CreateOrderCriteria{
	AccountIdentifier:  "test",
	Amount:             50.0,
	CustomerIdentifier: "test",
	UTID:               "U12345",
}

client := tangocard.NewClient(
	"PlatformIDHere",
	"PlatformKeyHere",
	nil,
)
res, err := client.CreateOrder(req)
fmt.Println(err)
fmt.Println(res)
```
