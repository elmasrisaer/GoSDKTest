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
