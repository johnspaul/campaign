// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": Application User Types
//
// Command:
// $ goagen
// --design=campaign/design
// --out=$(GOPATH)/src/campaign
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// Message  attached to a campaign
type campaignMessagePayload struct {
	// The message id
	MessageID *string `form:"messageId,omitempty" json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The percentage pf this message to be used
	Percentage *float64 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// Validate validates the campaignMessagePayload type instance.
func (ut *campaignMessagePayload) Validate() (err error) {
	if ut.MessageID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "messageId"))
	}
	if ut.Percentage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "percentage"))
	}
	if ut.MessageID != nil {
		if utf8.RuneCountInString(*ut.MessageID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.messageId`, *ut.MessageID, utf8.RuneCountInString(*ut.MessageID), 1, true))
		}
	}
	if ut.MessageID != nil {
		if utf8.RuneCountInString(*ut.MessageID) > 160 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.messageId`, *ut.MessageID, utf8.RuneCountInString(*ut.MessageID), 160, false))
		}
	}
	return
}

// Publicize creates CampaignMessagePayload from campaignMessagePayload
func (ut *campaignMessagePayload) Publicize() *CampaignMessagePayload {
	var pub CampaignMessagePayload
	if ut.MessageID != nil {
		pub.MessageID = *ut.MessageID
	}
	if ut.Percentage != nil {
		pub.Percentage = *ut.Percentage
	}
	return &pub
}

// Message  attached to a campaign
type CampaignMessagePayload struct {
	// The message id
	MessageID string `form:"messageId" json:"messageId" xml:"messageId"`
	// The percentage pf this message to be used
	Percentage float64 `form:"percentage" json:"percentage" xml:"percentage"`
}

// Validate validates the CampaignMessagePayload type instance.
func (ut *CampaignMessagePayload) Validate() (err error) {
	if ut.MessageID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "messageId"))
	}

	if utf8.RuneCountInString(ut.MessageID) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.messageId`, ut.MessageID, utf8.RuneCountInString(ut.MessageID), 1, true))
	}
	if utf8.RuneCountInString(ut.MessageID) > 160 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.messageId`, ut.MessageID, utf8.RuneCountInString(ut.MessageID), 160, false))
	}
	return
}

// The Campaign object
type campaignPayload struct {
	// Duration in which campaign will be running starting from activeStartTime -in minutes
	ActiveHours *int `form:"activeHours,omitempty" json:"activeHours,omitempty" xml:"activeHours,omitempty"`
	// The active start hour -  campaign execution will be starting from this time
	ActiveStartHour *int `form:"activeStartHour,omitempty" json:"activeStartHour,omitempty" xml:"activeStartHour,omitempty"`
	// The active start minutes
	ActiveStartMinute *int `form:"activeStartMinute,omitempty" json:"activeStartMinute,omitempty" xml:"activeStartMinute,omitempty"`
	// End date of the Campaign
	EndDate *int `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Frequency in which the campaign needs to be executed - in days
	ExecutionFrequency *int `form:"executionFrequency,omitempty" json:"executionFrequency,omitempty" xml:"executionFrequency,omitempty"`
	// Message content to be attached
	Messages []*campaignMessagePayload `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue - in minutes
	PollingInterval *int `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Product Id for which the campaign is created
	ProductID *string `form:"productId,omitempty" json:"productId,omitempty" xml:"productId,omitempty"`
	// Start date of the Campaign
	StartDate *int `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
}

// Finalize sets the default values for campaignPayload type instance.
func (ut *campaignPayload) Finalize() {
	var defaultExecutionFrequency = 1
	if ut.ExecutionFrequency == nil {
		ut.ExecutionFrequency = &defaultExecutionFrequency
	}
}

// Validate validates the campaignPayload type instance.
func (ut *campaignPayload) Validate() (err error) {
	if ut.ProductID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "productId"))
	}
	if ut.StartDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "startDate"))
	}
	if ut.EndDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "endDate"))
	}
	if ut.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "messages"))
	}
	if ut.ActiveStartHour != nil {
		if *ut.ActiveStartHour < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.activeStartHour`, *ut.ActiveStartHour, 0, true))
		}
	}
	if ut.ActiveStartHour != nil {
		if *ut.ActiveStartHour > 23 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.activeStartHour`, *ut.ActiveStartHour, 23, false))
		}
	}
	if ut.ActiveStartMinute != nil {
		if *ut.ActiveStartMinute < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.activeStartMinute`, *ut.ActiveStartMinute, 0, true))
		}
	}
	if ut.ActiveStartMinute != nil {
		if *ut.ActiveStartMinute > 59 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.activeStartMinute`, *ut.ActiveStartMinute, 59, false))
		}
	}
	if ut.ExecutionFrequency != nil {
		if *ut.ExecutionFrequency < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.executionFrequency`, *ut.ExecutionFrequency, 1, true))
		}
	}
	for _, e := range ut.Messages {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if ut.ProductID != nil {
		if utf8.RuneCountInString(*ut.ProductID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.productId`, *ut.ProductID, utf8.RuneCountInString(*ut.ProductID), 1, true))
		}
	}
	if ut.ProductID != nil {
		if utf8.RuneCountInString(*ut.ProductID) > 36 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.productId`, *ut.ProductID, utf8.RuneCountInString(*ut.ProductID), 36, false))
		}
	}
	return
}

