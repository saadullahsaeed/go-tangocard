package tangocard

import "time"

const (
	// AccountStatusActive is when an account is active
	AccountStatusActive = "active"
)

// Account struct represents the accounts in the tango system
type Account struct {
	AccountIdentifier string     `json:"accountIdentifier"`
	DisplayName       string     `json:"displayName"`
	CreatedAt         *time.Time `json:"createdAt"`
	Status            string     `json:"status"`
}

// IsActive checks if the account status is active
func (a *Account) IsActive() bool {
	return a.Status == AccountStatusActive
}

// Customer struct represents a customer in the tango system
type Customer struct {
	CustomerIdentifier string     `json:"customerIdentifier"`
	DisplayName        string     `json:"displayName"`
	CreatedAt          *time.Time `json:"createdAt"`
	Status             string     `json:"status"`
	Accounts           []Account  `json:"accounts"`
}
