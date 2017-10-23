package design
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("campaign", func() {

	Title("The Campaign Management API")
	Description("Trusting social campaign management")
	Host("localhost:8050")
	Scheme("http")
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
var campaignMessagePayload = Type("CampaignMessagePayload", func() {

	Description("Message  attached to a campaign")
	Attribute("campaignId", String, "The campaign id", func() {
		MinLength(1)
		MaxLength(36)
	})

	Attribute("messageId", String, "The message id", func() {
		MinLength(1)
		MaxLength(160)
	})

	Attribute("percentage",Number, "The percentage pf this message to be used", func() {

	})

	Required ("campaignId")
	Required ("messageId")
	Required ("percentage")

})

var campaignMessageMedia = MediaType("application/ts.campaign.message", func() {

	TypeName("CampaignMessage")

	Attributes(func() {
		Attribute("campaignId", String, "campaign id") // Operation results attribute
		Attribute("message", messagePayload, "The message") // Operation results attribute
		Attribute("percentage", String, "Percentage of sms to be made out of this message") // Operation results attribute
	})

	View("default", func(){
		Attribute("campaignId")
		Attribute("message")
		Attribute("percentage")

	})


})


var campaignPayload = Type("CampaignPayload", func() {

	Description("The Campaign object")

	Attribute("productId", String, "Product Id for which the campaign is created", func() {
		MinLength(1)
		MaxLength(36)
	})
	Attribute("state", Integer, "State of the campaign", func() {

	})

	Attribute("startDate", DateTime, "Start time of the campaign")

	Attribute("endDate", DateTime, "End date of the campaign")

	Attribute("pollingInterval", Number, "Interval in which the campaign need to poll the lead queue")

	Attribute("messages", ArrayOf(campaignMessagePayload),"Message content to be attached")

	Required ("productId")
	Required ("startDate")
	Required ("messages")
})


var campaignUpdatePayload = Type("CampaignUpdatePayload", func() {

	Description("The Campaign object")

	Attribute("campaignId", String, "campaign id")
	Attribute("state", Integer, "State of the campaign", func() {

	})

	Attribute("startDate", DateTime, "Start time of the campaign")

	Attribute("endDate", DateTime, "End date of the campaign")

	Attribute("pollingInterval", Number, "Interval in which the campaign need to poll the lead queue")

	Attribute("messages", ArrayOf("CampaignMessagePayload"),"Message content to be attached")

})



var campaignMedia = MediaType("application/ts.campaign", func() {

	TypeName("Campaign")
	Reference(campaignPayload)
	Attributes(func() {
		Attribute("campaignId", String, "Campaign id") // Operation results attribute
		Attribute("state", Integer, "State of the Campaign") // Operation results attribute
		Attribute("startDate", Integer, "Start date of the Campaign") // Operation results attribute
		Attribute("endDate", Integer, "End date of the Campaign") // Operation results attribute
		Attribute("pollingInterval", Number, "Interval in which the campaign need to poll the lead queue")
		Attribute("messages", ArrayOf(campaignMessageMedia))
	})

	View("default", func(){
		Attribute("campaignId")
		Attribute("state")
		Attribute("startDate")
		Attribute("endDate")
		Attribute("pollingInterval")
	})

	View("detailed"  , func(){
		Attribute("pollingInterval")
		Attribute("campaignId")
		Attribute("state")
		Attribute("startDate")
		Attribute("endDate")
		Attribute("messages", ArrayOf(campaignMessageMedia))
	})

})

var campaignExecutionMedia = MediaType("application/ts.campaignexecution", func() {

	TypeName("CampaignExecutionContext")

	Attributes(func() {
		Attribute("campaignId", String, "campaign id")
		Attribute("executionId", String, "execution id")
		Attribute("startTime", DateTime, "execution start time")
		Attribute("endTime", DateTime, "execution endTime time")
		Attribute("numMessagesSent", Number, "Number of message sent in this execution")
	})

	View("default", func(){
		Attribute("campaignId")
		Attribute("executionId")
		Attribute("startTime")
		Attribute("endTime")
		Attribute("numMessagesSent")
	})

})


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
			Param("messageId", String, "The message content id")

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
			Media(messageMediaType, "default")
		})
	})


	Action("update", func() {

		Routing(PUT("/:messageId"))
		Description("Updates a  message.")
		Payload(messageUpdatePayload)

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
			Description("This is the success response.")
			Status(204)
		})
		Response(NotFound, func() {
			Description("Item not found")
			Status(404)
		})
	})
})


var _ = Resource("campaigns", func() {
	Description("Represents Campaign instance for a Product.")
	BasePath("/campaigns")

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
		Description("Creates a campaign.")
		Payload(campaignPayload)
		Response("Created", func() {
			Description("Organization created successfully.")
			Status(201)
			Media(campaignMedia, "default")
		})
	})
	Action("getAll", func() {
		Routing(GET("/"))
		Description("Returns all the campaigns")
		Params(func() {
			Param("state", Integer, "The state", func() {
			})
		})
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(campaignMedia, "default")
		})

	})

	Action("get", func() {
		Routing(GET("/:campaignId"))
		Description("Returns a campaign with  specific id.")
		Params(func() {
			Param("campaignId", String, "The campaignId")

		})
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(campaignMedia, "detailed")
		})
		Response(NotFound, func() {
			Description("Response when Not Found")
			Status(404)
		})
	})

	Action("update", func() {

		Routing(PATCH("/:campaignId"))
		Description("Updates a campaign.")
		Params(func() {
			Param("campaignId",String,"The id of the campaign to be updated")
		})
		Payload(campaignUpdatePayload)
		Response("Updated", func() {
			Description("This is the success response.")
			Status(202)
			Media(campaignMedia, "default")
		})
		Response(NotFound, func() {
			Description("Item not found")
			Status(404)
		})
	})

	Action("delete", func() {

		Routing(DELETE("/:campaignId"))
		Description("Deletes a terminated campaign.")
		Params(func() {
			Param("campaignId",String,"The id of the campaign to be updated")
		})
		Response(OK, func() {
			Description("This is the success response for delete of a campaign.")
			Status(204)
		})
		Response(NotFound, func() {
			Description("Item not found")
			Status(404)
		})
	})
	Action("getAllCampaignExecution", func() {
		Routing(GET("/:campaignId/executions"))
		Description("Returns all campaign  execution details.")
		Params(func() {
			Param("campaignId", String, "The campaignId")

		})
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(campaignExecutionMedia, "default")
		})
		Response(NotFound, func() {
			Description("Response for Not Found")
			Status(404)
		})
	})

	Action("getCampaignExecution", func() {
		Routing(GET("/:campaignId/executions/:executionId"))
		Description("Returns a specific campaign execution details")

		Params(func() {
			Param("executionId", String, "The execution Id")
			Param("campaignId",String,"The campaign Id")

		})
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(campaignExecutionMedia, "default")
		})
		Response(NotFound, func() {
			Description("Response for Not Found")
			Status(404)
		})
	})
})




var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
})