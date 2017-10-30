package main

import (
	"campaign/app"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
)

// ProductsController implements the products resource.
type ProductsController struct {
	*goa.Controller
}

// NewProductsController creates a products controller.
func NewProductsController(service *goa.Service) *ProductsController {
	return &ProductsController{Controller: service.NewController("ProductsController")}
}

// Create runs the create action.
func (c *ProductsController) Create(ctx *app.CreateProductsContext) error {
	// ProductsController_Create: start_implement

	// ProductsController_Create: end_implement
	return nil
}

// Get runs the get action.
func (c *ProductsController) Get(ctx *app.GetProductsContext) error {
	// ProductsController_Get: start_implement
	var p ProductPayload
	p.init()
	res, _ := json.Marshal(productMap[ctx.ProductID])
	fmt.Fprintf(ctx.ResponseWriter, string(res))

	// ProductsController_Get: end_implement
	return nil
}
