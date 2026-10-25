# TopUpEsimResponse

**Properties**

| Name     | Type                          | Required | Description |
| :------- | :---------------------------- | :------- | :---------- |
| Purchase | api.TopUpEsimResponsePurchase | ✅       |             |
| Profile  | api.TopUpEsimResponseProfile  | ✅       |             |

# TopUpEsimResponsePurchase

**Properties**

| Name        | Type      | Required | Description                                                                |
| :---------- | :-------- | :------- | :------------------------------------------------------------------------- |
| ID          | string    | ✅       | ID of the purchase                                                         |
| PackageID   | string    | ✅       | ID of the package                                                          |
| StartDate   | time.Time | ✅       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ' |
| EndDate     | time.Time | ✅       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'   |
| CreatedDate | time.Time | ✅       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'        |
| StartTime   | float64   | ❌       | Epoch value representing the start time of the package's validity          |
| EndTime     | float64   | ❌       | Epoch value representing the end time of the package's validity            |

# TopUpEsimResponseProfile

**Properties**

| Name  | Type   | Required | Description    |
| :---- | :----- | :------- | :------------- |
| Iccid | string | ✅       | ID of the eSIM |
