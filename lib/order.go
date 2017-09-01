package tangocard

// Denomination struct represents the amount charge for the order
type Denomination struct {
	CurencyCode  string  `json:"currencyCode"`
	ExchangeRate float64 `json:"exchangeRate"`
	Fee          float64 `json:"fee"`
	Total        float64 `json:"total"`
	Value        float64 `json:"value"`
}

// Contact struct represents a contact for the order (email recipient OR sender)
type Contact struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Credential ...
type Credential struct {
	Label string `json:"label"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Reward ...
type Reward struct {
	CredentialList         []Credential      `json:"credentialList"`
	Credentials            map[string]string `json:"credentials"`
	RedemptionInstructions string            `json:"redemptionInstructions"`
}

// Order struct represents a giftcard/reward order with the TangoCard API
type Order struct {
	AccountIdentifier      string       `json:"accountIdentifier"`
	AmountCharged          Denomination `json:"amountCharged"`
	Campaign               string       `json:"campaign"`
	CustomerIdentifier     string       `json:"customerIdentifier"`
	Denomination           Denomination `json:"denomination"`
	EmailSubject           string       `json:"emailSubject"`
	ETID                   string       `json:"etid"`
	ExternalRefID          string       `json:"externalRefID"`
	Message                string       `json:"message"`
	Notes                  string       `json:"notes"`
	Recipient              *Contact     `json:"recipient"`
	RedemptionInstructions string       `json:"redemptionInstructions"`
	Reward                 *Reward      `json:"reward"`
	ReferenceOrderID       string       `json:"referenceOrderID"`
	RewardName             string       `json:"rewardName"`
	SendEmail              bool         `json:"sendEmail"`
	Sender                 *Contact     `json:"sender"`
	Status                 string       `json:"status"`
	UTID                   string       `json:"utid"`
}

// CreateOrderCriteria represents the request to create a new order
type CreateOrderCriteria struct {
	AccountIdentifier  string   `json:"accountIdentifier"`
	Amount             float64  `json:"amount"`
	Campaign           string   `json:"campaign"`
	CustomerIdentifier string   `json:"customerIdentifier"`
	EmailSubject       string   `json:"emailSubject"`
	ExternalRefID      string   `json:"externalRefID"`
	Message            string   `json:"message"`
	Notes              string   `json:"notes"`
	Recipient          *Contact `json:"recipient"`
	SendEmail          bool     `json:"sendEmail"`
	Sender             *Contact `json:"sender"`
	UTID               string   `json:"utid"`
}
