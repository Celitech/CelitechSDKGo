package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type ListPurchasesResponse struct {
	Purchases []*ListPurchasesResponsePurchasesItem `json:"purchases" xml:"purchases" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination.
	AfterCursor     *string `json:"afterCursor" xml:"afterCursor" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListPurchasesResponse) GetPurchases() (value []*ListPurchasesResponsePurchasesItem) {
	if l == nil {
		return
	}
	return l.Purchases
}

func (l *ListPurchasesResponse) GetAfterCursor() (value *string) {
	if l == nil {
		return
	}
	return l.AfterCursor
}

func (l *ListPurchasesResponse) SetPurchases(purchases []*ListPurchasesResponsePurchasesItem) {
	l.Purchases = purchases
}

func (l *ListPurchasesResponse) SetAfterCursor(afterCursor *string) {
	l.AfterCursor = afterCursor
}

func (l *ListPurchasesResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListPurchasesResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPurchasesResponse to string"
	}
	return string(jsonData)
}

func (l *ListPurchasesResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListPurchasesResponse) MarshalJSON() ([]byte, error) {
	type embed ListPurchasesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
