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

	models "github.com/crazedpeanut/luis/models"
)

// NewCloneVersionParams creates a new CloneVersionParams object
// with the default values initialized.
func NewCloneVersionParams() *CloneVersionParams {
	var ()
	return &CloneVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloneVersionParamsWithTimeout creates a new CloneVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloneVersionParamsWithTimeout(timeout time.Duration) *CloneVersionParams {
	var ()
	return &CloneVersionParams{

		timeout: timeout,
	}
}

// NewCloneVersionParamsWithContext creates a new CloneVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloneVersionParamsWithContext(ctx context.Context) *CloneVersionParams {
	var ()
	return &CloneVersionParams{

		Context: ctx,
	}
}

// NewCloneVersionParamsWithHTTPClient creates a new CloneVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloneVersionParamsWithHTTPClient(client *http.Client) *CloneVersionParams {
	var ()
	return &CloneVersionParams{
		HTTPClient: client,
	}
}

/*CloneVersionParams contains all the parameters to send to the API endpoint
for the clone version operation typically these are written to a http.Request
*/
type CloneVersionParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*TaskUpdateObject
	  An object containing the new version ID.

	*/
	TaskUpdateObject *models.TaskUpdateObject
	/*VersionID
	  The application version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the clone version params
func (o *CloneVersionParams) WithTimeout(timeout time.Duration) *CloneVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone version params
func (o *CloneVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone version params
func (o *CloneVersionParams) WithContext(ctx context.Context) *CloneVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone version params
func (o *CloneVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone version params
func (o *CloneVersionParams) WithHTTPClient(client *http.Client) *CloneVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone version params
func (o *CloneVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the clone version params
func (o *CloneVersionParams) WithAppID(appID string) *CloneVersionParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the clone version params
func (o *CloneVersionParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithTaskUpdateObject adds the taskUpdateObject to the clone version params
func (o *CloneVersionParams) WithTaskUpdateObject(taskUpdateObject *models.TaskUpdateObject) *CloneVersionParams {
	o.SetTaskUpdateObject(taskUpdateObject)
	return o
}

// SetTaskUpdateObject adds the taskUpdateObject to the clone version params
func (o *CloneVersionParams) SetTaskUpdateObject(taskUpdateObject *models.TaskUpdateObject) {
	o.TaskUpdateObject = taskUpdateObject
}

// WithVersionID adds the versionID to the clone version params
func (o *CloneVersionParams) WithVersionID(versionID string) *CloneVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the clone version params
func (o *CloneVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *CloneVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.TaskUpdateObject != nil {
		if err := r.SetBodyParam(o.TaskUpdateObject); err != nil {
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
