package tangocard

// GetCatalogResponse ...
type GetCatalogResponse struct {
	Brands      []*Brand `json:"brands"`
	CatalogName string   `json:"catalogName"`
}

// GetOrdersResponse ...
type GetOrdersResponse struct {
	Orders []*Order `json:"orders"`
}
