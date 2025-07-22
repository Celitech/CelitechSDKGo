```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"
  "github.com/Celitech/CelitechSDKGo/pkg/util"
  "github.com/Celitech/CelitechSDKGo/pkg/oauth"
)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BaseUrl")
config.SetTimeout("Timeout")
client := celitech.NewCelitech(config)


request := oauth.GetAccessTokenRequest{

}

response, err := client.OAuth.GetAccessToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
