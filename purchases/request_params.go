package purchases

import (
	"github.com/Celitech/CelitechSDKGo/param"
)

// CreatePurchaseRequestParams holds the optional parameters for the API request.
type CreatePurchaseRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

// ListPurchasesRequestParams holds the optional parameters for the API request.
type ListPurchasesRequestParams struct {
	PurchaseID  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"purchaseId"`
	Iccid       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"iccid"`
	AfterDate   *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"afterDate"`
	BeforeDate  *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"beforeDate"`
	Email       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"email"`
	ReferenceID *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"referenceId"`
	AfterCursor *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"afterCursor"`
	Limit       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"limit"`
	After       *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"after"`
	Before      *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"before"`
	Accept      *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