// Publicize creates CampaignPayload from campaignPayload
func (ut *campaignPayload) Publicize() *CampaignPayload {
	var pub CampaignPayload
	if ut.ActiveHours != nil {
		pub.ActiveHours = ut.ActiveHours
	}
	if ut.ActiveStartHour != nil {
		pub.ActiveStartHour = ut.ActiveStartHour
	}
	if ut.ActiveStartMinute != nil {
		pub.ActiveStartMinute = ut.ActiveStartMinute
	}
	if ut.EndDate != nil {
		pub.EndDate = *ut.EndDate
	}
	if ut.ExecutionFrequency != nil {
		pub.ExecutionFrequency = *ut.ExecutionFrequency
	}
	if ut.Messages != nil {
		pub.Messages = make([]*CampaignMessagePayload, len(ut.Messages))
		for i2, elem2 := range ut.Messages {
			pub.Messages[i2] = elem2.Publicize()
		}
	}
	if ut.PollingInterval != nil {
		pub.PollingInterval = ut.PollingInterval
	}
	if ut.ProductID != nil {
		pub.ProductID = *ut.ProductID
	}
	if ut.StartDate != nil {
		pub.StartDate = *ut.StartDate
	}
	return &pub
}

// The Campaign object
type CampaignPayload struct {
	// Duration in which campaign will be running starting from activeStartTime -in minutes
	ActiveHours *int `form:"activeHours,omitempty" json:"activeHours,omitempty" xml:"activeHours,omitempty"`
	// The active start hour -  campaign execution will be starting from this time
	ActiveStartHour *int `form:"activeStartHour,omitempty" json:"activeStartHour,omitempty" xml:"activeStartHour,omitempty"`
	// The active start minutes
	ActiveStartMinute *int `form:"activeStartMinute,omitempty" json:"activeStartMinute,omitempty" xml:"activeStartMinute,omitempty"`
	// End date of the Campaign
	EndDate int `form:"endDate" json:"endDate" xml:"endDate"`
	// Frequency in which the campaign needs to be executed - in days
	ExecutionFrequency int `form:"executionFrequency" json:"executionFrequency" xml:"executionFrequency"`
	// Message content to be attached
	Messages []*CampaignMessagePayload `form:"messages" json:"messages" xml:"messages"`
	// Interval in which the campaign need to poll the lead queue - in minutes
	PollingInterval *int `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Product Id for which the campaign is created
	ProductID string `form:"productId" json:"productId" xml:"productId"`
	// Start date of the Campaign
	StartDate int `form:"startDate" json:"startDate" xml:"startDate"`
}

// Validate validates the CampaignPayload type instance.
func (ut *CampaignPayload) Validate() (err error) {
	if ut.ProductID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "productId"))
	}

	if ut.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "messages"))
	}
	if ut.ActiveStartHour != nil {
		if *ut.ActiveStartHour < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.activeStartHour`, *ut.ActiveStartHour, 0, true))
		}
	}
	if ut.ActiveStartHour != nil {
		if *ut.ActiveStartHour > 23 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.activeStartHour`, *ut.ActiveStartHour, 23, false))
		}
	}
	if ut.ActiveStartMinute != nil {
		if *ut.ActiveStartMinute < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.activeStartMinute`, *ut.ActiveStartMinute, 0, true))
		}
	}
	if ut.ActiveStartMinute != nil {
		if *ut.ActiveStartMinute > 59 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`type.activeStartMinute`, *ut.ActiveStartMinute, 59, false))
		}
	}
	if ut.ExecutionFrequency < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.executionFrequency`, ut.ExecutionFrequency, 1, true))
	}
	for _, e := range ut.Messages {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if utf8.RuneCountInString(ut.ProductID) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.productId`, ut.ProductID, utf8.RuneCountInString(ut.ProductID), 1, true))
	}
	if utf8.RuneCountInString(ut.ProductID) > 36 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.productId`, ut.ProductID, utf8.RuneCountInString(ut.ProductID), 36, false))
	}
	return
}

// The Campaign object
type campaignUpdatePayload struct {
	// Duration in which campaign will be running starting from activeStartTime -in minutes
	ActiveHours *int `form:"activeHours,omitempty" json:"activeHours,omitempty" xml:"activeHours,omitempty"`
	// The active start hour -  campaign execution will be starting from this time
	ActiveStartHour *int `form:"activeStartHour,omitempty" json:"activeStartHour,omitempty" xml:"activeStartHour,omitempty"`
	// The active start minutes
	ActiveStartMinute *int `form:"activeStartMinute,omitempty" json:"activeStartMinute,omitempty" xml:"activeStartMinute,omitempty"`
	// End date of the campaign
	EndDate *int `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Interval in which the campaign needs to be executed - in days
	ExecutionFrequency *int `form:"executionFrequency,omitempty" json:"executionFrequency,omitempty" xml:"executionFrequency,omitempty"`
	// Message content to be attached
	Messages []*campaignMessagePayload `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *int `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Start time of the campaign
	StartDate *int `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the campaign
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// Validate validates the campaignUpdatePayload type instance.
func (ut *campaignUpdatePayload) Validate() (err error) {
	for _, e := range ut.Messages {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Publicize creates CampaignUpdatePayload from campaignUpdatePayload
func (ut *campaignUpdatePayload) Publicize() *CampaignUpdatePayload {
	var pub CampaignUpdatePayload
	if ut.ActiveHours != nil {
		pub.ActiveHours = ut.ActiveHours
	}
	if ut.ActiveStartHour != nil {
		pub.ActiveStartHour = ut.ActiveStartHour
	}
	if ut.ActiveStartMinute != nil {
		pub.ActiveStartMinute = ut.ActiveStartMinute
	}
	if ut.EndDate != nil {
		pub.EndDate = ut.EndDate
	}
	if ut.ExecutionFrequency != nil {
		pub.ExecutionFrequency = ut.ExecutionFrequency
	}
	if ut.Messages != nil {
		pub.Messages = make([]*CampaignMessagePayload, len(ut.Messages))
		for i2, elem2 := range ut.Messages {
			pub.Messages[i2] = elem2.Publicize()
		}
	}
	if ut.PollingInterval != nil {
		pub.PollingInterval = ut.PollingInterval
	}
	if ut.StartDate != nil {
		pub.StartDate = ut.StartDate
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	return &pub
}

// The Campaign object
type CampaignUpdatePayload struct {
	// Duration in which campaign will be running starting from activeStartTime -in minutes
	ActiveHours *int `form:"activeHours,omitempty" json:"activeHours,omitempty" xml:"activeHours,omitempty"`
	// The active start hour -  campaign execution will be starting from this time
	ActiveStartHour *int `form:"activeStartHour,omitempty" json:"activeStartHour,omitempty" xml:"activeStartHour,omitempty"`
	// The active start minutes
	ActiveStartMinute *int `form:"activeStartMinute,omitempty" json:"activeStartMinute,omitempty" xml:"activeStartMinute,omitempty"`
	// End date of the campaign
	EndDate *int `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Interval in which the campaign needs to be executed - in days
	ExecutionFrequency *int `form:"executionFrequency,omitempty" json:"executionFrequency,omitempty" xml:"executionFrequency,omitempty"`
	// Message content to be attached
	Messages []*CampaignMessagePayload `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *int `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Start time of the campaign
	StartDate *int `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the campaign
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// Validate validates the CampaignUpdatePayload type instance.
func (ut *CampaignUpdatePayload) Validate() (err error) {
	for _, e := range ut.Messages {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// The Message content object
type messageContentPayload struct {
	// The message content
	MessageContent *string `form:"messageContent,omitempty" json:"messageContent,omitempty" xml:"messageContent,omitempty"`
}

// Validate validates the messageContentPayload type instance.
func (ut *messageContentPayload) Validate() (err error) {
	if ut.MessageContent == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "messageContent"))
	}
	if ut.MessageContent != nil {
		if utf8.RuneCountInString(*ut.MessageContent) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.messageContent`, *ut.MessageContent, utf8.RuneCountInString(*ut.MessageContent), 1, true))
		}
	}
	if ut.MessageContent != nil {
		if utf8.RuneCountInString(*ut.MessageContent) > 160 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.messageContent`, *ut.MessageContent, utf8.RuneCountInString(*ut.MessageContent), 160, false))
		}
	}
	return
}

// Publicize creates MessageContentPayload from messageContentPayload
func (ut *messageContentPayload) Publicize() *MessageContentPayload {
	var pub MessageContentPayload
	if ut.MessageContent != nil {
		pub.MessageContent = *ut.MessageContent
	}
	return &pub
}

// The Message content object
type MessageContentPayload struct {
	// The message content
	MessageContent string `form:"messageContent" json:"messageContent" xml:"messageContent"`
}

// Validate validates the MessageContentPayload type instance.
func (ut *MessageContentPayload) Validate() (err error) {
	if ut.MessageContent == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "messageContent"))
	}
	if utf8.RuneCountInString(ut.MessageContent) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.messageContent`, ut.MessageContent, utf8.RuneCountInString(ut.MessageContent), 1, true))
	}
	if utf8.RuneCountInString(ut.MessageContent) > 160 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.messageContent`, ut.MessageContent, utf8.RuneCountInString(ut.MessageContent), 160, false))
	}
	return
}

// The Message content object
type messageContentUpdatePayload struct {
	// The message content
	MessageContent *string `form:"messageContent,omitempty" json:"messageContent,omitempty" xml:"messageContent,omitempty"`
}

// Validate validates the messageContentUpdatePayload type instance.
func (ut *messageContentUpdatePayload) Validate() (err error) {
	if ut.MessageContent == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "messageContent"))
	}
	if ut.MessageContent != nil {
		if utf8.RuneCountInString(*ut.MessageContent) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.messageContent`, *ut.MessageContent, utf8.RuneCountInString(*ut.MessageContent), 1, true))
		}
	}
	if ut.MessageContent != nil {
		if utf8.RuneCountInString(*ut.MessageContent) > 160 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.messageContent`, *ut.MessageContent, utf8.RuneCountInString(*ut.MessageContent), 160, false))
		}
	}
	return
}

// Publicize creates MessageContentUpdatePayload from messageContentUpdatePayload
func (ut *messageContentUpdatePayload) Publicize() *MessageContentUpdatePayload {
	var pub MessageContentUpdatePayload
	if ut.MessageContent != nil {
		pub.MessageContent = *ut.MessageContent
	}
	return &pub
}

// The Message content object
type MessageContentUpdatePayload struct {
	// The message content
	MessageContent string `form:"messageContent" json:"messageContent" xml:"messageContent"`
}

// Validate validates the MessageContentUpdatePayload type instance.
func (ut *MessageContentUpdatePayload) Validate() (err error) {
	if ut.MessageContent == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "messageContent"))
	}
	if utf8.RuneCountInString(ut.MessageContent) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.messageContent`, ut.MessageContent, utf8.RuneCountInString(ut.MessageContent), 1, true))
	}
	if utf8.RuneCountInString(ut.MessageContent) > 160 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.messageContent`, ut.MessageContent, utf8.RuneCountInString(ut.MessageContent), 160, false))
	}
	return
}

// productPayload user type.
type productPayload struct {
	AvailableLocations []*productLocation `form:"availableLocations,omitempty" json:"availableLocations,omitempty" xml:"availableLocations,omitempty"`
	ClientCode         *string            `form:"clientCode,omitempty" json:"clientCode,omitempty" xml:"clientCode,omitempty"`
	Criteria           []*productCriteria `form:"criteria,omitempty" json:"criteria,omitempty" xml:"criteria,omitempty"`
	DailyVolume        *int               `form:"dailyVolume,omitempty" json:"dailyVolume,omitempty" xml:"dailyVolume,omitempty"`
	ProductCode        *string            `form:"productCode,omitempty" json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductType        *string            `form:"productType,omitempty" json:"productType,omitempty" xml:"productType,omitempty"`
}

