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

// NewNr591045b75aca2f0b48c76bdcParams creates a new Nr591045b75aca2f0b48c76bdcParams object
// with the default values initialized.
func NewNr591045b75aca2f0b48c76bdcParams() *Nr591045b75aca2f0b48c76bdcParams {
	var ()
	return &Nr591045b75aca2f0b48c76bdcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr591045b75aca2f0b48c76bdcParamsWithTimeout creates a new Nr591045b75aca2f0b48c76bdcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr591045b75aca2f0b48c76bdcParamsWithTimeout(timeout time.Duration) *Nr591045b75aca2f0b48c76bdcParams {
	var ()
	return &Nr591045b75aca2f0b48c76bdcParams{

		timeout: timeout,
	}
}

// NewNr591045b75aca2f0b48c76bdcParamsWithContext creates a new Nr591045b75aca2f0b48c76bdcParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr591045b75aca2f0b48c76bdcParamsWithContext(ctx context.Context) *Nr591045b75aca2f0b48c76bdcParams {
	var ()
	return &Nr591045b75aca2f0b48c76bdcParams{

		Context: ctx,
	}
}

// NewNr591045b75aca2f0b48c76bdcParamsWithHTTPClient creates a new Nr591045b75aca2f0b48c76bdcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr591045b75aca2f0b48c76bdcParamsWithHTTPClient(client *http.Client) *Nr591045b75aca2f0b48c76bdcParams {
	var ()
	return &Nr591045b75aca2f0b48c76bdcParams{
		HTTPClient: client,
	}
}

/*Nr591045b75aca2f0b48c76bdcParams contains all the parameters to send to the API endpoint
for the 591045b75aca2f0b48c76bdc operation typically these are written to a http.Request
*/
type Nr591045b75aca2f0b48c76bdcParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*PrebuiltDomainCreateBaseObject
	  A prebuilt domain create object containing the name of the domain

	*/
	PrebuiltDomainCreateBaseObject *models.PrebuiltDomainCreateBaseObject
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) WithTimeout(timeout time.Duration) *Nr591045b75aca2f0b48c76bdcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) WithContext(ctx context.Context) *Nr591045b75aca2f0b48c76bdcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) WithHTTPClient(client *http.Client) *Nr591045b75aca2f0b48c76bdcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) WithAppID(appID string) *Nr591045b75aca2f0b48c76bdcParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithPrebuiltDomainCreateBaseObject adds the prebuiltDomainCreateBaseObject to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) WithPrebuiltDomainCreateBaseObject(prebuiltDomainCreateBaseObject *models.PrebuiltDomainCreateBaseObject) *Nr591045b75aca2f0b48c76bdcParams {
	o.SetPrebuiltDomainCreateBaseObject(prebuiltDomainCreateBaseObject)
	return o
}

// SetPrebuiltDomainCreateBaseObject adds the prebuiltDomainCreateBaseObject to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) SetPrebuiltDomainCreateBaseObject(prebuiltDomainCreateBaseObject *models.PrebuiltDomainCreateBaseObject) {
	o.PrebuiltDomainCreateBaseObject = prebuiltDomainCreateBaseObject
}

// WithVersionID adds the versionID to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) WithVersionID(versionID string) *Nr591045b75aca2f0b48c76bdcParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 591045b75aca2f0b48c76bdc params
func (o *Nr591045b75aca2f0b48c76bdcParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr591045b75aca2f0b48c76bdcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.PrebuiltDomainCreateBaseObject != nil {
		if err := r.SetBodyParam(o.PrebuiltDomainCreateBaseObject); err != nil {
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