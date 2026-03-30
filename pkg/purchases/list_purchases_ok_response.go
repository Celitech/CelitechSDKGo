package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type ListPurchasesOkResponse struct {
	Purchases []Purchases `json:"purchases,omitempty" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination.
	AfterCursor *util.Nullable[string] `json:"afterCursor,omitempty" required:"true"`
}

func (l *ListPurchasesOkResponse) GetPurchases() []Purchases {
	if l == nil {
		return nil
	}
	return l.Purchases
}

func (l *ListPurchasesOkResponse) SetPurchases(purchases []Purchases) {
	l.Purchases = purchases
}

func (l *ListPurchasesOkResponse) GetAfterCursor() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.AfterCursor
}

func (l *ListPurchasesOkResponse) SetAfterCursor(afterCursor util.Nullable[string]) {
	l.AfterCursor = &afterCursor
}

func (l *ListPurchasesOkResponse) SetAfterCursorNull() {
	l.AfterCursor = &util.Nullable[string]{IsNull: true}
}

func (l ListPurchasesOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPurchasesOkResponse to string"
	}
	return string(jsonData)
}

func (l *ListPurchasesOkResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}
