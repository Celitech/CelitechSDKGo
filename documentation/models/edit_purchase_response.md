# EditPurchaseResponse

**Properties**

| Name         | Type      | Required | Description                                                                |
| :----------- | :-------- | :------- | :------------------------------------------------------------------------- |
| PurchaseID   | string    | ✅       | ID of the purchase                                                         |
| NewStartDate | time.Time | ✅       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ' |
| NewEndDate   | time.Time | ✅       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'   |
| NewStartTime | float64   | ❌       | Epoch value representing the new start time of the package's validity      |
| NewEndTime   | float64   | ❌       | Epoch value representing the new end time of the package's validity        |
