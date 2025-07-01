# EditPurchaseOkResponse

**Properties**

| Name         | Type    | Required | Description                                                                |
| :----------- | :------ | :------- | :------------------------------------------------------------------------- |
| PurchaseId   | string  | ❌       | ID of the purchase                                                         |
| NewStartDate | string  | ❌       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ' |
| NewEndDate   | string  | ❌       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'   |
| NewStartTime | float64 | ❌       | Epoch value representing the new start time of the package's validity      |
| NewEndTime   | float64 | ❌       | Epoch value representing the new end time of the package's validity        |
