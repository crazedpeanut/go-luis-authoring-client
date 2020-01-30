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

// NewNr5ade0f32d5b81c209ce2e59dParams creates a new Nr5ade0f32d5b81c209ce2e59dParams object
// with the default values initialized.
func NewNr5ade0f32d5b81c209ce2e59dParams() *Nr5ade0f32d5b81c209ce2e59dParams {
	var ()
	return &Nr5ade0f32d5b81c209ce2e59dParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade0f32d5b81c209ce2e59dParamsWithTimeout creates a new Nr5ade0f32d5b81c209ce2e59dParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade0f32d5b81c209ce2e59dParamsWithTimeout(timeout time.Duration) *Nr5ade0f32d5b81c209ce2e59dParams {
	var ()
	return &Nr5ade0f32d5b81c209ce2e59dParams{

		timeout: timeout,
	}
}

// NewNr5ade0f32d5b81c209ce2e59dParamsWithContext creates a new Nr5ade0f32d5b81c209ce2e59dParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade0f32d5b81c209ce2e59dParamsWithContext(ctx context.Context) *Nr5ade0f32d5b81c209ce2e59dParams {
	var ()
	return &Nr5ade0f32d5b81c209ce2e59dParams{

		Context: ctx,
	}
}

// NewNr5ade0f32d5b81c209ce2e59dParamsWithHTTPClient creates a new Nr5ade0f32d5b81c209ce2e59dParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade0f32d5b81c209ce2e59dParamsWithHTTPClient(client *http.Client) *Nr5ade0f32d5b81c209ce2e59dParams {
	var ()
	return &Nr5ade0f32d5b81c209ce2e59dParams{
		HTTPClient: client,
	}
}

/*Nr5ade0f32d5b81c209ce2e59dParams contains all the parameters to send to the API endpoint
for the 5ade0f32d5b81c209ce2e59d operation typically these are written to a http.Request
*/
type Nr5ade0f32d5b81c209ce2e59dParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Body
	  A model object containing the name for the new role.

	The name of the role must be unique across the application version.

	*/
	Body interface{}
	/*EntityID
	  Format - guid. The regular expression entity ID.

	*/
	EntityID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithTimeout(timeout time.Duration) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithContext(ctx context.Context) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithHTTPClient(client *http.Client) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithAppID(appID string) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithBody adds the body to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithBody(body interface{}) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetBody(body interface{}) {
	o.Body = body
}

// WithEntityID adds the entityID to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithEntityID(entityID string) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithVersionID adds the versionID to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WithVersionID(versionID string) *Nr5ade0f32d5b81c209ce2e59dParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade0f32d5b81c209ce2e59d params
func (o *Nr5ade0f32d5b81c209ce2e59dParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade0f32d5b81c209ce2e59dParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param entityId
	if err := r.SetPathParam("entityId", o.EntityID); err != nil {
		return err
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