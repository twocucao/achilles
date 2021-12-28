package response

import "achilles/models/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
