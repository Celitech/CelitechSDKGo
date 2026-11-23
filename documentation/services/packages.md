# packages

A list of all methods in the `packages` service. Click on the method name to view detailed information about that method.

| Methods                       | Description   |
| :---------------------------- | :------------ |
| [ListPackages](#listpackages) | List Packages |

## ListPackages

List Packages

- HTTP Method: `GET`
- Endpoint: `/packages`

**Parameters**

| Name             | Type      | Required | Description                                                                                                                                                                                                                                                                                          |
| :--------------- | :-------- | :------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context   | ✅        | Default go language context                                                                                                                                                                                                                                                                          |
| Destination      | string    | ❌        | ISO representation of the package's destination. Supports both ISO2 (e.g., 'FR') and ISO3 (e.g., 'FRA') country codes.                                                                                                                                                                               |
| DataLimitInGb    | float64   | ❌        | Filter packages by data limit in GB. When provided, only packages with this exact data limit are returned. Use `-1` together with `includeUnlimited=true` to return only unlimited packages. A value of `0` is ignored.                                                                              |
| StartDate        | time.Time | ❌        | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.                                                                                                                                                       |
| EndDate          | time.Time | ❌        | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                                                                                                                                                                     |
| AfterCursor      | string    | ❌        | To get the next batch of results, use this parameter. It tells the API where to start fetching data after the last item you received. It helps you avoid repeats and efficiently browse through large sets of data.                                                                                  |
| Limit            | float64   | ❌        | Maximum number of packages to be returned in the response. The value must be greater than 0 and less than or equal to 160. If not provided, the default value is 20                                                                                                                                  |
| StartTime        | int       | ❌        | Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months                                                                                                                                               |
| EndTime          | int       | ❌        | Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time                                                                                                                                                                                    |
| IncludeUnlimited | bool      | ❌        | Whether to include unlimited (date-based) packages in the results. Unlimited packages are excluded by default; set this to `true` to include them. An unlimited package has `dataLimitInGB` and `dataLimitInBytes` equal to `-1`, and is offered for 3 to 30 days with `minDays` equal to `maxDays`. |

**Return Type**

[`ListPackagesResponse`](../models/list_packages_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Packages.ListPackages(context.Background(), &api.ListPackagesRequest{})
if err != nil {
  panic(err)
}

fmt.Println(response)
```
