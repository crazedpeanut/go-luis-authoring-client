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

// NewGetCustomPrebuiltDomainsByCultureParams creates a new GetCustomPrebuiltDomainsByCultureParams object
// with the default values initialized.
func NewGetCustomPrebuiltDomainsByCultureParams() *GetCustomPrebuiltDomainsByCultureParams {
	var ()
	return &GetCustomPrebuiltDomainsByCultureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomPrebuiltDomainsByCultureParamsWithTimeout creates a new GetCustomPrebuiltDomainsByCultureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomPrebuiltDomainsByCultureParamsWithTimeout(timeout time.Duration) *GetCustomPrebuiltDomainsByCultureParams {
	var ()
	return &GetCustomPrebuiltDomainsByCultureParams{

		timeout: timeout,
	}
}

// NewGetCustomPrebuiltDomainsByCultureParamsWithContext creates a new GetCustomPrebuiltDomainsByCultureParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomPrebuiltDomainsByCultureParamsWithContext(ctx context.Context) *GetCustomPrebuiltDomainsByCultureParams {
	var ()
	return &GetCustomPrebuiltDomainsByCultureParams{

		Context: ctx,
	}
}

// NewGetCustomPrebuiltDomainsByCultureParamsWithHTTPClient creates a new GetCustomPrebuiltDomainsByCultureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomPrebuiltDomainsByCultureParamsWithHTTPClient(client *http.Client) *GetCustomPrebuiltDomainsByCultureParams {
	var ()
	return &GetCustomPrebuiltDomainsByCultureParams{
		HTTPClient: client,
	}
}

/*GetCustomPrebuiltDomainsByCultureParams contains all the parameters to send to the API endpoint
for the get custom prebuilt domains by culture operation typically these are written to a http.Request
*/
type GetCustomPrebuiltDomainsByCultureParams struct {

	/*Culture
	  Format - string. Culture

	*/
	Culture string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) WithTimeout(timeout time.Duration) *GetCustomPrebuiltDomainsByCultureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) WithContext(ctx context.Context) *GetCustomPrebuiltDomainsByCultureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) WithHTTPClient(client *http.Client) *GetCustomPrebuiltDomainsByCultureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCulture adds the culture to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) WithCulture(culture string) *GetCustomPrebuiltDomainsByCultureParams {
	o.SetCulture(culture)
	return o
}

// SetCulture adds the culture to the get custom prebuilt domains by culture params
func (o *GetCustomPrebuiltDomainsByCultureParams) SetCulture(culture string) {
	o.Culture = culture
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomPrebuiltDomainsByCultureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param culture
	if err := r.SetPathParam("culture", o.Culture); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
