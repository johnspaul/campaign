//go:generate goagen bootstrap -d campaign/design

package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	_"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	// Create service
	service := goa.New("campaign")
	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	db, _ := sql.Open("mysql", "root:pass123@/campaign?charset=utf8&parseTime=True&loc=Local")
	// Mount "campaigns" controller
	c := NewCampaignsController(service,db)
	app.MountCampaignsController(service, c)
	// Mount "messagecontents" controller
	c2 := NewMessagecontentsController(service,db)
	app.MountMessagecontentsController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
    c4 := NewProductsController(service)
    app.MountProductsController(service,c4)
    c5 := NewLeadController(service)
    app.MountLeadController(service,c5)
	// Start service

	defer db.Close()
	if err := service.ListenAndServe(":8050"); err != nil {
		service.LogError("startup", "err", err)
	}
}


