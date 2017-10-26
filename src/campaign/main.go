//go:generate goagen bootstrap -d campaign/design

package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	_"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/DavidHuie/gomigrate"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create service
	service := goa.New("campaign")
	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	db, err := gorm.Open("mysql", "root:root@/campaign?charset=utf8&parseTime=True&loc=Local")

	// Mount "campaigns" controller
	c := NewCampaignsController(service)
	app.MountCampaignsController(service, c)
	// Mount "messagecontents" controller
	c2 := NewMessagecontentsController(service)
	app.MountMessagecontentsController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
    c4 := NewProductsController(service)
    app.MountProductsController(service,c4)
    c5 := NewLeadController(service)
    app.MountLeadController(service,c5)
	// Start service
	if err != nil {
		log.Fatal(err)
	}
	runMigrations(db)
	defer db.Close()
	if err := service.ListenAndServe(":8050"); err != nil {
		service.LogError("startup", "err", err)
	}
}
func runMigrations(db *gorm.DB){
	//dbMap.Exec("set search_path TO usermanagement")
	dbMap := db.DB()
	migrator, err := gomigrate.NewMigrator(dbMap, gomigrate.Mysql{}, "migrations")
    if err!=nil {
    	fmt.Println(err)
	}

	err = migrator.Migrate()
	if err!=nil {
		fmt.Println(err)
	}
}

