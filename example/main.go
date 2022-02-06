package main

import (
	"encoding/json"
	"fmt"

	coinbase "github.com/02amanag/coinbase-commerce-go"
)

func main() {
	_ = `{
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

	body1 := coinbase.ChargeParam{
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
	}

	client := coinbase.Client("ac3e823b-554c-490b-a557-c0ad730b50a7")
	res, _ := client.Charge.Api.Charge.Create(body1)
	data, _ := json.Marshal(res)
	fmt.Println("json output==> ", string(data))
}
