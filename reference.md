# Reference

## Destinations

### `ListDestinations`

`GET /destinations`

List Destinations

## Packages

### `ListPackages`

`GET /packages`

List Packages

**Parameters:**

- `destination`
- `dataLimitInGb`
- `startDate`
- `endDate`
- `afterCursor`
- `limit`
- `startTime`
- `endTime`
- `includeUnlimited`

## Purchases

### `CreatePurchaseV2`

`POST /purchases/v2`

Create Purchase V2

### `ListPurchases`

`GET /purchases`

List Purchases

**Parameters:**

- `purchaseID`
- `iccid`
- `afterDate`
- `beforeDate`
- `email`
- `referenceID`
- `afterCursor`
- `limit`
- `after`
- `before`

### `CreatePurchase`

`POST /purchases`

Create Purchase

### `TopUpEsim`

`POST /purchases/topup`

Top-up eSIM

### `EditPurchase`

`POST /purchases/edit`

Edit Purchase

### `GetPurchaseConsumption`

`GET /purchases/{purchaseId}/consumption`

Get Purchase Consumption

**Parameters:**

- `purchaseID` (required)

## ESim

### `GetEsim`

`GET /esim`

Get eSIM

**Parameters:**

- `iccid` (required)

### `GetEsimDevice`

`GET /esim/{iccid}/device`

Get eSIM Device

**Parameters:**

- `iccid` (required)

### `GetEsimHistory`

`GET /esim/{iccid}/history`

Get eSIM History

**Parameters:**

- `iccid` (required)

### `SuspendEsim`

`POST /esim/{iccid}/suspend`

Suspend eSIM

**Parameters:**

- `iccid` (required)

## IFrame

### `Token`

`POST /iframe/token`

Generate Token

## OAuth

### `GetAccessToken`

`POST /oauth2/token`
