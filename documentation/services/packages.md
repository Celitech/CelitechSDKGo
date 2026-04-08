# Packages

A list of all methods in the `Packages` service. Click on the method name to view detailed information about that method.

| Methods                       | Description   |
| :---------------------------- | :------------ |
| [ListPackages](#listpackages) | List Packages |

## ListPackages

List Packages

- HTTP Method: `GET`
- Endpoint: `/packages`

**Parameters**

| Name   | Type                      | Required | Description                   |
| :----- | :------------------------ | :------- | :---------------------------- |
| ctx    | Context                   | ✅       | Default go language context   |
| params | ListPackagesRequestParams | ✅       | Additional request parameters |

**Return Type**

`ListPackagesOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "example.com/celitech"
  "example.com/celitech/packages"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := packages.ListPackagesRequestParams{
  Destination: celitech.Ptr("FRA"),
  StartDate: celitech.Ptr("2023-11-01"),
  EndDate: celitech.Ptr("2023-11-20"),
  AfterCursor: celitech.Ptr("Y3JlYXRlZEF0OjE1OTk0OTMwOTgsZGVzdGluYXRpb246QVVTLG1pbkRheXM6MCxkYXRhTGltaXRJbkJ5dGVzOjUzNjg3MDkxMjA"),
  Limit: celitech.Ptr(float64(20)),
  StartTime: celitech.Ptr(int64(2)),
  EndTime: celitech.Ptr(int64(0)),
}

response, err := client.Packages.ListPackages(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
