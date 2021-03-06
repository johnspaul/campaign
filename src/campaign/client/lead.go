// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "campaign": lead Resource Client
//
// Command:
// $ goagen
// --design=campaign/design
// --out=$(GOPATH)/src/campaign
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetLeadPath computes a request path to the get action of lead.
func GetLeadPath(productID string) string {
	param0 := productID

	return fmt.Sprintf("/lead/%s", param0)
}

// GetLead makes a request to the get action endpoint of the lead resource
func (c *Client) GetLead(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetLeadRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetLeadRequest create the request corresponding to the get action endpoint of the lead resource.
func (c *Client) NewGetLeadRequest(ctx context.Context, path string) (*http.Request, error) {
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
