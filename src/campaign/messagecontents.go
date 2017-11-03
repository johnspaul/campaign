package main

import (
	"campaign/app"
	"github.com/goadesign/goa"
	"database/sql"
	"github.com/satori/go.uuid"
	"fmt"
)

// MessagecontentsController implements the messagecontents resource.
type MessagecontentsController struct {
	*goa.Controller
	*sql.DB

}

// NewMessagecontentsController creates a messagecontents controller.
func NewMessagecontentsController(service *goa.Service,db *sql.DB) *MessagecontentsController {
	return &MessagecontentsController{Controller: service.NewController("MessagecontentsController"),DB: db}
}

// Create runs the create action.
func (c *MessagecontentsController) Create(ctx *app.CreateMessagecontentsContext) error {
	// MessagecontentsController_Create: start_implement
	id := uuid.NewV4()
	stmt,err := c.DB.Prepare("insert INTO message_content(messageId,content) VALUES (?,?)")
	fmt.Println(ctx.Payload.MessageContent)
	_,err = stmt.Query(id,ctx.Payload.MessageContent)
	checkErr(err)

	//fmt.Println(res.RowsAffected())

	// MessagecontentsController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *MessagecontentsController) Delete(ctx *app.DeleteMessagecontentsContext) error {
	// MessagecontentsController_Delete: start_implement

	// Put your logic here
	stmt,err := c.DB.Prepare("DELETE from message_content where messageId = ?")
	fmt.Println(ctx.MessageID)
	_,err = stmt.Query(ctx.MessageID)
	checkErr(err)

	// MessagecontentsController_Delete: end_implement
	return nil
}

// Get runs the get action.
func (c *MessagecontentsController) Get(ctx *app.GetMessagecontentsContext) error {
	// MessagecontentsController_Get: start_implement

	stmt,_ := c.DB.Prepare("select messageId,content from message_content where messageId =?")
	fmt.Println(ctx.MessageID)
	row,err := stmt.Query(ctx.MessageID)
	if err!= nil {
		fmt.Println(err)
	}
	row.Next()
	var CampaignMessageContent app.CampaignMessageContent
	err=row.Scan(&CampaignMessageContent.MessageID,
		&CampaignMessageContent.Content)

	fmt.Println(CampaignMessageContent)
	// MessagecontentsController_Get: end_implement
	//res := &app.CampaignMessageContent{}
	//return ctx.OK(res)
	return ctx.OK(&CampaignMessageContent)
}

// List runs the list action.
func (c *MessagecontentsController) List(ctx *app.ListMessagecontentsContext) error {
	// MessagecontentsController_List: start_implement
	stmt,_ := c.DB.Prepare("select messageId,content from message_content ")

	row,err := stmt.Query()
	if err!= nil {
		fmt.Println(err)
	}

	var messages app.CampaignMessageContentCollection
	var CampaignMessageContent app.CampaignMessageContent


	for row.Next() {
		row.Scan(&CampaignMessageContent.MessageID,
			&CampaignMessageContent.Content)
			message := CampaignMessageContent
		messages = append(messages,&message)
	}

	fmt.Println(messages)

	return ctx.OK(messages)
	//// MessagecontentsController_List: end_implement
	//res := &app.CampaignMessageContent{}
	//return ctx.OK(res)
}

// Update runs the update action.
func (c *MessagecontentsController) Update(ctx *app.UpdateMessagecontentsContext) error {
	// MessagecontentsController_Update: start_implement

	stmt,_ := c.DB.Prepare("UPDATE message_content SET content=? WHERE messageId = ?")

	_,err := stmt.Query(ctx.Payload.MessageContent,ctx.MessageID)
	if err!= nil {
		fmt.Println(err)
	}
	checkErr(err)

	// MessagecontentsController_Update: end_implement
	return nil
}
