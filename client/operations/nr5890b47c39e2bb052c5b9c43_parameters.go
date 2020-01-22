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

	models "github.com/crazedpeanut/go-luis-authoring-client/models"
)

// NewNr5890b47c39e2bb052c5b9c43Params creates a new Nr5890b47c39e2bb052c5b9c43Params object
// with the default values initialized.
func NewNr5890b47c39e2bb052c5b9c43Params() *Nr5890b47c39e2bb052c5b9c43Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c43Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c43ParamsWithTimeout creates a new Nr5890b47c39e2bb052c5b9c43Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5890b47c39e2bb052c5b9c43ParamsWithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c43Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c43Params{

		timeout: timeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c43ParamsWithContext creates a new Nr5890b47c39e2bb052c5b9c43Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5890b47c39e2bb052c5b9c43ParamsWithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c43Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c43Params{

		Context: ctx,
	}
}

// NewNr5890b47c39e2bb052c5b9c43ParamsWithHTTPClient creates a new Nr5890b47c39e2bb052c5b9c43Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5890b47c39e2bb052c5b9c43ParamsWithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c43Params {
	var ()
	return &Nr5890b47c39e2bb052c5b9c43Params{
		HTTPClient: client,
	}
}

/*Nr5890b47c39e2bb052c5b9c43Params contains all the parameters to send to the API endpoint
for the 5890b47c39e2bb052c5b9c43 operation typically these are written to a http.Request
*/
type Nr5890b47c39e2bb052c5b9c43Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ExternalKeyUpdateObject
	  The external api key to be assigned.

	*/
	ExternalKeyUpdateObject *models.ExternalKeyUpdateObject
	/*VersionID
	  The application version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) WithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c43Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) WithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c43Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) WithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c43Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) WithAppID(appID string) *Nr5890b47c39e2bb052c5b9c43Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithExternalKeyUpdateObject adds the externalKeyUpdateObject to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) WithExternalKeyUpdateObject(externalKeyUpdateObject *models.ExternalKeyUpdateObject) *Nr5890b47c39e2bb052c5b9c43Params {
	o.SetExternalKeyUpdateObject(externalKeyUpdateObject)
	return o
}

// SetExternalKeyUpdateObject adds the externalKeyUpdateObject to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) SetExternalKeyUpdateObject(externalKeyUpdateObject *models.ExternalKeyUpdateObject) {
	o.ExternalKeyUpdateObject = externalKeyUpdateObject
}

// WithVersionID adds the versionID to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) WithVersionID(versionID string) *Nr5890b47c39e2bb052c5b9c43Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5890b47c39e2bb052c5b9c43 params
func (o *Nr5890b47c39e2bb052c5b9c43Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5890b47c39e2bb052c5b9c43Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.ExternalKeyUpdateObject != nil {
		if err := r.SetBodyParam(o.ExternalKeyUpdateObject); err != nil {
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
