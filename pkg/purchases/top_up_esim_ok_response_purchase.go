package purchases

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/unmarshal"
	"github.com/Celitech/CelitechSDKGo/pkg/util"
)

type TopUpEsimOkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty" required:"true"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty" required:"true"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *util.Nullable[string] `json:"startDate,omitempty" required:"true"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *util.Nullable[string] `json:"endDate,omitempty" required:"true"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty" required:"true"`
	// Epoch value representing the start time of the package's validity
	StartTime *util.Nullable[float64] `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *util.Nullable[float64] `json:"endTime,omitempty"`
}

func (t *TopUpEsimOkResponsePurchase) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TopUpEsimOkResponsePurchase) SetId(id string) {
	t.Id = &id
}

func (t *TopUpEsimOkResponsePurchase) GetPackageId() *string {
	if t == nil {
		return nil
	}
	return t.PackageId
}

func (t *TopUpEsimOkResponsePurchase) SetPackageId(packageId string) {
	t.PackageId = &packageId
}

func (t *TopUpEsimOkResponsePurchase) GetStartDate() *util.Nullable[string] {
	if t == nil {
		return nil
	}
	return t.StartDate
}

func (t *TopUpEsimOkResponsePurchase) SetStartDate(startDate util.Nullable[string]) {
	t.StartDate = &startDate
}

func (t *TopUpEsimOkResponsePurchase) SetStartDateNull() {
	t.StartDate = &util.Nullable[string]{IsNull: true}
}

func (t *TopUpEsimOkResponsePurchase) GetEndDate() *util.Nullable[string] {
	if t == nil {
		return nil
	}
	return t.EndDate
}

func (t *TopUpEsimOkResponsePurchase) SetEndDate(endDate util.Nullable[string]) {
	t.EndDate = &endDate
}

func (t *TopUpEsimOkResponsePurchase) SetEndDateNull() {
	t.EndDate = &util.Nullable[string]{IsNull: true}
}

func (t *TopUpEsimOkResponsePurchase) GetCreatedDate() *string {
	if t == nil {
		return nil
	}
	return t.CreatedDate
}

func (t *TopUpEsimOkResponsePurchase) SetCreatedDate(createdDate string) {
	t.CreatedDate = &createdDate
}

func (t *TopUpEsimOkResponsePurchase) GetStartTime() *util.Nullable[float64] {
	if t == nil {
		return nil
	}
	return t.StartTime
}

func (t *TopUpEsimOkResponsePurchase) SetStartTime(startTime util.Nullable[float64]) {
	t.StartTime = &startTime
}

func (t *TopUpEsimOkResponsePurchase) SetStartTimeNull() {
	t.StartTime = &util.Nullable[float64]{IsNull: true}
}

func (t *TopUpEsimOkResponsePurchase) GetEndTime() *util.Nullable[float64] {
	if t == nil {
		return nil
	}
	return t.EndTime
}

func (t *TopUpEsimOkResponsePurchase) SetEndTime(endTime util.Nullable[float64]) {
	t.EndTime = &endTime
}

func (t *TopUpEsimOkResponsePurchase) SetEndTimeNull() {
	t.EndTime = &util.Nullable[float64]{IsNull: true}
}

func (t TopUpEsimOkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponsePurchase to string"
	}
	return string(jsonData)
}

func (t *TopUpEsimOkResponsePurchase) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, t)
}
