# go-tangocard

Golang client for the Tangocard RaaS API.

This is a Work in Progress.

Currently supported methods:
* Get Catalog
* Get Orders
* Create Order


### Examples

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
