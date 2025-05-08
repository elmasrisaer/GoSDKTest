# PurchasesService

A list of all methods in the `PurchasesService` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description                                                                                                                                                                                                                                                                                                                                                      |
| :------------------------------------------------ | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [CreatePurchaseV2](#createpurchasev2)             | This endpoint is used to purchase a new eSIM by providing the package details.                                                                                                                                                                                                                                                                                   |
| [ListPurchases](#listpurchases)                   | This endpoint can be used to list all the successful purchases made between a given interval.                                                                                                                                                                                                                                                                    |
| [TopUpEsim](#topupesim)                           | This endpoint is used to top-up an eSIM with the previously associated destination by providing an existing ICCID and the package details. The top-up is only feasible for eSIMs in "ENABLED" or "INSTALLED" state. You can check this state using the Get eSIM Status endpoint.                                                                                 |
| [EditPurchase](#editpurchase)                     | This endpoint allows you to modify the dates of an existing package with a future activation start time. Editing can only be performed for packages that have not been activated, and it cannot change the package size. The modification must not change the package duration category to ensure pricing consistency. Duration based packages cannot be edited. |
| [GetPurchaseConsumption](#getpurchaseconsumption) | This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.                                                                                                                                                                                 |

## CreatePurchaseV2

This endpoint is used to purchase a new eSIM by providing the package details.

- HTTP Method: `POST`
- Endpoint: `/purchases/v2`

**Parameters**

| Name                    | Type                    | Required | Description                 |
| :---------------------- | :---------------------- | :------- | :-------------------------- |
| ctx                     | Context                 | ✅       | Default go language context |
| createPurchaseV2Request | CreatePurchaseV2Request | ✅       |                             |

**Return Type**

`[]CreatePurchaseV2OkResponse`

**Example Usage Code Snippet**

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

## ListPurchases

This endpoint can be used to list all the successful purchases made between a given interval.

- HTTP Method: `GET`
- Endpoint: `/purchases`

**Parameters**

| Name   | Type                       | Required | Description                   |
| :----- | :------------------------- | :------- | :---------------------------- |
| ctx    | Context                    | ✅       | Default go language context   |
| params | ListPurchasesRequestParams | ✅       | Additional request parameters |

**Return Type**

`ListPurchasesOkResponse`

**Example Usage Code Snippet**

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


params := purchases.ListPurchasesRequestParams{

}

response, err := client.Purchases.ListPurchases(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## TopUpEsim

This endpoint is used to top-up an eSIM with the previously associated destination by providing an existing ICCID and the package details. The top-up is only feasible for eSIMs in "ENABLED" or "INSTALLED" state. You can check this state using the Get eSIM Status endpoint.

- HTTP Method: `POST`
- Endpoint: `/purchases/topup`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| topUpEsimRequest | TopUpEsimRequest | ✅       |                             |

**Return Type**

`TopUpEsimOkResponse`

**Example Usage Code Snippet**

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


request := purchases.TopUpEsimRequest{
  Iccid: util.ToPointer("Iccid"),
  DataLimitInGb: util.ToPointer(float64(123)),
}

response, err := client.Purchases.TopUpEsim(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditPurchase

This endpoint allows you to modify the dates of an existing package with a future activation start time. Editing can only be performed for packages that have not been activated, and it cannot change the package size. The modification must not change the package duration category to ensure pricing consistency. Duration based packages cannot be edited.

- HTTP Method: `POST`
- Endpoint: `/purchases/edit`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| editPurchaseRequest | EditPurchaseRequest | ✅       |                             |

**Return Type**

`EditPurchaseOkResponse`

**Example Usage Code Snippet**

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

## GetPurchaseConsumption

This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.

- HTTP Method: `GET`
- Endpoint: `/purchases/{purchaseId}/consumption`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| purchaseId | string  | ✅       | ID of the purchase          |

**Return Type**

`GetPurchaseConsumptionOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"

)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)

response, err := client.Purchases.GetPurchaseConsumption(context.Background(), "purchaseId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
