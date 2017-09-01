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

// GetCustomers returns a list of all customers
func (c *Client) GetCustomers() ([]*Customer, error) {
	customers := []*Customer{}
	err := c.requestAndParseResponse("GET", "/customers", nil, customers)
	return customers, err
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
	js, err := c.request(method, url, body)
	if err != nil {
		return err
	}
	fmt.Println(string(js))
	err = json.Unmarshal(js, &resObj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(method, url string, body interface{}) (json.RawMessage, error) {
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
