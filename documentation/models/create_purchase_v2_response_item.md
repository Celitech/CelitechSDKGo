# CreatePurchaseV2ResponseItem

**Properties**

| Name     | Type                                     | Required | Description |
| :------- | :--------------------------------------- | :------- | :---------- |
| Purchase | api.CreatePurchaseV2ResponseItemPurchase | ✅       |             |
| Profile  | api.CreatePurchaseV2ResponseItemProfile  | ✅       |             |

# CreatePurchaseV2ResponseItemPurchase

**Properties**

| Name        | Type      | Required | Description                                                         |
| :---------- | :-------- | :------- | :------------------------------------------------------------------ |
| ID          | string    | ✅       | ID of the purchase                                                  |
| PackageID   | string    | ✅       | ID of the package                                                   |
| CreatedDate | time.Time | ✅       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ' |

# CreatePurchaseV2ResponseItemProfile

**Properties**

| Name                  | Type   | Required | Description                         |
| :-------------------- | :----- | :------- | :---------------------------------- |
| Iccid                 | string | ✅       | ID of the eSIM                      |
| ActivationCode        | string | ✅       | QR Code of the eSIM as base64       |
| ManualActivationCode  | string | ✅       | Manual Activation Code of the eSIM  |
| IosActivationLink     | string | ✅       | iOS Activation Link of the eSIM     |
| AndroidActivationLink | string | ✅       | Android Activation Link of the eSIM |
