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

// NewGetPatternFeaturesParams creates a new GetPatternFeaturesParams object
// with the default values initialized.
func NewGetPatternFeaturesParams() *GetPatternFeaturesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPatternFeaturesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatternFeaturesParamsWithTimeout creates a new GetPatternFeaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatternFeaturesParamsWithTimeout(timeout time.Duration) *GetPatternFeaturesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPatternFeaturesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: timeout,
	}
}

// NewGetPatternFeaturesParamsWithContext creates a new GetPatternFeaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatternFeaturesParamsWithContext(ctx context.Context) *GetPatternFeaturesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPatternFeaturesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		Context: ctx,
	}
}

// NewGetPatternFeaturesParamsWithHTTPClient creates a new GetPatternFeaturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatternFeaturesParamsWithHTTPClient(client *http.Client) *GetPatternFeaturesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetPatternFeaturesParams{
		Skip:       &skipDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*GetPatternFeaturesParams contains all the parameters to send to the API endpoint
for the get pattern features operation typically these are written to a http.Request
*/
type GetPatternFeaturesParams struct {

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
	  The version ID of the task.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pattern features params
func (o *GetPatternFeaturesParams) WithTimeout(timeout time.Duration) *GetPatternFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pattern features params
func (o *GetPatternFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pattern features params
func (o *GetPatternFeaturesParams) WithContext(ctx context.Context) *GetPatternFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pattern features params
func (o *GetPatternFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pattern features params
func (o *GetPatternFeaturesParams) WithHTTPClient(client *http.Client) *GetPatternFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pattern features params
func (o *GetPatternFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get pattern features params
func (o *GetPatternFeaturesParams) WithAppID(appID string) *GetPatternFeaturesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get pattern features params
func (o *GetPatternFeaturesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithSkip adds the skip to the get pattern features params
func (o *GetPatternFeaturesParams) WithSkip(skip *int64) *GetPatternFeaturesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get pattern features params
func (o *GetPatternFeaturesParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithTake adds the take to the get pattern features params
func (o *GetPatternFeaturesParams) WithTake(take *int64) *GetPatternFeaturesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the get pattern features params
func (o *GetPatternFeaturesParams) SetTake(take *int64) {
	o.Take = take
}

// WithVersionID adds the versionID to the get pattern features params
func (o *GetPatternFeaturesParams) WithVersionID(versionID string) *GetPatternFeaturesParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get pattern features params
func (o *GetPatternFeaturesParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatternFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