// Publicize creates ProductPayload from productPayload
func (ut *productPayload) Publicize() *ProductPayload {
	var pub ProductPayload
	if ut.AvailableLocations != nil {
		pub.AvailableLocations = make([]*ProductLocation, len(ut.AvailableLocations))
		for i2, elem2 := range ut.AvailableLocations {
			pub.AvailableLocations[i2] = elem2.Publicize()
		}
	}
	if ut.ClientCode != nil {
		pub.ClientCode = ut.ClientCode
	}
	if ut.Criteria != nil {
		pub.Criteria = make([]*ProductCriteria, len(ut.Criteria))
		for i2, elem2 := range ut.Criteria {
			pub.Criteria[i2] = elem2.Publicize()
		}
	}
	if ut.DailyVolume != nil {
		pub.DailyVolume = ut.DailyVolume
	}
	if ut.ProductCode != nil {
		pub.ProductCode = ut.ProductCode
	}
	if ut.ProductType != nil {
		pub.ProductType = ut.ProductType
	}
	return &pub
}

// ProductPayload user type.
type ProductPayload struct {
	AvailableLocations []*ProductLocation `form:"availableLocations,omitempty" json:"availableLocations,omitempty" xml:"availableLocations,omitempty"`
	ClientCode         *string            `form:"clientCode,omitempty" json:"clientCode,omitempty" xml:"clientCode,omitempty"`
	Criteria           []*ProductCriteria `form:"criteria,omitempty" json:"criteria,omitempty" xml:"criteria,omitempty"`
	DailyVolume        *int               `form:"dailyVolume,omitempty" json:"dailyVolume,omitempty" xml:"dailyVolume,omitempty"`
	ProductCode        *string            `form:"productCode,omitempty" json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductType        *string            `form:"productType,omitempty" json:"productType,omitempty" xml:"productType,omitempty"`
}

// customerServicePayload user type.
type customerServicePayload struct {
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

// Publicize creates CustomerServicePayload from customerServicePayload
func (ut *customerServicePayload) Publicize() *CustomerServicePayload {
	var pub CustomerServicePayload
	if ut.PhoneNumber != nil {
		pub.PhoneNumber = ut.PhoneNumber
	}
	return &pub
}

// CustomerServicePayload user type.
type CustomerServicePayload struct {
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

// productCriteria user type.
type productCriteria struct {
	ComparisonType *string `form:"comparisonType,omitempty" json:"comparisonType,omitempty" xml:"comparisonType,omitempty"`
	MaxValue       *int    `form:"maxValue,omitempty" json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue       *int    `form:"minValue,omitempty" json:"minValue,omitempty" xml:"minValue,omitempty"`
	Variable       *string `form:"variable,omitempty" json:"variable,omitempty" xml:"variable,omitempty"`
}

