package purchases

import "encoding/json"

type TopUpEsimOkResponse struct {
	Purchase *TopUpEsimOkResponsePurchase `json:"purchase,omitempty" required:"true"`
	Profile  *TopUpEsimOkResponseProfile  `json:"profile,omitempty" required:"true"`
}

func (t *TopUpEsimOkResponse) GetPurchase() *TopUpEsimOkResponsePurchase {
	if t == nil {
		return nil
	}
	return t.Purchase
}

func (t *TopUpEsimOkResponse) SetPurchase(purchase TopUpEsimOkResponsePurchase) {
	t.Purchase = &purchase
}

func (t *TopUpEsimOkResponse) GetProfile() *TopUpEsimOkResponseProfile {
	if t == nil {
		return nil
	}
	return t.Profile
}

func (t *TopUpEsimOkResponse) SetProfile(profile TopUpEsimOkResponseProfile) {
	t.Profile = &profile
}

func (t TopUpEsimOkResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponse to string"
	}
	return string(jsonData)
}
