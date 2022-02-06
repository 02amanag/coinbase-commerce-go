package coinbase

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

// Create will create new payment session with respect to the information passed
func (c *ChargeService) Create(body interface{}) (interface{}, error) {
	response, err := c.Api.Fetch("POST", "/charges", body)
	if err != nil {
		return nil, err
	}
	return response, nil

}

// Get will give information of any specific payment with respect to there code/Id
func (c *ChargeService) Get(code string) (interface{}, error) {
	response, err := c.Api.Fetch("GET", "/charges/"+code, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Cancel will update the payment status of respective Id to cancel.
func (a *ChargeService) Cancel(id string) (interface{}, error) {
	response, err := a.Api.Fetch("POST", "/charges/"+id+"/cancel", nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// List will response bac with all the payment sessions that are created with there complete information.
func (a *ChargeService) List() (interface{}, error) {
	response, err := a.Api.Fetch("GET", "/charges/", nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}
