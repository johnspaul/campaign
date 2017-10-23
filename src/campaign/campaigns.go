package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
)

// CampaignsController implements the campaigns resource.
type CampaignsController struct {
	*goa.Controller
}

// NewCampaignsController creates a campaigns controller.
func NewCampaignsController(service *goa.Service) *CampaignsController {
	return &CampaignsController{Controller: service.NewController("CampaignsController")}
}

// Create runs the create action.
func (c *CampaignsController) Create(ctx *app.CreateCampaignsContext) error {
	// CampaignsController_Create: start_implement

	// Put your logic here

	// CampaignsController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *CampaignsController) Delete(ctx *app.DeleteCampaignsContext) error {
	// CampaignsController_Delete: start_implement

	// Put your logic here

	// CampaignsController_Delete: end_implement
	return nil
}

// Get runs the get action.
func (c *CampaignsController) Get(ctx *app.GetCampaignsContext) error {
	// CampaignsController_Get: start_implement

	// Put your logic here

	// CampaignsController_Get: end_implement
	res := &app.CampaignDetailed{}
	return ctx.OKDetailed(res)
}

// GetAll runs the getAll action.
func (c *CampaignsController) GetAll(ctx *app.GetAllCampaignsContext) error {
	// CampaignsController_GetAll: start_implement

	// Put your logic here

	// CampaignsController_GetAll: end_implement
	res := &app.Campaign{}
	return ctx.OK(res)
}

// GetAllCampaignExecution runs the getAllCampaignExecution action.
func (c *CampaignsController) GetAllCampaignExecution(ctx *app.GetAllCampaignExecutionCampaignsContext) error {
	// CampaignsController_GetAllCampaignExecution: start_implement

	// Put your logic here

	// CampaignsController_GetAllCampaignExecution: end_implement
	res := &app.CampaignExecutionContext{}
	return ctx.OK(res)
}

// GetCampaignExecution runs the getCampaignExecution action.
func (c *CampaignsController) GetCampaignExecution(ctx *app.GetCampaignExecutionCampaignsContext) error {
	// CampaignsController_GetCampaignExecution: start_implement

	// Put your logic here

	// CampaignsController_GetCampaignExecution: end_implement
	res := &app.CampaignExecutionContext{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *CampaignsController) Update(ctx *app.UpdateCampaignsContext) error {
	// CampaignsController_Update: start_implement

	// Put your logic here

	// CampaignsController_Update: end_implement
	return nil
}
