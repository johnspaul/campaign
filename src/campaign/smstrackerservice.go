package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
	"time"
	"net/http"
	"bytes"
	"strconv"
)

// SmstrackerserviceController implements the smstrackerservice resource.
type SmstrackerserviceController struct {
	*goa.Controller
}

// NewSmstrackerserviceController creates a smstrackerservice controller.
func NewSmstrackerserviceController(service *goa.Service) *SmstrackerserviceController {
	return &SmstrackerserviceController{Controller: service.NewController("SmstrackerserviceController")}
}

// Create runs the create action.
func (c *SmstrackerserviceController) Create(ctx *app.CreateSmstrackerserviceContext) error {
	// SmstrackerserviceController_Create: start_implement

	go func() {
		time.Sleep(1000*time.Millisecond)
		http.Post(*ctx.Payload.CallbackApiforSend,"application/json",bytes.NewBufferString(`{phoneNumber:`+strconv.Itoa(*ctx.Payload.PhoneNumber)+`}`))
	}()
	// SmstrackerserviceController_Create: end_implement
	res := &app.SmsMedia{}
	return ctx.OK(res)
}
