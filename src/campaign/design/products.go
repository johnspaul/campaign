package campaign
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)
var productMedia = MediaType("application/ts.campaign.product", func() {

	TypeName("campaignProductMedia")

	Attribute("productId",String)
	Attribute("productType",String)
	Attribute("productCode",String)
	Attribute("clientCode",String)
	Attribute("dailyVolume",Integer)
	Attribute("availableLocations",ArrayOf(productLocation))
	Attribute("criteria",ArrayOf(productCriteria))

	View("default", func() {
		Attribute("productId")
		Attribute("productType")
		Attribute("productCode")
		Attribute("clientCode")
		Attribute("dailyVolume")
		Attribute("availableLocations")
		Attribute("criteria")
	})
})

var productPayload = Type("Product Payload", func() {
	Description("")

	Attribute("productType", String)
	Attribute("productCode",String)
	Attribute("clientCode",String)
	Attribute("dailyVolume",Integer)
	Attribute("availableLocations",ArrayOf(productLocation))
	Attribute("criteria",ArrayOf(productCriteria))

})
var productLocation = Type("product location", func() {
	Description("")

	Attribute("province",String)
	Attribute("district",String)

})

var productCriteria = Type("product criteria", func() {
	Description("")
	Attribute("variable",String)
	Attribute("comparisonType",String)
	Attribute("minValue",Integer)
	Attribute("maxValue",Integer)

})
var _= Resource("products", func() {
	Description("The details of a product")
	BasePath("/products")

	Action("create", func() {
		Routing(POST("/"))
		Description("creates a product")
		Payload(productPayload)
		Response("created", func() {
			Description("Product Created Successfully")
			Status(201)
		})
	})

	Action("get", func() {
		Routing(GET("/:productId"))
		Description("")
		Params(func() {
			Param("productId", String, "the product id")
		})
		Response(OK, func() {
			Media(productMedia,"default")
			Status(201)
			Description("get product details by id")
		})
		Response(NotFound, func() {
			Status(404)
			Description("product detail not found")
		})
	})
})
