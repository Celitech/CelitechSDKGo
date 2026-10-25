package api

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/v3/internal"
	"github.com/Celitech/CelitechSDKGo/v3/internal/unmarshal"
)

type ListPackagesResponse struct {
	Packages []*ListPackagesResponsePackagesItem `json:"packages" xml:"packages" required:"true"`
	// The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination
	AfterCursor     *string `json:"afterCursor" xml:"afterCursor" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListPackagesResponse) GetPackages() (value []*ListPackagesResponsePackagesItem) {
	if l == nil {
		return
	}
	return l.Packages
}

func (l *ListPackagesResponse) GetAfterCursor() (value *string) {
	if l == nil {
		return
	}
	return l.AfterCursor
}

func (l *ListPackagesResponse) SetPackages(packages []*ListPackagesResponsePackagesItem) {
	l.Packages = packages
}

func (l *ListPackagesResponse) SetAfterCursor(afterCursor *string) {
	l.AfterCursor = afterCursor
}

func (l *ListPackagesResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListPackagesResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListPackagesResponse to string"
	}
	return string(jsonData)
}

func (l *ListPackagesResponse) UnmarshalJSON(data []byte) error {
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

func (l *ListPackagesResponse) MarshalJSON() ([]byte, error) {
	type embed ListPackagesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
