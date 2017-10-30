package main

import (
	"campaign/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"database/sql"
	"github.com/satori/go.uuid"
	"fmt"
)

// CampaignsController implements the campaigns resource.
type CampaignsController struct {
	*goa.Controller
	*sql.DB
}

// NewCampaignsController creates a campaigns controller.
func NewCampaignsController(service *goa.Service,db *sql.DB) *CampaignsController {
	return &CampaignsController{Controller: service.NewController("CampaignsController"),DB: db}
}

// Create runs the create action.
func (c *CampaignsController) Create(ctx *app.CreateCampaignsContext) error {
	// CampaignsController_Create: start_implement

	id := uuid.NewV4()
	stmt,err := c.DB.Prepare("insert INTO campaign(campaignId,productId,status,startdate," +
		"enddate,activeStartHour,activeStartMinute,activeHours,pollingInterval,executionFrequency) VALUES (?,?,0,?,?,?,?,?,?,?)")
	res,err := stmt.Exec(id,
		ctx.Payload.ProductID,
		ctx.Payload.StartDate,
		ctx.Payload.EndDate,
		ctx.Payload.ActiveStartHour,
		ctx.Payload.ActiveStartMinute,
		ctx.Payload.ActiveHours,
		ctx.Payload.PollingInterval,
		ctx.Payload.ExecutionFrequency)
	checkErr(err)
	for  _,message := range ctx.Payload.Messages {
		campaignMessageId := uuid.NewV4()
		stmt1, _ := c.DB.Prepare("INSERT INTO campaign_message(campaignMessageId,campaignId,messageId,percentage) VALUES(?,?,?,?)")
		_,err = stmt1.Exec(campaignMessageId,id,message.MessageID,message.Percentage)
		checkErr(err)
	}
	fmt.Println(res.RowsAffected())
	if ctx.Payload.StartDate > ctx.Payload.EndDate {
		ctx.ResponseData.WriteHeader(500)
		return nil
	}
	// CampaignsController_Create: end_implement
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
// Delete runs the delete action.
func (c *CampaignsController) Delete(ctx *app.DeleteCampaignsContext) error {
	// CampaignsController_Delete: start_implement
//
//stmt,_ := c.DB.Prepare("DELETE from campaign where campaignId =?")
//row, _ :=
	// CampaignsController_Delete: end_implement
	return nil
}

// Get runs the get action.
func (c *CampaignsController) Get(ctx *app.GetCampaignsContext) error {
	// CampaignsController_Get: start_implement

	stmt,_ := c.DB.Prepare("select campaignId,productId,status,startdate," +
		"enddate,activeStartHour,activeStartMinute,activeHours,pollingInterval,executionFrequency from campaign where campaignId =?")
	row,_ := stmt.Query(ctx.CampaignID)
	//checkErr(err)
	row.Next()
	var campaign app.CampaignDetailed
	_ = row.Scan(&campaign.CampaignID,
		&campaign.ProductID,
		&campaign.Status,
		&campaign.StartDate,
		&campaign.EndDate,
		&campaign.ActiveStartHour,
		&campaign.ActiveStartMinute,
		&campaign.ActiveHours,
		&campaign.PollingInterval,
		&campaign.ExecutionFrequency)
		fmt.Println(campaign)
	// CampaignsController_Get: end_implement
	//res := &app.CampaignDetailed{}
	return ctx.OKDetailed(&campaign)
}

// GetAll runs the getAll action.
func (c *CampaignsController) GetAll(ctx *app.GetAllCampaignsContext) error {
	// CampaignsController_GetAll: start_implement


	stmt,_ := c.DB.Prepare("select campaignId,productId,status,startdate," +
		"enddate,activeStartHour,activeStartMinute,activeHours,pollingInterval,executionFrequency from campaign")
	row,err := stmt.Query()
	checkErr(err)
	var campaigns app.CampaignCollection
	var campaign app.Campaign
	 for row.Next() {
		 row.Scan(&campaign.CampaignID,
		 &campaign.ProductID,
			 &campaign.Status,
			 &campaign.StartDate,
			 &campaign.EndDate,
			 &campaign.ActiveStartHour,
			 &campaign.ActiveStartMinute,
			 &campaign.ActiveHours,
			 &campaign.PollingInterval,
			 &campaign.ExecutionFrequency)
		 campaigns = append(campaigns,&campaign)
	 }
		fmt.Println(campaign)

	// CampaignsController_GetAll: end_implement
	//res := app.CampaignCollection{}
	return ctx.OK(campaigns)
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
