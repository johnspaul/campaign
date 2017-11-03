package campaign
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("messagecontents", func() {
	Description("Represents  a message content to be attached to  1 or more campaigns.")
	BasePath("/messagecontents")

	Response("Unauthorized", func() {
		Description("Response sent for unauthorized requests.")
		Status(401)
	})
	Response("Bad Request", func() {
		Description("Response sent for bad requests.")
		Media("application/text")
		Status(400)
	})
	Response("Internal Server Error", func() {
		Description("Response sent for Database or Internal Server Errors.")
		Media("application/text")
		Status(500)
	})

	Action("create", func() {

		Routing(POST("/"))
		Description("Creates a message content.")
		Payload(messagePayload)
		Response("Created", func() {
			Description("Message content  created successfully.")
			Status(201)
			Media(messageMediaType, "default")
		})
	})

	Action("get", func() {
		Routing(GET("/:messageId"))
		Description("Returns the messages defined for specific product type.")
		Params(func() {
			Param("messageId", String, "The message contentid")

		})
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(messageMediaType, "default")
		})
		Response(NotFound, func() {
			Description("Response for Not Found")
			Status(404)
		})
	})


	Action("list", func() {
		Routing(GET("/"))
		Description("Returns all the messages.")
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(CollectionOf(messageMediaType), "default")
		})
	})

	Action("update", func() {
		Routing(PUT("/:messageId"))
		Description("Updates a  message.")
		Payload(messageUpdatePayload)
		Params(func() {
			Param("messageId", String, "The message content id")

		})
		Response("Updated", func() {
			Description("This is the success response.")
			Status(202)
			Media(messageMediaType, "default")
		})
		Response(NotFound, func() {
			Description("Item not found")
			Status(404)
		})
	})

	Action("delete", func() {

		Routing(DELETE("/:messageId"))
		Description("Deletes a message.")
		Params(func() {
			Param("messageId", String, "The message content id")

		})
		Response(OK, func() {
			Description("Deleted successfully")
			Status(204)
		})
		Response(NotFound, func() {
			Description("Item not found")
			Status(404)
		})
	})
})
var messagePayload = Type("MessageContentPayload", func() {

	Description("The Message content object")
	Attribute("messageContent", String, "The message content", func() {
		MinLength(1)
		MaxLength(160)
	})
	Required ("messageContent")
})

var messageUpdatePayload = Type("MessageContentUpdatePayload", func() {

	Description("The Message content object")
	Attribute("messageContent", String, "The message content", func() {
		MinLength(1)
		MaxLength(160)
	})

	Required ("messageContent")

})



var messageMediaType = MediaType("application/ts.campaign.messagecontent", func() {

	TypeName("CampaignMessageContent")

	Attributes(func() {
		Attribute("messageId", String, "Message content  id") // Operation results attribute
		Attribute("content", String, "The Message") // Operation results attribute
	})

	View("default", func(){
		Attribute("messageId")
		Attribute("content")

	})

})

