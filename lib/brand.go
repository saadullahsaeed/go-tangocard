package tangocard

// Status is a type to define the status of entities like brand or item
type Status string

const (
	// StatusActive is to identify that the entity is active
	StatusActive Status = "active"
)

// ValueType defines the type of value (fixed or variable) for card item
type ValueType string

const (
	// ValueTypeFixed is when the items have defined values
	ValueTypeFixed ValueType = "FIXED_VALUE"

	// ValueTypeVariable can have a min and max field for items
	ValueTypeVariable ValueType = "VARIABLE_VALUE"
)

// RewardType ...
type RewardType string

const (
	// RewardTypeGiftCard is for gift card types
	RewardTypeGiftCard RewardType = "gift card"
)

// Brand is the sponsor for the card
type Brand struct {
	Key              string            `json:"brandKey"`
	Name             string            `json:"brandName"`
	Disclaimer       string            `json:"disclaimer"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription"`
	Terms            string            `json:"terms"`
	ImageURLs        map[string]string `json:"imageUrls"`
	Status           Status            `json:"status"`
	Items            []Item            `json:"items"`
}

// Item represents a denomination of the gift card from the brand
type Item struct {
	UTID         string     `json:"utid"`
	RewardName   string     `json:"rewardName"`
	CurrencyCode string     `json:"currencyCode"`
	Status       Status     `json:"status"`
	ValueType    ValueType  `json:"valueType"`
	RewardType   RewardType `json:"rewardType"`
	FaceValue    float64    `json:"faceValue"`
	MinValue     float64    `json:"minValue"`
	MaxValue     float64    `json:"maxValue"`
	Countries    []string   `json:"countries"`
}
