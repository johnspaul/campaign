package campaign
import (
. "github.com/goadesign/goa/design"
. "github.com/goadesign/goa/design/apidsl"
)


var leadPoolMedia = MediaType("application/ts.leadpool", func() {
	TypeName("leadPoolLength")

	Attributes(func() {
		Attribute("length",String,"length of lead pool")
	})
	View("default", func() {
		Attribute("length")
	})
})

var _ = Resource("lead", func() {
	Description("Lead module")
	BasePath("/lead")
	Action("get", func() {
		Routing(GET("/:productId"))
		Params(func() {
			Param("productId",String,"the product id")
		})
		Response(OK, func() {
			Description("The length of lead pool")
			Status(201)
			Media(leadPoolMedia,"default")
		})
		Response(NotFound, func() {
			Status(404)
			Description("length not found")
		})
	})
})