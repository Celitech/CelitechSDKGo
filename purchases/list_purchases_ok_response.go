package purchases

import (
	"encoding/json"
	"example.com/celitech/internal/unmarshal"
	"example.com/celitech/param"
)

type ListPurchasesOkResponse struct {
	Purchases []Purchases `json:"purchases" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination.
	AfterCursor *param.Nullable[string] `json:"afterCursor" required:"true"`
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
