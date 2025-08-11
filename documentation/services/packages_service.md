# PackagesService

A list of all methods in the `PackagesService` service. Click on the method name to view detailed information about that method.

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
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

  "github.com/Celitech/CelitechSDKGo/pkg/packages"
)

config := celitechconfig.NewConfig()
config.SetClientId("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := packages.ListPackagesRequestParams{

}

response, err := client.Packages.ListPackages(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
