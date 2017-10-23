// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": Application User Types
//
// Command:
// $ goagen
// --design=campaign/design
// --out=$(GOPATH)/src/campaign
// --regen=true
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// Message  attached to a campaign
type campaignMessagePayload struct {
	// The campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// The message id
	MessageID *string `form:"messageId,omitempty" json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The percentage pf this message to be used
	Percentage *float64 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// Validate validates the campaignMessagePayload type instance.
func (ut *campaignMessagePayload) Validate() (err error) {
	if ut.CampaignID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "campaignId"))
	}
	if ut.MessageID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "messageId"))
	}
	if ut.Percentage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "percentage"))
	}
	if ut.CampaignID != nil {
		if utf8.RuneCountInString(*ut.CampaignID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.campaignId`, *ut.CampaignID, utf8.RuneCountInString(*ut.CampaignID), 1, true))
		}
	}
	if ut.CampaignID != nil {
		if utf8.RuneCountInString(*ut.CampaignID) > 36 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.campaignId`, *ut.CampaignID, utf8.RuneCountInString(*ut.CampaignID), 36, false))
		}
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
	if ut.CampaignID != nil {
		pub.CampaignID = *ut.CampaignID
	}
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
	// The campaign id
	CampaignID string `form:"campaignId" json:"campaignId" xml:"campaignId"`
	// The message id
	MessageID string `form:"messageId" json:"messageId" xml:"messageId"`
	// The percentage pf this message to be used
	Percentage float64 `form:"percentage" json:"percentage" xml:"percentage"`
}

// Validate validates the CampaignMessagePayload type instance.
func (ut *CampaignMessagePayload) Validate() (err error) {
	if ut.CampaignID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "campaignId"))
	}
	if ut.MessageID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "messageId"))
	}

	if utf8.RuneCountInString(ut.CampaignID) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.campaignId`, ut.CampaignID, utf8.RuneCountInString(ut.CampaignID), 1, true))
	}
	if utf8.RuneCountInString(ut.CampaignID) > 36 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.campaignId`, ut.CampaignID, utf8.RuneCountInString(ut.CampaignID), 36, false))
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
	// End date of the campaign
	EndDate *time.Time `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Message content to be attached
	Messages []*campaignMessagePayload `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *float64 `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Product Id for which the campaign is created
	ProductID *string `form:"productId,omitempty" json:"productId,omitempty" xml:"productId,omitempty"`
	// Start time of the campaign
	StartDate *time.Time `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the campaign
	State *int `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the campaignPayload type instance.
func (ut *campaignPayload) Validate() (err error) {
	if ut.ProductID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "productId"))
	}
	if ut.StartDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "startDate"))
	}
	if ut.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "messages"))
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
	if ut.EndDate != nil {
		pub.EndDate = ut.EndDate
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
	if ut.State != nil {
		pub.State = ut.State
	}
	return &pub
}

// The Campaign object
type CampaignPayload struct {
	// End date of the campaign
	EndDate *time.Time `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Message content to be attached
	Messages []*CampaignMessagePayload `form:"messages" json:"messages" xml:"messages"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *float64 `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Product Id for which the campaign is created
	ProductID string `form:"productId" json:"productId" xml:"productId"`
	// Start time of the campaign
	StartDate time.Time `form:"startDate" json:"startDate" xml:"startDate"`
	// State of the campaign
	State *int `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the CampaignPayload type instance.
func (ut *CampaignPayload) Validate() (err error) {
	if ut.ProductID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "productId"))
	}

	if ut.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "messages"))
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
	// campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// End date of the campaign
	EndDate *time.Time `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Message content to be attached
	Messages []*campaignMessagePayload `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *float64 `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Start time of the campaign
	StartDate *time.Time `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the campaign
	State *int `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
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
	if ut.CampaignID != nil {
		pub.CampaignID = ut.CampaignID
	}
	if ut.EndDate != nil {
		pub.EndDate = ut.EndDate
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
	if ut.State != nil {
		pub.State = ut.State
	}
	return &pub
}

// The Campaign object
type CampaignUpdatePayload struct {
	// campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// End date of the campaign
	EndDate *time.Time `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Message content to be attached
	Messages []*CampaignMessagePayload `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *float64 `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Start time of the campaign
	StartDate *time.Time `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the campaign
	State *int `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
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