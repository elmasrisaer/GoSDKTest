# ListPurchasesOkResponse

**Properties**

| Name        | Type                  | Required | Description                                                                                                                                                                                                                                                                                     |
| :---------- | :-------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Purchases   | []purchases.Purchases | ❌       |                                                                                                                                                                                                                                                                                                 |
| AfterCursor | string                | ❌       | The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination. |

# Purchases

**Properties**

| Name         | Type                    | Required | Description                                                                                                                                                                                                    |
| :----------- | :---------------------- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Id           | string                  | ❌       | ID of the purchase                                                                                                                                                                                             |
| StartDate    | string                  | ❌       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                     |
| EndDate      | string                  | ❌       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                       |
| Duration     | float64                 | ❌       | It designates the number of days the eSIM is valid for within 90-day validity from issuance date.                                                                                                              |
| CreatedDate  | string                  | ❌       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                            |
| StartTime    | float64                 | ❌       | Epoch value representing the start time of the package's validity                                                                                                                                              |
| EndTime      | float64                 | ❌       | Epoch value representing the end time of the package's validity                                                                                                                                                |
| CreatedAt    | float64                 | ❌       | Epoch value representing the date of creation of the purchase                                                                                                                                                  |
| Package\_    | purchases.Package\_     | ❌       |                                                                                                                                                                                                                |
| Esim         | purchases.PurchasesEsim | ❌       |                                                                                                                                                                                                                |
| Source       | string                  | ❌       | The `source` indicates whether the purchase was made from the API, dashboard, landing-page, promo-page or iframe. For purchases made before September 8, 2023, the value will be displayed as 'Not available'. |
| PurchaseType | string                  | ❌       | The `purchaseType` indicates whether this is the initial purchase that creates the eSIM (First Purchase) or a subsequent top-up on an existing eSIM (Top-up Purchase).                                         |
| ReferenceId  | string                  | ❌       | The `referenceId` that was provided by the partner during the purchase or top-up flow. This identifier can be used for analytics and debugging purposes.                                                       |

# Package\_

**Properties**

| Name             | Type    | Required | Description                                                                                                                                                                 |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Id               | string  | ❌       | ID of the package                                                                                                                                                           |
| DataLimitInBytes | float64 | ❌       | Size of the package in bytes. For **limited packages**, this field will return the data limit in bytes. For **unlimited packages**, it will return **-1** as an identifier. |
| Destination      | string  | ❌       | ISO representation of the package's destination.                                                                                                                            |
| DestinationName  | string  | ❌       | Name of the package's destination                                                                                                                                           |
| PriceInCents     | float64 | ❌       | Price of the package in cents                                                                                                                                               |

# PurchasesEsim

**Properties**

| Name  | Type   | Required | Description    |
| :---- | :----- | :------- | :------------- |
| Iccid | string | ❌       | ID of the eSIM |
