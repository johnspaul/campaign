// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": Application Media Types
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
	"net/http"
	"time"
)

// Campaign media type (default view)
//
// Identifier: application/ts.campaign; view=default
type Campaign struct {
	// Campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// End date of the Campaign
	EndDate *int `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *float64 `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Start date of the Campaign
	StartDate *int `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the Campaign
	State *int `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Campaign media type (detailed view)
//
// Identifier: application/ts.campaign; view=detailed
type CampaignDetailed struct {
	// Campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// End date of the Campaign
	EndDate *int `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Message content to be attached
	Messages []*CampaignMessage `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
	// Interval in which the campaign need to poll the lead queue
	PollingInterval *float64 `form:"pollingInterval,omitempty" json:"pollingInterval,omitempty" xml:"pollingInterval,omitempty"`
	// Start date of the Campaign
	StartDate *int `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// State of the Campaign
	State *int `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the CampaignDetailed media type instance.
func (mt *CampaignDetailed) Validate() (err error) {
	for _, e := range mt.Messages {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeCampaign decodes the Campaign instance encoded in resp body.
func (c *Client) DecodeCampaign(resp *http.Response) (*Campaign, error) {
	var decoded Campaign
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCampaignDetailed decodes the CampaignDetailed instance encoded in resp body.
func (c *Client) DecodeCampaignDetailed(resp *http.Response) (*CampaignDetailed, error) {
	var decoded CampaignDetailed
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CampaignMessage media type (default view)
//
// Identifier: application/ts.campaign.message; view=default
type CampaignMessage struct {
	// campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// The message
	Message *MessageContentPayload `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Percentage of sms to be made out of this message
	Percentage *string `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// Validate validates the CampaignMessage media type instance.
func (mt *CampaignMessage) Validate() (err error) {
	if mt.Message != nil {
		if err2 := mt.Message.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeCampaignMessage decodes the CampaignMessage instance encoded in resp body.
func (c *Client) DecodeCampaignMessage(resp *http.Response) (*CampaignMessage, error) {
	var decoded CampaignMessage
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CampaignMessageContent media type (default view)
//
// Identifier: application/ts.campaign.messagecontent; view=default
type CampaignMessageContent struct {
	// The Message
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Message content  id
	MessageID *string `form:"messageId,omitempty" json:"messageId,omitempty" xml:"messageId,omitempty"`
}

// DecodeCampaignMessageContent decodes the CampaignMessageContent instance encoded in resp body.
func (c *Client) DecodeCampaignMessageContent(resp *http.Response) (*CampaignMessageContent, error) {
	var decoded CampaignMessageContent
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CampaignExecutionContext media type (default view)
//
// Identifier: application/ts.campaignexecution; view=default
type CampaignExecutionContext struct {
	// campaign id
	CampaignID *string `form:"campaignId,omitempty" json:"campaignId,omitempty" xml:"campaignId,omitempty"`
	// execution endTime time
	EndTime *time.Time `form:"endTime,omitempty" json:"endTime,omitempty" xml:"endTime,omitempty"`
	// execution id
	ExecutionID *string `form:"executionId,omitempty" json:"executionId,omitempty" xml:"executionId,omitempty"`
	// Number of message sent in this execution
	NumMessagesSent *float64 `form:"numMessagesSent,omitempty" json:"numMessagesSent,omitempty" xml:"numMessagesSent,omitempty"`
	// execution start time
	StartTime *time.Time `form:"startTime,omitempty" json:"startTime,omitempty" xml:"startTime,omitempty"`
}

// DecodeCampaignExecutionContext decodes the CampaignExecutionContext instance encoded in resp body.
func (c *Client) DecodeCampaignExecutionContext(resp *http.Response) (*CampaignExecutionContext, error) {
	var decoded CampaignExecutionContext
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}