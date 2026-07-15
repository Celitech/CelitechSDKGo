# GetPurchaseConsumptionOkResponse

**Properties**

| Name                      | Type    | Required | Description                                                                     |
| :------------------------ | :------ | :------- | :------------------------------------------------------------------------------ |
| DataUsageRemainingInBytes | float64 | ✅       | Remaining balance of the package in bytes. Returns `-1` for unlimited packages. |
| DataUsageRemainingInGb    | float64 | ✅       | Remaining balance of the package in GB. Returns `-1` for unlimited packages.    |
| Status                    | string  | ✅       | Status of the connectivity, possible values are 'ACTIVE' or 'NOT_ACTIVE'        |
