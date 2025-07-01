package purchases

import "encoding/json"

type TopUpEsimOkResponse struct {
	Purchase *TopUpEsimOkResponsePurchase `json:"purchase,omitempty"`
	Profile  *TopUpEsimOkResponseProfile  `json:"profile,omitempty"`
}

func (t *TopUpEsimOkResponse) GetPurchase() *TopUpEsimOkResponsePurchase {
	if t == nil {
		return nil
	}
	return t.Purchase
}

func (t *TopUpEsimOkResponse) SetPurchase(purchase TopUpEsimOkResponsePurchase) {
	t.Purchase = &purchase
}

func (t *TopUpEsimOkResponse) GetProfile() *TopUpEsimOkResponseProfile {
	if t == nil {
		return nil
	}
	return t.Profile
}

func (t *TopUpEsimOkResponse) SetProfile(profile TopUpEsimOkResponseProfile) {
	t.Profile = &profile
}

func (t TopUpEsimOkResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponse to string"
	}
	return string(jsonData)
}

type TopUpEsimOkResponsePurchase struct {
	// ID of the purchase
	Id *string `json:"id,omitempty"`
	// ID of the package
	PackageId *string `json:"packageId,omitempty"`
	// Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	StartDate *string `json:"startDate,omitempty"`
	// End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'
	EndDate *string `json:"endDate,omitempty"`
	// Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'
	CreatedDate *string `json:"createdDate,omitempty"`
	// Epoch value representing the start time of the package's validity
	StartTime *float64 `json:"startTime,omitempty"`
	// Epoch value representing the end time of the package's validity
	EndTime *float64 `json:"endTime,omitempty"`
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

func (t *TopUpEsimOkResponsePurchase) GetStartDate() *string {
	if t == nil {
		return nil
	}
	return t.StartDate
}

func (t *TopUpEsimOkResponsePurchase) SetStartDate(startDate string) {
	t.StartDate = &startDate
}

func (t *TopUpEsimOkResponsePurchase) GetEndDate() *string {
	if t == nil {
		return nil
	}
	return t.EndDate
}

func (t *TopUpEsimOkResponsePurchase) SetEndDate(endDate string) {
	t.EndDate = &endDate
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

func (t *TopUpEsimOkResponsePurchase) GetStartTime() *float64 {
	if t == nil {
		return nil
	}
	return t.StartTime
}

func (t *TopUpEsimOkResponsePurchase) SetStartTime(startTime float64) {
	t.StartTime = &startTime
}

func (t *TopUpEsimOkResponsePurchase) GetEndTime() *float64 {
	if t == nil {
		return nil
	}
	return t.EndTime
}

func (t *TopUpEsimOkResponsePurchase) SetEndTime(endTime float64) {
	t.EndTime = &endTime
}

func (t TopUpEsimOkResponsePurchase) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponsePurchase to string"
	}
	return string(jsonData)
}

type TopUpEsimOkResponseProfile struct {
	// ID of the eSIM
	Iccid *string `json:"iccid,omitempty" maxLength:"22" minLength:"18"`
}

func (t *TopUpEsimOkResponseProfile) GetIccid() *string {
	if t == nil {
		return nil
	}
	return t.Iccid
}

func (t *TopUpEsimOkResponseProfile) SetIccid(iccid string) {
	t.Iccid = &iccid
}

func (t TopUpEsimOkResponseProfile) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopUpEsimOkResponseProfile to string"
	}
	return string(jsonData)
}
