# CreatePurchaseRequest

**Properties**

| Name          | Type    | Required | Description                                                                                                                                                                                                                  |
| :------------ | :------ | :------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Destination   | string  | ✅       | ISO representation of the package's destination.                                                                                                                                                                             |
| DataLimitInGb | float64 | ✅       | Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20GB                                                                                                                                                |
| StartDate     | string  | ✅       | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.                                                                               |
| EndDate       | string  | ✅       | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                                                                                             |
| Email         | string  | ❌       | Email address where the purchase confirmation email will be sent (including QR Code & activation steps)                                                                                                                      |
| ReferenceId   | string  | ❌       | An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.                                                                                            |
| NetworkBrand  | string  | ❌       | Customize the network brand of the issued eSIM. The `networkBrand` parameter cannot exceed 15 characters in length and must contain only letters and numbers. This feature is available to platforms with Diamond tier only. |
| EmailBrand    | string  | ❌       | Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.       |
| StartTime     | float64 | ❌       | Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.                                                                      |
| EndTime       | float64 | ❌       | Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.                                                                                                           |
