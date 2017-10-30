// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": campaigns Resource Client
//
// Command:
// $ goagen
// --design=campaign/design
// --out=$(GOPATH)/src/campaign
// --regen=true
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateCampaignsPath computes a request path to the create action of campaigns.
func CreateCampaignsPath() string {

	return fmt.Sprintf("/campaigns/")
}

// Creates a campaign.
func (c *Client) CreateCampaigns(ctx context.Context, path string, payload *CampaignPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCampaignsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCampaignsRequest create the request corresponding to the create action endpoint of the campaigns resource.
func (c *Client) NewCreateCampaignsRequest(ctx context.Context, path string, payload *CampaignPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteCampaignsPath computes a request path to the delete action of campaigns.
func DeleteCampaignsPath(campaignID string) string {
	param0 := campaignID

	return fmt.Sprintf("/campaigns/%s", param0)
}

// Deletes a terminated campaign.
func (c *Client) DeleteCampaigns(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCampaignsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCampaignsRequest create the request corresponding to the delete action endpoint of the campaigns resource.
func (c *Client) NewDeleteCampaignsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetCampaignsPath computes a request path to the get action of campaigns.
func GetCampaignsPath(campaignID string) string {
	param0 := campaignID

	return fmt.Sprintf("/campaigns/%s", param0)
}

// Returns a campaign with  specific id.
func (c *Client) GetCampaigns(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetCampaignsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCampaignsRequest create the request corresponding to the get action endpoint of the campaigns resource.
func (c *Client) NewGetCampaignsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetAllCampaignsPath computes a request path to the getAll action of campaigns.
func GetAllCampaignsPath() string {

	return fmt.Sprintf("/campaigns/")
}

// Returns all the campaigns
func (c *Client) GetAllCampaigns(ctx context.Context, path string, state *float64) (*http.Response, error) {
	req, err := c.NewGetAllCampaignsRequest(ctx, path, state)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAllCampaignsRequest create the request corresponding to the getAll action endpoint of the campaigns resource.
func (c *Client) NewGetAllCampaignsRequest(ctx context.Context, path string, state *float64) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if state != nil {
		tmp18 := strconv.FormatFloat(*state, 'f', -1, 64)
		values.Set("state", tmp18)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetAllCampaignExecutionCampaignsPath computes a request path to the getAllCampaignExecution action of campaigns.
func GetAllCampaignExecutionCampaignsPath(campaignID string) string {
	param0 := campaignID

	return fmt.Sprintf("/campaigns/%s/executions", param0)
}

// Returns all campaign  execution details.
func (c *Client) GetAllCampaignExecutionCampaigns(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetAllCampaignExecutionCampaignsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAllCampaignExecutionCampaignsRequest create the request corresponding to the getAllCampaignExecution action endpoint of the campaigns resource.
func (c *Client) NewGetAllCampaignExecutionCampaignsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetCampaignExecutionCampaignsPath computes a request path to the getCampaignExecution action of campaigns.
func GetCampaignExecutionCampaignsPath(campaignID string, executionID string) string {
	param0 := campaignID
	param1 := executionID

	return fmt.Sprintf("/campaigns/%s/executions/%s", param0, param1)
}

// Returns a specific campaign execution details
func (c *Client) GetCampaignExecutionCampaigns(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetCampaignExecutionCampaignsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCampaignExecutionCampaignsRequest create the request corresponding to the getCampaignExecution action endpoint of the campaigns resource.
func (c *Client) NewGetCampaignExecutionCampaignsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateCampaignsPath computes a request path to the update action of campaigns.
func UpdateCampaignsPath(campaignID string) string {
	param0 := campaignID

	return fmt.Sprintf("/campaigns/%s", param0)
}

// Updates a campaign.
func (c *Client) UpdateCampaigns(ctx context.Context, path string, payload *CampaignUpdatePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateCampaignsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCampaignsRequest create the request corresponding to the update action endpoint of the campaigns resource.
func (c *Client) NewUpdateCampaignsRequest(ctx context.Context, path string, payload *CampaignUpdatePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
