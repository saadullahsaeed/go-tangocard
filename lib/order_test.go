package tangocard_test

import (
	"testing"

	tangocard "github.com/saadullahsaeed/go-tangocard/lib"
)

func TestReward_HasRedemptionURL(t *testing.T) {
	r := &tangocard.Reward{}

	if r.HasRedemptionURL() {
		t.FailNow()
	}

	r.Credentials = map[string]string{
		"Redemption URL": "x",
	}
	if !r.HasRedemptionURL() {
		t.FailNow()
	}
}
