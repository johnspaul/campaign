package main

import (
	"campaign/app"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"math/rand"
	_ "math/rand"
	"strconv"
)

// LeadController implements the lead resource.
type LeadController struct {
	*goa.Controller
}

// NewLeadController creates a lead controller.
func NewLeadController(service *goa.Service) *LeadController {
	return &LeadController{Controller: service.NewController("LeadController")}
}

// Get runs the get action.
func (c *LeadController) Get(ctx *app.GetLeadContext) error {
	// LeadController_Get: start_implement

	var p *ProductPayload
	p.init()
	result, _ := json.Marshal(productMap[ctx.ProductID].DailyVolume)
	resval, _ := strconv.Atoi(string(result))
	length := rand.Intn(resval)
	fmt.Fprintf(ctx.ResponseWriter, strconv.Itoa(length))

	// LeadController_Get: end_implement
	return nil
}
