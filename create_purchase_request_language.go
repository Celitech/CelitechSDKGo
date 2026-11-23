package api

import (
	"fmt"
)

// Language of the confirmation email sent to the customer.
type CreatePurchaseRequestLanguage string

const (
	CreatePurchaseRequestLanguageEn   CreatePurchaseRequestLanguage = "en"
	CreatePurchaseRequestLanguageEs   CreatePurchaseRequestLanguage = "es"
	CreatePurchaseRequestLanguageFr   CreatePurchaseRequestLanguage = "fr"
	CreatePurchaseRequestLanguageDe   CreatePurchaseRequestLanguage = "de"
	CreatePurchaseRequestLanguagePtBr CreatePurchaseRequestLanguage = "pt-br"
)

func NewCreatePurchaseRequestLanguageFromString(s string) (CreatePurchaseRequestLanguage, error) {
	switch s {
	case "en":
		return CreatePurchaseRequestLanguageEn, nil
	case "es":
		return CreatePurchaseRequestLanguageEs, nil
	case "fr":
		return CreatePurchaseRequestLanguageFr, nil
	case "de":
		return CreatePurchaseRequestLanguageDe, nil
	case "pt-br":
		return CreatePurchaseRequestLanguagePtBr, nil
	}
	var t CreatePurchaseRequestLanguage
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreatePurchaseRequestLanguage) Ptr() *CreatePurchaseRequestLanguage {
	return &e
}
