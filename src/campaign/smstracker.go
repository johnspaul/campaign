package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
	"time"
	"net/http"
)

// SmstrackerController implements the smstracker resource.
type SmstrackerController struct {
	*goa.Controller
}

// NewSmstrackerController creates a smstracker controller.
func NewSmstrackerController(service *goa.Service) *SmstrackerController {
	return &SmstrackerController{Controller: service.NewController("SmstrackerController")}
}

// Create runs the create action.
func (c *SmstrackerController) Create(ctx *app.CreateSmstrackerContext) error {
	// SmstrackerController_Create: start_implement

	go func() {
		time.Sleep(1000*time.Millisecond)
		http.Post()
	}()

	// SmstrackerController_Create: end_implement
	return nil
}
