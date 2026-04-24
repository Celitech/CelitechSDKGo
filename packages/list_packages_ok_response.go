package packages

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/param"
)

type ListPackagesOkResponse struct {
	Packages []Packages `json:"packages" xml:"packages" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination
	AfterCursor *param.Nullable[string] `json:"afterCursor" xml:"afterCursor" required:"true"`
}

func (l ListPackagesOkResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPackagesOkResponse to string"
	}
	return string(jsonData)
}

func (l *ListPackagesOkResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}
