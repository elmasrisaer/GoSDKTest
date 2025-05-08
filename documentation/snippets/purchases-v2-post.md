```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"
  "github.com/elmasrisaer/GoSDKTest/pkg/util"
  "github.com/elmasrisaer/GoSDKTest/pkg/purchases"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


request := purchases.CreatePurchaseV2Request{
  Destination: util.ToPointer("Destination"),
  DataLimitInGb: util.ToPointer(float64(123)),
  Quantity: util.ToPointer(float64(123)),
}

response, err := client.Purchases.CreatePurchaseV2(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
