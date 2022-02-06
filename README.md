# coinbase-commerce-go
Coinbase Commerce Golang

# Table of contents

<!--ts-->
   * [Documentation](#documentation)
   * [Installation](#installation)
   * [Usage](#usage)
      * [Charges](#charges)
<!--te-->


## Documentation

For more details visit [Coinbase API docs](https://commerce.coinbase.com/docs/api/)
To start using library, you need to register on [Commerce SignUp](https://commerce.coinbase.com/signup).
And get your ``API_KEY`` from user settings.

Next create a ``APIClient`` object for interacting with the API:
```golang
import "github.com/02amanag/coinbase-commerce-go"

client := coinbase.Client(API_KEY)
```

``Client`` contains link to an Golang Class representing ``Charge``

You can call ``Create, List, Get, Cancel`` methods from an API resource classes

```golang
client.Charge.Create
client.Charge.Cancel
client.Charge.List
client.Charge.Get
```

Client method returns an ``API resource instance`` (``APICharge``) representing the response from the API, all of the models are dumpable with JSON.\
The response data is parsed into Golang interface.

Client support Common Errors and Warnings handling.
All errors occuring during interaction with the API will be return.


| Error                    | Status Code |
|--------------------------|-------------|
| APIError                 |      *      |   
| InvalidRequestError      |     400     |   
| ParamRequiredError       |     400     |  
| ValidationError          |     400     |  
| AuthenticationError      |     401     |  
| ResourceNotFoundError    |     404     |
| RateLimitExceededError   |     429     |
| InternalServerError      |     500     |
| ServiceUnavailableError  |     503     |

## Installation

Install with ``go get``:

    go get "github.com/02amanag/coinbase-commerce-go"


## Usage
```golang
import "github.com/02amanag/coinbase-commerce-go"

client := coinbase.Client(API_KEY)
```

## Charges
[Charges API docs](https://commerce.coinbase.com/docs/api/#charges)
### Get
```golang
charge ,err := client.Charge.Get(<charge_id>)
```
### Create
```golang
#by struct
charge, err := client.Charge.Create(coinbase.ChargeParam{
		Name:        "Test Name",
		Description: "test description for testing",
		Local_price: coinbase.Money{
			Amount:   "101.00", //amount to be paid
			Currency: "USD",
		},
		Pricing_type: "fixed_price",
		Metadata: coinbase.Metadata{ //extra information
			Customer_id:   "ID001",
			Customer_name: "Test Customer Name",
		},
		Redirect_url: "https://google.com",   //success page
		Cancel_url:   "https://facebook.com", //cancel page
	})

#or directly by json
charge_info := `{
	        "name": "Test Name 2",
	        "description": "testing description",
	        "local_price": {
	          "amount": "101.00",
	          "currency": "USD"
	        },
	        "pricing_type": "fixed_price",
	        "metadata": {
	          "customer_id": "id_1005",
	          "customer_name": "test customer Name "
	        },
	        "redirect_url": "https://google.com",
	        "cancel_url": "https://facebook.com"
	      }`
charge, err := client.Charge.Create(charge_info)
```
### List
```golang
charges, err := client.Charge.List()
```
### Cancel
```golang
charges, err := client.Charge.Cancel(<charge_id>)
```

## Types
### Charge
#### Charge Structures
```golang

type ChargeService struct {
	Api *APIClient
}

type ChargeParam struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Local_price  Money    `json:"local_price"`
	Pricing_type string   `json:"pricing_type"`
	Metadata     Metadata `json:"metadata"`
	Redirect_url string   `json:"redirect_url"`
	Cancel_url   string   `json:"cancel_url"`
}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Metadata struct {
	Customer_id   string `json:"customer_id"`
	Customer_name string `json:"customer_name"`
}
```

# Example
## You can find complete implementation in Example folder and execute it directly through
``` golang
go run example/main.go
```