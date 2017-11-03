// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": smstrackerservice Resource Client
//
// Command:
// $ goagen
// --design=campaign/design
// --out=$(GOPATH)/src/campaign
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateSmstrackerservicePath computes a request path to the create action of smstrackerservice.
func CreateSmstrackerservicePath() string {

	return fmt.Sprintf("/smstrackerservice/")
}

// To send an sms
func (c *Client) CreateSmstrackerservice(ctx context.Context, path string, payload *SmsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateSmstrackerserviceRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateSmstrackerserviceRequest create the request corresponding to the create action endpoint of the smstrackerservice resource.
func (c *Client) NewCreateSmstrackerserviceRequest(ctx context.Context, path string, payload *SmsPayload, contentType string) (*http.Request, error) {
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
