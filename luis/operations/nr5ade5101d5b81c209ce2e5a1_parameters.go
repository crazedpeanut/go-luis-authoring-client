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

// NewNr5ade5101d5b81c209ce2e5a1Params creates a new Nr5ade5101d5b81c209ce2e5a1Params object
// with the default values initialized.
func NewNr5ade5101d5b81c209ce2e5a1Params() *Nr5ade5101d5b81c209ce2e5a1Params {
	var ()
	return &Nr5ade5101d5b81c209ce2e5a1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade5101d5b81c209ce2e5a1ParamsWithTimeout creates a new Nr5ade5101d5b81c209ce2e5a1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade5101d5b81c209ce2e5a1ParamsWithTimeout(timeout time.Duration) *Nr5ade5101d5b81c209ce2e5a1Params {
	var ()
	return &Nr5ade5101d5b81c209ce2e5a1Params{

		timeout: timeout,
	}
}

// NewNr5ade5101d5b81c209ce2e5a1ParamsWithContext creates a new Nr5ade5101d5b81c209ce2e5a1Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade5101d5b81c209ce2e5a1ParamsWithContext(ctx context.Context) *Nr5ade5101d5b81c209ce2e5a1Params {
	var ()
	return &Nr5ade5101d5b81c209ce2e5a1Params{

		Context: ctx,
	}
}

// NewNr5ade5101d5b81c209ce2e5a1ParamsWithHTTPClient creates a new Nr5ade5101d5b81c209ce2e5a1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade5101d5b81c209ce2e5a1ParamsWithHTTPClient(client *http.Client) *Nr5ade5101d5b81c209ce2e5a1Params {
	var ()
	return &Nr5ade5101d5b81c209ce2e5a1Params{
		HTTPClient: client,
	}
}

/*Nr5ade5101d5b81c209ce2e5a1Params contains all the parameters to send to the API endpoint
for the 5ade5101d5b81c209ce2e5a1 operation typically these are written to a http.Request
*/
type Nr5ade5101d5b81c209ce2e5a1Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Body
	  A model object containing the name and the explicit list for the new Pattern.Any entity extractor.

	The name of the regex entity must be unique in the application and must not be used by a prebuilt entity.

	*/
	Body interface{}
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WithTimeout(timeout time.Duration) *Nr5ade5101d5b81c209ce2e5a1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WithContext(ctx context.Context) *Nr5ade5101d5b81c209ce2e5a1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WithHTTPClient(client *http.Client) *Nr5ade5101d5b81c209ce2e5a1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WithAppID(appID string) *Nr5ade5101d5b81c209ce2e5a1Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithBody adds the body to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WithBody(body interface{}) *Nr5ade5101d5b81c209ce2e5a1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) SetBody(body interface{}) {
	o.Body = body
}

// WithVersionID adds the versionID to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WithVersionID(versionID string) *Nr5ade5101d5b81c209ce2e5a1Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade5101d5b81c209ce2e5a1 params
func (o *Nr5ade5101d5b81c209ce2e5a1Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade5101d5b81c209ce2e5a1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}