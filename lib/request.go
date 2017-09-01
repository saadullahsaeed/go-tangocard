package tangocard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	defaultIntegrationHost = "https://integration-api.tangocard.com/raas/v2"
)

// GetOrdersRequest represents the request to get a list of all orders
type GetOrdersRequest struct {
	AccountIdentifier  string     `json:"accountIdentifier"`  //which account to query
	CustomerIdentifier string     `json:"customerIdentifier"` //which custtomer to query
	ExternalRefID      string     `json:"externalRefID"`      //
	StartDate          *time.Time `json:"startDate"`
	EndDate            *time.Time `json:"endDate"`
	ElementsPerBlock   int        `json:"elementsPerBlock"`
	Page               int        `json:"page"`
}

// sendRequest
func sendRequest(host, method, url, username, password string, body interface{}) (json.RawMessage, error) {
	if len(host) == 0 {
		host = defaultIntegrationHost
	}

	jstr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jstr))
	client := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", host, url), bytes.NewBuffer(jstr))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyjson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		err := &ErrorResponse{}
		json.Unmarshal(bodyjson, err)
		return json.RawMessage(bodyjson), err
	}
	return json.RawMessage(bodyjson), nil
}
