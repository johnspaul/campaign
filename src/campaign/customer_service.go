package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
)

// CustomerServiceController implements the customerService resource.
type CustomerServiceController struct {
	*goa.Controller
}

// NewCustomerServiceController creates a customerService controller.
func NewCustomerServiceController(service *goa.Service) *CustomerServiceController {
	return &CustomerServiceController{Controller: service.NewController("CustomerServiceController")}
}

// Create runs the create action.
func (c *CustomerServiceController) Create(ctx *app.CreateCustomerServiceContext) error {
	// CustomerServiceController_Create: start_implement

	// Put your logic here

	// CustomerServiceController_Create: end_implement
	return nil
}
