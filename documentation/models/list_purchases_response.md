# ListPurchasesResponse

**Properties**

| Name        | Type                                       | Required | Description                                                                                                                                                                                                                                                                                     |
| :---------- | :----------------------------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Purchases   | []\*api.ListPurchasesResponsePurchasesItem | ✅       |                                                                                                                                                                                                                                                                                                 |
| AfterCursor | string                                     | ✅       | The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination. |

# ListPurchasesResponsePurchasesItem

**Properties**

| Name         | Type                                          | Required | Description                                                                                                                                                                                                    |
| :----------- | :-------------------------------------------- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ID           | string                                        | ✅       | ID of the purchase                                                                                                                                                                                             |
| StartDate    | time.Time                                     | ✅       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                     |
| EndDate      | time.Time                                     | ✅       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                       |
| CreatedDate  | time.Time                                     | ✅       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                            |
| Package      | api.ListPurchasesResponsePurchasesItemPackage | ✅       |                                                                                                                                                                                                                |
| Esim         | api.ListPurchasesResponsePurchasesItemEsim    | ✅       |                                                                                                                                                                                                                |
| Source       | string                                        | ✅       | The `source` indicates whether the purchase was made from the API, dashboard, landing-page, promo-page or iframe. For purchases made before September 8, 2023, the value will be displayed as 'Not available'. |
| PurchaseType | string                                        | ✅       | The `purchaseType` indicates whether this is the initial purchase that creates the eSIM (First Purchase) or a subsequent top-up on an existing eSIM (Top-up Purchase).                                         |
| Duration     | float64                                       | ❌       | Duration of the package in days. Possible values are 1, 2, 7, 14, 30, or 90. `null` for unlimited (date-based) packages.                                                                                       |
| StartTime    | float64                                       | ❌       | Epoch value representing the start time of the package's validity                                                                                                                                              |
| EndTime      | float64                                       | ❌       | Epoch value representing the end time of the package's validity                                                                                                                                                |
| CreatedAt    | float64                                       | ❌       | Epoch value representing the date of creation of the purchase                                                                                                                                                  |
| ReferenceID  | string                                        | ❌       | The `referenceId` that was provided by the partner during the purchase or top-up flow. This identifier can be used for analytics and debugging purposes.                                                       |

# ListPurchasesResponsePurchasesItemPackage

**Properties**

| Name             | Type    | Required | Description                                                                             |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------- |
| ID               | string  | ✅       | ID of the package                                                                       |
| DataLimitInBytes | float64 | ✅       | Size of the package in Bytes. A value of `-1` indicates an unlimited package.           |
| DataLimitInGb    | float64 | ✅       | Size of the package in GB. A value of `-1` indicates an unlimited (date-based) package. |
| Destination      | string  | ✅       | ISO3 representation of the package's destination.                                       |
| DestinationIso2  | string  | ✅       | ISO2 representation of the package's destination.                                       |
| DestinationName  | string  | ✅       | Name of the package's destination                                                       |
| PriceInCents     | float64 | ✅       | Price of the package in cents                                                           |

# ListPurchasesResponsePurchasesItemEsim

**Properties**

| Name  | Type   | Required | Description    |
| :---- | :----- | :------- | :------------- |
| Iccid | string | ✅       | ID of the eSIM |
