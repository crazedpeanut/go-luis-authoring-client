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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteExampleParams creates a new DeleteExampleParams object
// with the default values initialized.
func NewDeleteExampleParams() *DeleteExampleParams {
	var ()
	return &DeleteExampleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExampleParamsWithTimeout creates a new DeleteExampleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteExampleParamsWithTimeout(timeout time.Duration) *DeleteExampleParams {
	var ()
	return &DeleteExampleParams{

		timeout: timeout,
	}
}

// NewDeleteExampleParamsWithContext creates a new DeleteExampleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteExampleParamsWithContext(ctx context.Context) *DeleteExampleParams {
	var ()
	return &DeleteExampleParams{

		Context: ctx,
	}
}

// NewDeleteExampleParamsWithHTTPClient creates a new DeleteExampleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteExampleParamsWithHTTPClient(client *http.Client) *DeleteExampleParams {
	var ()
	return &DeleteExampleParams{
		HTTPClient: client,
	}
}

/*DeleteExampleParams contains all the parameters to send to the API endpoint
for the delete example operation typically these are written to a http.Request
*/
type DeleteExampleParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ExampleID
	  The example ID.

	*/
	ExampleID int64
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete example params
func (o *DeleteExampleParams) WithTimeout(timeout time.Duration) *DeleteExampleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete example params
func (o *DeleteExampleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete example params
func (o *DeleteExampleParams) WithContext(ctx context.Context) *DeleteExampleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete example params
func (o *DeleteExampleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete example params
func (o *DeleteExampleParams) WithHTTPClient(client *http.Client) *DeleteExampleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete example params
func (o *DeleteExampleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete example params
func (o *DeleteExampleParams) WithAppID(appID string) *DeleteExampleParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete example params
func (o *DeleteExampleParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithExampleID adds the exampleID to the delete example params
func (o *DeleteExampleParams) WithExampleID(exampleID int64) *DeleteExampleParams {
	o.SetExampleID(exampleID)
	return o
}

// SetExampleID adds the exampleId to the delete example params
func (o *DeleteExampleParams) SetExampleID(exampleID int64) {
	o.ExampleID = exampleID
}

// WithVersionID adds the versionID to the delete example params
func (o *DeleteExampleParams) WithVersionID(versionID string) *DeleteExampleParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete example params
func (o *DeleteExampleParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExampleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param exampleId
	if err := r.SetPathParam("exampleId", swag.FormatInt64(o.ExampleID)); err != nil {
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
