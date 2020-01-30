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

// NewGetPrebuildEntitiesParams creates a new GetPrebuildEntitiesParams object
// with the default values initialized.
func NewGetPrebuildEntitiesParams() *GetPrebuildEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPrebuildEntitiesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrebuildEntitiesParamsWithTimeout creates a new GetPrebuildEntitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrebuildEntitiesParamsWithTimeout(timeout time.Duration) *GetPrebuildEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPrebuildEntitiesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: timeout,
	}
}

// NewGetPrebuildEntitiesParamsWithContext creates a new GetPrebuildEntitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrebuildEntitiesParamsWithContext(ctx context.Context) *GetPrebuildEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPrebuildEntitiesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		Context: ctx,
	}
}

// NewGetPrebuildEntitiesParamsWithHTTPClient creates a new GetPrebuildEntitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrebuildEntitiesParamsWithHTTPClient(client *http.Client) *GetPrebuildEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPrebuildEntitiesParams{
		Skip:       &skipDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*GetPrebuildEntitiesParams contains all the parameters to send to the API endpoint
for the get prebuild entities operation typically these are written to a http.Request
*/
type GetPrebuildEntitiesParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Skip
	  The number of entries to skip. Default value is 0.

	*/
	Skip *int64
	/*Take
	  The number of entries to return. Maximum page size is 500. Default is 100.

	*/
	Take *int64
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithTimeout(timeout time.Duration) *GetPrebuildEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithContext(ctx context.Context) *GetPrebuildEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithHTTPClient(client *http.Client) *GetPrebuildEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithAppID(appID string) *GetPrebuildEntitiesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithSkip adds the skip to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithSkip(skip *int64) *GetPrebuildEntitiesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithTake adds the take to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithTake(take *int64) *GetPrebuildEntitiesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetTake(take *int64) {
	o.Take = take
}

// WithVersionID adds the versionID to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) WithVersionID(versionID string) *GetPrebuildEntitiesParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get prebuild entities params
func (o *GetPrebuildEntitiesParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrebuildEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int64
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt64(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int64
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt64(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
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