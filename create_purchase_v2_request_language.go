package api

import (
	"fmt"
)

// Language of the confirmation email sent to the customer.
type CreatePurchaseV2RequestLanguage string

const (
	CreatePurchaseV2RequestLanguageEn   CreatePurchaseV2RequestLanguage = "en"
	CreatePurchaseV2RequestLanguageEs   CreatePurchaseV2RequestLanguage = "es"
	CreatePurchaseV2RequestLanguageFr   CreatePurchaseV2RequestLanguage = "fr"
	CreatePurchaseV2RequestLanguageDe   CreatePurchaseV2RequestLanguage = "de"
	CreatePurchaseV2RequestLanguagePtBr CreatePurchaseV2RequestLanguage = "pt-br"
)

func NewCreatePurchaseV2RequestLanguageFromString(s string) (CreatePurchaseV2RequestLanguage, error) {
	switch s {
	case "en":
		return CreatePurchaseV2RequestLanguageEn, nil
	case "es":
		return CreatePurchaseV2RequestLanguageEs, nil
	case "fr":
		return CreatePurchaseV2RequestLanguageFr, nil
	case "de":
		return CreatePurchaseV2RequestLanguageDe, nil
	case "pt-br":
		return CreatePurchaseV2RequestLanguagePtBr, nil
	}
	var t CreatePurchaseV2RequestLanguage
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreatePurchaseV2RequestLanguage) Ptr() *CreatePurchaseV2RequestLanguage {
	return &e
}
