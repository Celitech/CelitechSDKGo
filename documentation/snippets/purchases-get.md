```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

  "github.com/Celitech/CelitechSDKGo/pkg/purchases"
)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BaseUrl")
config.SetTimeout("Timeout")
client := celitech.NewCelitech(config)


params := purchases.ListPurchasesRequestParams{

}

response, err := client.Purchases.ListPurchases(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
