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

// NewGetVersionsParams creates a new GetVersionsParams object
// with the default values initialized.
func NewGetVersionsParams() *GetVersionsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetVersionsParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionsParamsWithTimeout creates a new GetVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersionsParamsWithTimeout(timeout time.Duration) *GetVersionsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetVersionsParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: timeout,
	}
}

// NewGetVersionsParamsWithContext creates a new GetVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersionsParamsWithContext(ctx context.Context) *GetVersionsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetVersionsParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		Context: ctx,
	}
}

// NewGetVersionsParamsWithHTTPClient creates a new GetVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVersionsParamsWithHTTPClient(client *http.Client) *GetVersionsParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetVersionsParams{
		Skip:       &skipDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*GetVersionsParams contains all the parameters to send to the API endpoint
for the get versions operation typically these are written to a http.Request
*/
type GetVersionsParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get versions params
func (o *GetVersionsParams) WithTimeout(timeout time.Duration) *GetVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versions params
func (o *GetVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versions params
func (o *GetVersionsParams) WithContext(ctx context.Context) *GetVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versions params
func (o *GetVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versions params
func (o *GetVersionsParams) WithHTTPClient(client *http.Client) *GetVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versions params
func (o *GetVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get versions params
func (o *GetVersionsParams) WithAppID(appID string) *GetVersionsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get versions params
func (o *GetVersionsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithSkip adds the skip to the get versions params
func (o *GetVersionsParams) WithSkip(skip *int64) *GetVersionsParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get versions params
func (o *GetVersionsParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithTake adds the take to the get versions params
func (o *GetVersionsParams) WithTake(take *int64) *GetVersionsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the get versions params
func (o *GetVersionsParams) SetTake(take *int64) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
