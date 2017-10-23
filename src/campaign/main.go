//go:generate goagen bootstrap -d campaign/design

package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("campaign")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "campaigns" controller
	c := NewCampaignsController(service)
	app.MountCampaignsController(service, c)
	// Mount "messagecontents" controller
	c2 := NewMessagecontentsController(service)
	app.MountMessagecontentsController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8050"); err != nil {
		service.LogError("startup", "err", err)
	}

}