// Publicize creates ProductCriteria from productCriteria
func (ut *productCriteria) Publicize() *ProductCriteria {
	var pub ProductCriteria
	if ut.ComparisonType != nil {
		pub.ComparisonType = ut.ComparisonType
	}
	if ut.MaxValue != nil {
		pub.MaxValue = ut.MaxValue
	}
	if ut.MinValue != nil {
		pub.MinValue = ut.MinValue
	}
	if ut.Variable != nil {
		pub.Variable = ut.Variable
	}
	return &pub
}

// ProductCriteria user type.
type ProductCriteria struct {
	ComparisonType *string `form:"comparisonType,omitempty" json:"comparisonType,omitempty" xml:"comparisonType,omitempty"`
	MaxValue       *int    `form:"maxValue,omitempty" json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue       *int    `form:"minValue,omitempty" json:"minValue,omitempty" xml:"minValue,omitempty"`
	Variable       *string `form:"variable,omitempty" json:"variable,omitempty" xml:"variable,omitempty"`
}

// productLocation user type.
type productLocation struct {
	District *string `form:"district,omitempty" json:"district,omitempty" xml:"district,omitempty"`
	Province *string `form:"province,omitempty" json:"province,omitempty" xml:"province,omitempty"`
}

// Publicize creates ProductLocation from productLocation
func (ut *productLocation) Publicize() *ProductLocation {
	var pub ProductLocation
	if ut.District != nil {
		pub.District = ut.District
	}
	if ut.Province != nil {
		pub.Province = ut.Province
	}
	return &pub
}

// ProductLocation user type.
type ProductLocation struct {
	District *string `form:"district,omitempty" json:"district,omitempty" xml:"district,omitempty"`
	Province *string `form:"province,omitempty" json:"province,omitempty" xml:"province,omitempty"`
}

// smsPayload user type.
type smsPayload struct {
	// The url of callback api for response
	CallbackApiforResponse *string `form:"callbackApiforResponse,omitempty" json:"callbackApiforResponse,omitempty" xml:"callbackApiforResponse,omitempty"`
	// The url of callback api for send
	CallbackApiforSend *string `form:"callbackApiforSend,omitempty" json:"callbackApiforSend,omitempty" xml:"callbackApiforSend,omitempty"`
	//  the campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// The content of sms
	MessageContent *string `form:"messageContent,omitempty" json:"messageContent,omitempty" xml:"messageContent,omitempty"`
	// phone number to sent sms
	PhoneNumber *int `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

// Publicize creates SmsPayload from smsPayload
func (ut *smsPayload) Publicize() *SmsPayload {
	var pub SmsPayload
	if ut.CallbackApiforResponse != nil {
		pub.CallbackApiforResponse = ut.CallbackApiforResponse
	}
	if ut.CallbackApiforSend != nil {
		pub.CallbackApiforSend = ut.CallbackApiforSend
	}
	if ut.CampaignID != nil {
		pub.CampaignID = ut.CampaignID
	}
	if ut.MessageContent != nil {
		pub.MessageContent = ut.MessageContent
	}
	if ut.PhoneNumber != nil {
		pub.PhoneNumber = ut.PhoneNumber
	}
	return &pub
}

// SmsPayload user type.
type SmsPayload struct {
	// The url of callback api for response
	CallbackApiforResponse *string `form:"callbackApiforResponse,omitempty" json:"callbackApiforResponse,omitempty" xml:"callbackApiforResponse,omitempty"`
	// The url of callback api for send
	CallbackApiforSend *string `form:"callbackApiforSend,omitempty" json:"callbackApiforSend,omitempty" xml:"callbackApiforSend,omitempty"`
	//  the campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// The content of sms
	MessageContent *string `form:"messageContent,omitempty" json:"messageContent,omitempty" xml:"messageContent,omitempty"`
	// phone number to sent sms
	PhoneNumber *int `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}
