package packages

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type ListPackagesOkResponse struct {
	Packages []Packages `json:"packages,omitempty" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination
	AfterCursor *util.Nullable[string] `json:"afterCursor,omitempty" required:"true"`
}

func (l *ListPackagesOkResponse) GetPackages() []Packages {
	if l == nil {
		return nil
	}
	return l.Packages
}

func (l *ListPackagesOkResponse) SetPackages(packages []Packages) {
	l.Packages = packages
}

func (l *ListPackagesOkResponse) GetAfterCursor() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.AfterCursor
}

func (l *ListPackagesOkResponse) SetAfterCursor(afterCursor util.Nullable[string]) {
	l.AfterCursor = &afterCursor
}

func (l *ListPackagesOkResponse) SetAfterCursorNull() {
	l.AfterCursor = &util.Nullable[string]{IsNull: true}
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
