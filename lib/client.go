package tangocard

import (
	"encoding/json"
	"fmt"
)

// ClientOptions ...
type ClientOptions struct {
	Host string
}

// Client is used to make API requests to the Tangocard API
type Client struct {
	PlatformIdentifier string
	PlatformKey        string
	Options            ClientOptions
}

// CreateCustomer ...
func (c *Client) CreateCustomer(displayName, customerIdentifier string) (*Customer, error) {
	customer := &Customer{
		DisplayName:        displayName,
		CustomerIdentifier: customerIdentifier,
	}
	err := c.requestAndParseResponse("POST", "/customers", customer, customer)
	return customer, err
}

// GetCustomers returns a list of all customers
func (c *Client) GetCustomers() ([]*Customer, error) {
	customers := []*Customer{}
	err := c.requestAndParseResponse("GET", "/customers", nil, customers)
	return customers, err
}

// CreateAccount ...
func (c *Client) CreateAccount(email, displayName, accountIdentifier, customerIdentifier string) (*Account, error) {
	account := &Account{
		ContactEmail:      email,
		DisplayName:       displayName,
		AccountIdentifier: accountIdentifier,
	}
	path := fmt.Sprintf("/customers/%s/accounts", customerIdentifier)
	err := c.requestAndParseResponse("POST", path, account, account)
	return account, err
}

// GetAccounts returns a list of all customers
func (c *Client) GetAccounts() ([]*Account, error) {
	accounts := []*Account{}
	err := c.requestAndParseResponse("GET", "/accounts", nil, accounts)
	return accounts, err
}

// GetAccountDetails returns a list of all customers
func (c *Client) GetAccountDetails(accountID string) (*Account, error) {
	account := &Account{}
	path := fmt.Sprintf("/accounts/%s", accountID)
	err := c.requestAndParseResponse("GET", path, nil, account)
	return account, err
}

// GetCatalog gets all items in the Platform's Catalog
func (c *Client) GetCatalog() (*GetCatalogResponse, error) {
	cr := &GetCatalogResponse{}
	err := c.requestAndParseResponse("GET", "/catalogs", nil, cr)
	return cr, err
}

// GetOrders gets a list of Orders placed under the specified Platform
func (c *Client) GetOrders(r *GetOrdersRequest) (*GetOrdersResponse, error) {
	or := &GetOrdersResponse{}
	err := c.requestAndParseResponse("GET", "/orders", r, or)
	return or, err
}

// CreateOrder creates a new order of course
func (c *Client) CreateOrder(r *CreateOrderCriteria) (*Order, error) {
	or := &Order{}
	err := c.requestAndParseResponse("POST", "/orders", r, or)
	return or, err
}

// GetOrderDetails fetches detail for a specific order, identified by refOrderID
func (c *Client) GetOrderDetails(refOrderID string) (*Order, error) {
	or := &Order{}
	path := fmt.Sprintf("/orders/%s", refOrderID)
	err := c.requestAndParseResponse("GET", path, nil, or)
	return or, err
}

func (c *Client) requestAndParseResponse(method, url string, body interface{}, resObj interface{}) error {
	ar, err := c.request(method, url, body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(ar.Body, &resObj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(method, url string, body interface{}) (*APIResponse, error) {
	return sendRequest(c.Options.Host, method, url, c.PlatformIdentifier, c.PlatformKey, body)
}

// NewClient returns an API client
func NewClient(pid, pkey string, options *ClientOptions) *Client {
	if options == nil {
		options = &ClientOptions{}
	}

	return &Client{
		PlatformIdentifier: pid,
		PlatformKey:        pkey,
		Options:            *options,
	}
}
