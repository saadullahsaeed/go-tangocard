package tangocard

import "time"

// TangoError ...
type TangoError struct {
	Path         string `json:"path"`
	Message      string `json:"message"`
	InvalidValue string `json:"invalidValue"`
	Constraint   string `json:"constraint"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Timestamp  *time.Time    `json:"timestamp"`
	RequestID  string        `json:"requestId"`
	Path       string        `json:"path"`
	HTTPCode   int           `json:"httpCode"`
	HTTPPhrase int           `json:"httpPhrase"`
	Errors     []*TangoError `json:"errors"`
}

// GetCatalogResponse ...
type GetCatalogResponse struct {
	Brands      []*Brand `json:"brands"`
	CatalogName string   `json:"catalogName"`
}

// GetOrdersResponse ...
type GetOrdersResponse struct {
	Orders []*Order `json:"orders"`
}
