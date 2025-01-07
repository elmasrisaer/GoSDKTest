```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"
  "github.com/elmasrisaer/GoSDKTest/pkg/purchases"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


params := purchases.ListPurchasesRequestParams{}


response, err := client.Purchases.ListPurchases(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
