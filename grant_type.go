package api

import (
	"fmt"
)

type GrantType string

const (
	GrantTypeClientCredentials GrantType = "client_credentials"
)

func NewGrantTypeFromString(s string) (GrantType, error) {
	switch s {
	case "client_credentials":
		return GrantTypeClientCredentials, nil
	}
	var t GrantType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GrantType) Ptr() *GrantType {
	return &e
}
