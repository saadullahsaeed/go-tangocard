package tangocard_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/saadullahsaeed/go-tangocard/lib"
)

func TestAuth(t *testing.T) {
	var auth string

	tests := []struct {
		username     string
		password     string
		expectedAuth string
	}{
		{
			username:     "test",
			password:     "test",
			expectedAuth: "Basic dGVzdDp0ZXN0",
		},
	}

	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth = r.Header.Get("Authorization")
		}))

		client := tangocard.NewClient(
			tt.username,
			tt.password,
			&tangocard.ClientOptions{Host: ts.URL},
		)
		client.GetCatalog()
		ts.Close()

		if auth != tt.expectedAuth {
			t.FailNow()
		}
	}
}

func TestClient_GetCatalog(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
	"catalogName": "test catalog",
	"brands": [{
		"brandKey": "test_brand",
		"brandName": "In n out gift cards",
		"disclaimer": "test disclaimer"
	}]
}`)
	}))
	defer ts.Close()

	client := tangocard.NewClient(
		"",
		"",
		&tangocard.ClientOptions{Host: ts.URL},
	)

	cr, err := client.GetCatalog()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expectedRes := &tangocard.GetCatalogResponse{
		CatalogName: "test catalog",
		Brands: []*tangocard.Brand{
			&tangocard.Brand{
				Key:        "test_brand",
				Name:       "In n out gift cards",
				Disclaimer: "test disclaimer",
			},
		},
	}
	if !reflect.DeepEqual(expectedRes, cr) {
		t.FailNow()
	}
}

func TestClient_GetAccountDetails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-request-id", "4477b71e-85d0-45aa-b1af-9f1639080166")
		http.Error(w, `{
  "timestamp": "2017-09-08T19:00:59.516Z",
  "requestId": "4477b71e-85d0-45aa-b1af-9f1639080166",
  "path": "/raas/v2/accounts/x",
  "httpCode": 404,
  "httpPhrase": "Not Found",
  "errors": [
    {
      "message": "The account does not exist for this platform.",
      "code": 404001
    }
  ]
}`, http.StatusNotFound)

	}))

	defer ts.Close()

	client := tangocard.NewClient(
		"",
		"",
		&tangocard.ClientOptions{Host: ts.URL},
	)

	_, err := client.GetAccountDetails("saadtest")
	if err == nil {
		t.Error("expected an error but got none")
	}

	if err.Error() != "\nThe account does not exist for this platform." {
		t.FailNow()
	}

	if err.(*tangocard.ErrorResponse).RequestID != "4477b71e-85d0-45aa-b1af-9f1639080166" {
		t.FailNow()
	}
}
