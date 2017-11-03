
package campaign
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


var campaignMessagePayload = Type("CampaignMessagePayload", func() {

	Description("Message  attached to a campaign")
	Attribute("messageId", String, "The message id", func() {
		MinLength(1)
		MaxLength(160)
	})
	Attribute("percentage",Number, "The percentage pf this message to be used", func() {

	})
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
var customerServicePayload = Type("customerServicePayload", func() {
Attribute("phoneNumber")
})

var campaignPayload = Type("CampaignPayload", func() {
	Description("The Campaign object")

	Attribute("productId", String, "Product Id for which the campaign is created", func() {
		MinLength(1)
		MaxLength(36)
	})
	Attribute("startDate",Integer,"Start date of the Campaign") // Operation results attribute
	Attribute("endDate",Integer, "End date of the Campaign") // Operation results attribute

	Attribute("activeStartHour", Integer, "The active start hour -  campaign execution will be starting from this time", func() {
		Maximum(23)
		Minimum(0)
	})
	Attribute("activeStartMinute", Integer, "The active start minutes", func() {
		Maximum(59)
		Minimum(0)
	})
	Attribute("activeHours", Integer, "Duration in which campaign will be running starting from activeStartTime -in minutes")
	Attribute("executionFrequency", Integer,"Frequency in which the campaign needs to be executed - in days", func() {
		Default(1)
		Minimum(1)
	})
	Attribute("pollingInterval", Integer, "Interval in which the campaign need to poll the lead queue - in minutes")

	Attribute("messages", ArrayOf(campaignMessagePayload),"Message content to be attached")

	Required ("productId")
	Required ("startDate")
	Required ("endDate")
	Required ("messages")
})

var campaignUpdatePayload = Type("CampaignUpdatePayload", func() {

	Description("The Campaign object")

	Attribute("status", Integer, "State of the campaign")
	Attribute("startDate", Integer, "Start time of the campaign")
	Attribute("endDate", Integer, "End date of the campaign")

	Attribute("activeStartHour", Integer, "The active start hour -  campaign execution will be starting from this time")
	Attribute("activeStartMinute", Integer, "The active start minutes")
	Attribute("activeHours", Integer, "Duration in which campaign will be running starting from activeStartTime -in minutes")

	Attribute("executionFrequency", Integer,"Interval in which the campaign needs to be executed - in days")
	Attribute("pollingInterval", Integer, "Interval in which the campaign need to poll the lead queue")


	Attribute("messages", ArrayOf(campaignMessagePayload),"Message content to be attached")

})

var campaignMedia = MediaType("application/ts.campaign", func() {

	TypeName("Campaign")
	Reference(campaignPayload)
	Attributes(func() {

		Attribute("productId", String, "Product Id")
		Attribute("campaignId", String, "Campaign id") // Operation results attribute
		Attribute("status", Integer, "State of the Campaign") // Operation results attribute

		Attribute("approvedOn", Integer,"Approved On")
		Attribute("approvedBy", Integer,"Approved by")

		Attribute("startDate", Integer, "Start date of the Campaign") // Operation results attribute
		Attribute("endDate", Integer, "End date of the Campaign") // Operation results attribute

		Attribute("activeStartHour", Integer, "The active start hour -  campaign execution will be starting from this time")
		Attribute("activeStartMinute", Integer, "The active start minute")
		Attribute("activeHours", Integer, "Duration in which campaign will be running starting from activeStartTime -in minutes)")
		Attribute("executionFrequency", Integer,"Frequency in which the campaign needs to be executed - in days")
		Attribute("pollingInterval", Integer, "Interval in which the campaign need to poll the lead queue - in minutes")
		Attribute("lastExecutionTime", Integer, "The time at which the Campaign executed for the last time.") // Operation results attribute

		Attribute("currentExecutionCycleStartTime", Integer, "Start date of the current execution cycle") // Operation results attribute
		Attribute("currentTargetVolume" , Integer , "Lead volume that needs to be achieved with in current cycle") // Operation results attribute

		Attribute("messages", ArrayOf(campaignMessageMedia))
	})

	View("default", func(){

		Attribute("productId")
		Attribute("campaignId")
		Attribute("status")
		Attribute("approvedOn")
		Attribute("approvedBy")
		Attribute("startDate")
		Attribute("endDate")
		Attribute("activeStartHour")
		Attribute("activeStartMinute")
		Attribute("activeHours")
		Attribute("lastExecutionTime")
		Attribute("pollingInterval")
		Attribute("executionFrequency")
		Attribute("currentExecutionCycleStartTime")
		Attribute("currentTargetVolume")
	})

	View("detailed"  , func(){

		Attribute("productId")
		Attribute("campaignId")
		Attribute("status")
		Attribute("approvedOn")
		Attribute("approvedBy")
		Attribute("startDate")
		Attribute("endDate")
		Attribute("activeStartHour")
		Attribute("activeStartMinute")
		Attribute("activeHours")
		Attribute("lastExecutionTime")
		Attribute("pollingInterval")
		Attribute("executionFrequency")
		Attribute("currentExecutionCycleStartTime")
		Attribute("currentTargetVolume")
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

var smsPayload = Type("smsPayload", func() {
	Attribute("campaignId",String," the campaign id")
	Attribute("phoneNumber",Integer,"phone number to sent sms")
	Attribute("messageContent",String,"The content of sms")
	Attribute("callbackApiforSend",String,"The url of callback api for send")
	Attribute("callbackApiforResponse",String,"The url of callback api for response")
})
var smsMedia = MediaType("application/ts.smstracker", func() {
	TypeName("smsMedia")
	Attributes(func() {
		Attribute("smsId",String,"The id of sms created")
		Attribute("campaignId",String," the campaign id")
		Attribute("phoneNumber",Integer,"phone number to sent sms")
		Attribute("messageContent",String,"The content of sms")
		Attribute("callbackApiforSend",String,"The url of callback api for send")
		Attribute("callbackApiforResponse",String,"The url of callback api for response")
	})
	View("default", func() {
		Attribute("smsId",String,"The id of sms created")
		Attribute("campaignId",String," the campaign id")
		Attribute("phoneNumber",Integer,"phone number to sent sms")
		Attribute("messageContent",String,"The content of sms")
	})
})
var _ = Resource("customerService", func() {
	Description("Mock api for notifying the customer service when a customer replies for an sms")
	BasePath("/customerservcie")
	Action("create", func() {
		Routing(POST("/"))
		Description("informs the customer service when a customer replies for an sms")
	})
})
var _ = Resource("smstrackerservice", func() {
	Description("mock api for sms tracker.It creates an sms and registers callback ")
	BasePath("/smstrackerservice")
	Action("create", func() {
		Routing(POST("/"))
		Description("To send an sms")
		Payload(smsPayload)
		Response(OK, func() {
			Description("sms created ")
			Status(200)
			Media(smsMedia)
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
			Description("campaign created successfully.")
			Status(201)
			Media(campaignMedia, "default")
		})
	})
	Action("getAll", func() {
		Routing(GET("/"))
		Description("Returns all the campaigns")
		Params(func() {
			Param("state", Number, "The state", func() {
			})
		})
		Response(OK, func() {
			Description("This is the success response.")
			Status(200)
			Media(CollectionOf(campaignMedia), "default")
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
		Response("Deleted", func() {
			Description("This is the success response.")
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


