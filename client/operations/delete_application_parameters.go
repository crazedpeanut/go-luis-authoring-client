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

// NewDeleteApplicationParams creates a new DeleteApplicationParams object
// with the default values initialized.
func NewDeleteApplicationParams() *DeleteApplicationParams {
	var (
		forceDefault = string("false")
	)
	return &DeleteApplicationParams{
		Force: &forceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteApplicationParamsWithTimeout creates a new DeleteApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteApplicationParamsWithTimeout(timeout time.Duration) *DeleteApplicationParams {
	var (
		forceDefault = string("false")
	)
	return &DeleteApplicationParams{
		Force: &forceDefault,

		timeout: timeout,
	}
}

// NewDeleteApplicationParamsWithContext creates a new DeleteApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteApplicationParamsWithContext(ctx context.Context) *DeleteApplicationParams {
	var (
		forceDefault = string("false")
	)
	return &DeleteApplicationParams{
		Force: &forceDefault,

		Context: ctx,
	}
}

// NewDeleteApplicationParamsWithHTTPClient creates a new DeleteApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteApplicationParamsWithHTTPClient(client *http.Client) *DeleteApplicationParams {
	var (
		forceDefault = string("false")
	)
	return &DeleteApplicationParams{
		Force:      &forceDefault,
		HTTPClient: client,
	}
}

/*DeleteApplicationParams contains all the parameters to send to the API endpoint
for the delete application operation typically these are written to a http.Request
*/
type DeleteApplicationParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Force
	  A flag indicating whether to ignore application dependencies while deleting or not. Default is false.

	*/
	Force *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete application params
func (o *DeleteApplicationParams) WithTimeout(timeout time.Duration) *DeleteApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete application params
func (o *DeleteApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete application params
func (o *DeleteApplicationParams) WithContext(ctx context.Context) *DeleteApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete application params
func (o *DeleteApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete application params
func (o *DeleteApplicationParams) WithHTTPClient(client *http.Client) *DeleteApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete application params
func (o *DeleteApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete application params
func (o *DeleteApplicationParams) WithAppID(appID string) *DeleteApplicationParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete application params
func (o *DeleteApplicationParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithForce adds the force to the delete application params
func (o *DeleteApplicationParams) WithForce(force *string) *DeleteApplicationParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete application params
func (o *DeleteApplicationParams) SetForce(force *string) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Force != nil {

		// query param force
		var qrForce string
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := qrForce
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
