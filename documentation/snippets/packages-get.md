```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

  "github.com/Celitech/CelitechSDKGo/pkg/packages"
)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BASE_URL")
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
