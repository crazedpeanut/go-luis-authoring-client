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

// NewGetIntentsParams creates a new GetIntentsParams object
// with the default values initialized.
func NewGetIntentsParams() *GetIntentsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetIntentsParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntentsParamsWithTimeout creates a new GetIntentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntentsParamsWithTimeout(timeout time.Duration) *GetIntentsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetIntentsParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: timeout,
	}
}

// NewGetIntentsParamsWithContext creates a new GetIntentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntentsParamsWithContext(ctx context.Context) *GetIntentsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetIntentsParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		Context: ctx,
	}
}

// NewGetIntentsParamsWithHTTPClient creates a new GetIntentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntentsParamsWithHTTPClient(client *http.Client) *GetIntentsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetIntentsParams{
		Skip:       &skipDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*GetIntentsParams contains all the parameters to send to the API endpoint
for the get intents operation typically these are written to a http.Request
*/
type GetIntentsParams struct {

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

// WithTimeout adds the timeout to the get intents params
func (o *GetIntentsParams) WithTimeout(timeout time.Duration) *GetIntentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get intents params
func (o *GetIntentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get intents params
func (o *GetIntentsParams) WithContext(ctx context.Context) *GetIntentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get intents params
func (o *GetIntentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get intents params
func (o *GetIntentsParams) WithHTTPClient(client *http.Client) *GetIntentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get intents params
func (o *GetIntentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get intents params
func (o *GetIntentsParams) WithAppID(appID string) *GetIntentsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get intents params
func (o *GetIntentsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithSkip adds the skip to the get intents params
func (o *GetIntentsParams) WithSkip(skip *int64) *GetIntentsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get intents params
func (o *GetIntentsParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithTake adds the take to the get intents params
func (o *GetIntentsParams) WithTake(take *int64) *GetIntentsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the get intents params
func (o *GetIntentsParams) SetTake(take *int64) {
	o.Take = take
}

// WithVersionID adds the versionID to the get intents params
func (o *GetIntentsParams) WithVersionID(versionID string) *GetIntentsParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get intents params
func (o *GetIntentsParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
