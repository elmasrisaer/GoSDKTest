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


request := purchases.EditPurchaseRequest{}
request.SetPurchaseId("PurchaseId")
request.SetStartDate("StartDate")
request.SetEndDate("EndDate")

response, err := client.Purchases.EditPurchase(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
