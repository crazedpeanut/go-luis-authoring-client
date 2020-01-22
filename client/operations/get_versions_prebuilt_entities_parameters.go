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

// NewGetVersionsPrebuiltEntitiesParams creates a new GetVersionsPrebuiltEntitiesParams object
// with the default values initialized.
func NewGetVersionsPrebuiltEntitiesParams() *GetVersionsPrebuiltEntitiesParams {
	var ()
	return &GetVersionsPrebuiltEntitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionsPrebuiltEntitiesParamsWithTimeout creates a new GetVersionsPrebuiltEntitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersionsPrebuiltEntitiesParamsWithTimeout(timeout time.Duration) *GetVersionsPrebuiltEntitiesParams {
	var ()
	return &GetVersionsPrebuiltEntitiesParams{

		timeout: timeout,
	}
}

// NewGetVersionsPrebuiltEntitiesParamsWithContext creates a new GetVersionsPrebuiltEntitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersionsPrebuiltEntitiesParamsWithContext(ctx context.Context) *GetVersionsPrebuiltEntitiesParams {
	var ()
	return &GetVersionsPrebuiltEntitiesParams{

		Context: ctx,
	}
}

// NewGetVersionsPrebuiltEntitiesParamsWithHTTPClient creates a new GetVersionsPrebuiltEntitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVersionsPrebuiltEntitiesParamsWithHTTPClient(client *http.Client) *GetVersionsPrebuiltEntitiesParams {
	var ()
	return &GetVersionsPrebuiltEntitiesParams{
		HTTPClient: client,
	}
}

/*GetVersionsPrebuiltEntitiesParams contains all the parameters to send to the API endpoint
for the get versions prebuilt entities operation typically these are written to a http.Request
*/
type GetVersionsPrebuiltEntitiesParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) WithTimeout(timeout time.Duration) *GetVersionsPrebuiltEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) WithContext(ctx context.Context) *GetVersionsPrebuiltEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) WithHTTPClient(client *http.Client) *GetVersionsPrebuiltEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) WithAppID(appID string) *GetVersionsPrebuiltEntitiesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithVersionID adds the versionID to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) WithVersionID(versionID string) *GetVersionsPrebuiltEntitiesParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get versions prebuilt entities params
func (o *GetVersionsPrebuiltEntitiesParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionsPrebuiltEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
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
