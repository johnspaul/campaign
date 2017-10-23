package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
)

// MessagecontentsController implements the messagecontents resource.
type MessagecontentsController struct {
	*goa.Controller
}

// NewMessagecontentsController creates a messagecontents controller.
func NewMessagecontentsController(service *goa.Service) *MessagecontentsController {
	return &MessagecontentsController{Controller: service.NewController("MessagecontentsController")}
}

// Create runs the create action.
func (c *MessagecontentsController) Create(ctx *app.CreateMessagecontentsContext) error {
	// MessagecontentsController_Create: start_implement

	// Put your logic here

	// MessagecontentsController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *MessagecontentsController) Delete(ctx *app.DeleteMessagecontentsContext) error {
	// MessagecontentsController_Delete: start_implement

	// Put your logic here

	// MessagecontentsController_Delete: end_implement
	return nil
}

// Get runs the get action.
func (c *MessagecontentsController) Get(ctx *app.GetMessagecontentsContext) error {
	// MessagecontentsController_Get: start_implement

	// Put your logic here

	// MessagecontentsController_Get: end_implement
	res := &app.CampaignMessageContent{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *MessagecontentsController) List(ctx *app.ListMessagecontentsContext) error {
	// MessagecontentsController_List: start_implement

	// Put your logic here

	// MessagecontentsController_List: end_implement
	res := &app.CampaignMessageContent{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *MessagecontentsController) Update(ctx *app.UpdateMessagecontentsContext) error {
	// MessagecontentsController_Update: start_implement

	// Put your logic here

	// MessagecontentsController_Update: end_implement
	return nil
}
