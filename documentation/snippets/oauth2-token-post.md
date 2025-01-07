```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"
  "github.com/elmasrisaer/GoSDKTest/pkg/oauth"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


request := oauth.GetAccessTokenRequest{}

response, err := client.OAuth.GetAccessToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
