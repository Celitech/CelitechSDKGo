package api

import (
	"time"
)

// ListPackagesRequest holds the optional parameters for the API request.
type ListPackagesRequest struct {
	Destination      *string    `json:"-" url:"destination,omitempty"`
	DataLimitInGb    *float64   `json:"-" url:"dataLimitInGB,omitempty"`
	StartDate        *time.Time `json:"-" url:"startDate,omitempty" format:"date"`
	EndDate          *time.Time `json:"-" url:"endDate,omitempty" format:"date"`
	AfterCursor      *string    `json:"-" url:"afterCursor,omitempty"`
	Limit            *float64   `json:"-" url:"limit,omitempty"`
	StartTime        *int       `json:"-" url:"startTime,omitempty"`
	EndTime          *int       `json:"-" url:"endTime,omitempty"`
	IncludeUnlimited *bool      `json:"-" url:"includeUnlimited,omitempty"`
}

func (p *ListPackagesRequest) SetDestination(destination *string) {
	p.Destination = destination
}

func (p *ListPackagesRequest) SetDataLimitInGb(dataLimitInGb *float64) {
	p.DataLimitInGb = dataLimitInGb
}

func (p *ListPackagesRequest) SetStartDate(startDate *time.Time) {
	p.StartDate = startDate
}

func (p *ListPackagesRequest) SetEndDate(endDate *time.Time) {
	p.EndDate = endDate
}

func (p *ListPackagesRequest) SetAfterCursor(afterCursor *string) {
	p.AfterCursor = afterCursor
}

func (p *ListPackagesRequest) SetLimit(limit *float64) {
	p.Limit = limit
}

func (p *ListPackagesRequest) SetStartTime(startTime *int) {
	p.StartTime = startTime
}

func (p *ListPackagesRequest) SetEndTime(endTime *int) {
	p.EndTime = endTime
}

func (p *ListPackagesRequest) SetIncludeUnlimited(includeUnlimited *bool) {
	p.IncludeUnlimited = includeUnlimited
}

// ListPurchasesRequest holds the optional parameters for the API request.
type ListPurchasesRequest struct {
	PurchaseID  *string    `json:"-" url:"purchaseId,omitempty"`
	Iccid       *string    `json:"-" url:"iccid,omitempty"`
	AfterDate   *time.Time `json:"-" url:"afterDate,omitempty" format:"date"`
	BeforeDate  *time.Time `json:"-" url:"beforeDate,omitempty" format:"date"`
	Email       *string    `json:"-" url:"email,omitempty"`
	ReferenceID *string    `json:"-" url:"referenceId,omitempty"`
	AfterCursor *string    `json:"-" url:"afterCursor,omitempty"`
	Limit       *float64   `json:"-" url:"limit,omitempty"`
	After       *float64   `json:"-" url:"after,omitempty"`
	Before      *float64   `json:"-" url:"before,omitempty"`
}

func (p *ListPurchasesRequest) SetPurchaseID(purchaseID *string) {
	p.PurchaseID = purchaseID
}

func (p *ListPurchasesRequest) SetIccid(iccid *string) {
	p.Iccid = iccid
}

func (p *ListPurchasesRequest) SetAfterDate(afterDate *time.Time) {
	p.AfterDate = afterDate
}

func (p *ListPurchasesRequest) SetBeforeDate(beforeDate *time.Time) {
	p.BeforeDate = beforeDate
}

func (p *ListPurchasesRequest) SetEmail(email *string) {
	p.Email = email
}

func (p *ListPurchasesRequest) SetReferenceID(referenceID *string) {
	p.ReferenceID = referenceID
}

func (p *ListPurchasesRequest) SetAfterCursor(afterCursor *string) {
	p.AfterCursor = afterCursor
}

func (p *ListPurchasesRequest) SetLimit(limit *float64) {
	p.Limit = limit
}

func (p *ListPurchasesRequest) SetAfter(after *float64) {
	p.After = after
}

func (p *ListPurchasesRequest) SetBefore(before *float64) {
	p.Before = before
}

// GetPurchaseConsumptionRequest holds the optional parameters for the API request.
type GetPurchaseConsumptionRequest struct {
	PurchaseID string `json:"-" url:"-"`
}

func (p *GetPurchaseConsumptionRequest) SetPurchaseID(purchaseID string) {
	p.PurchaseID = purchaseID
}

// GetEsimRequest holds the optional parameters for the API request.
type GetEsimRequest struct {
	Iccid string `json:"-" url:"iccid"`
}

func (p *GetEsimRequest) SetIccid(iccid string) {
	p.Iccid = iccid
}

// GetEsimDeviceRequest holds the optional parameters for the API request.
type GetEsimDeviceRequest struct {
	Iccid string `json:"-" url:"-"`
}

func (p *GetEsimDeviceRequest) SetIccid(iccid string) {
	p.Iccid = iccid
}

// GetEsimHistoryRequest holds the optional parameters for the API request.
type GetEsimHistoryRequest struct {
	Iccid string `json:"-" url:"-"`
}

func (p *GetEsimHistoryRequest) SetIccid(iccid string) {
	p.Iccid = iccid
}

// SuspendEsimRequest holds the optional parameters for the API request.
type SuspendEsimRequest struct {
	Iccid          string  `json:"-" url:"-"`
	IdempotencyKey *string `json:"-" url:"-"`
}

func (p *SuspendEsimRequest) SetIccid(iccid string) {
	p.Iccid = iccid
}

func (p *SuspendEsimRequest) SetIdempotencyKey(idempotencyKey *string) {
	p.IdempotencyKey = idempotencyKey
}
