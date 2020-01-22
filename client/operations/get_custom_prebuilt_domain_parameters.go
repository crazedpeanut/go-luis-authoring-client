// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCustomPrebuiltDomainParams creates a new GetCustomPrebuiltDomainParams object
// with the default values initialized.
func NewGetCustomPrebuiltDomainParams() *GetCustomPrebuiltDomainParams {

	return &GetCustomPrebuiltDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomPrebuiltDomainParamsWithTimeout creates a new GetCustomPrebuiltDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomPrebuiltDomainParamsWithTimeout(timeout time.Duration) *GetCustomPrebuiltDomainParams {

	return &GetCustomPrebuiltDomainParams{

		timeout: timeout,
	}
}

// NewGetCustomPrebuiltDomainParamsWithContext creates a new GetCustomPrebuiltDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomPrebuiltDomainParamsWithContext(ctx context.Context) *GetCustomPrebuiltDomainParams {

	return &GetCustomPrebuiltDomainParams{

		Context: ctx,
	}
}

// NewGetCustomPrebuiltDomainParamsWithHTTPClient creates a new GetCustomPrebuiltDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomPrebuiltDomainParamsWithHTTPClient(client *http.Client) *GetCustomPrebuiltDomainParams {

	return &GetCustomPrebuiltDomainParams{
		HTTPClient: client,
	}
}

/*GetCustomPrebuiltDomainParams contains all the parameters to send to the API endpoint
for the get custom prebuilt domain operation typically these are written to a http.Request
*/
type GetCustomPrebuiltDomainParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get custom prebuilt domain params
func (o *GetCustomPrebuiltDomainParams) WithTimeout(timeout time.Duration) *GetCustomPrebuiltDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom prebuilt domain params
func (o *GetCustomPrebuiltDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom prebuilt domain params
func (o *GetCustomPrebuiltDomainParams) WithContext(ctx context.Context) *GetCustomPrebuiltDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom prebuilt domain params
func (o *GetCustomPrebuiltDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom prebuilt domain params
func (o *GetCustomPrebuiltDomainParams) WithHTTPClient(client *http.Client) *GetCustomPrebuiltDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom prebuilt domain params
func (o *GetCustomPrebuiltDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomPrebuiltDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
