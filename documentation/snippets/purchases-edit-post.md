```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"
  "github.com/Celitech/CelitechSDKGo/pkg/util"
  "github.com/Celitech/CelitechSDKGo/pkg/purchases"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


request := purchases.EditPurchaseRequest{
  PurchaseId: util.ToPointer("PurchaseId"),
  StartDate: util.ToPointer("StartDate"),
  EndDate: util.ToPointer("EndDate"),
}

response, err := client.Purchases.EditPurchase(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
