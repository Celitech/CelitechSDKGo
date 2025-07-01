# CreatePurchaseV2OkResponse

**Properties**

| Name     | Type                                         | Required | Description |
| :------- | :------------------------------------------- | :------- | :---------- |
| Purchase | purchases.CreatePurchaseV2OkResponsePurchase | ❌       |             |
| Profile  | purchases.CreatePurchaseV2OkResponseProfile  | ❌       |             |

# CreatePurchaseV2OkResponsePurchase

**Properties**

| Name        | Type   | Required | Description                                                         |
| :---------- | :----- | :------- | :------------------------------------------------------------------ |
| Id          | string | ❌       | ID of the purchase                                                  |
| PackageId   | string | ❌       | ID of the package                                                   |
| CreatedDate | string | ❌       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ' |

# CreatePurchaseV2OkResponseProfile

**Properties**

| Name                 | Type   | Required | Description                        |
| :------------------- | :----- | :------- | :--------------------------------- |
| Iccid                | string | ❌       | ID of the eSIM                     |
| ActivationCode       | string | ❌       | QR Code of the eSIM as base64      |
| ManualActivationCode | string | ❌       | Manual Activation Code of the eSIM |
