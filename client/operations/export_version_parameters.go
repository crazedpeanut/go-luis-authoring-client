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

// NewExportVersionParams creates a new ExportVersionParams object
// with the default values initialized.
func NewExportVersionParams() *ExportVersionParams {
	var (
		formatDefault = string("json")
	)
	return &ExportVersionParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewExportVersionParamsWithTimeout creates a new ExportVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportVersionParamsWithTimeout(timeout time.Duration) *ExportVersionParams {
	var (
		formatDefault = string("json")
	)
	return &ExportVersionParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewExportVersionParamsWithContext creates a new ExportVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportVersionParamsWithContext(ctx context.Context) *ExportVersionParams {
	var (
		formatDefault = string("json")
	)
	return &ExportVersionParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewExportVersionParamsWithHTTPClient creates a new ExportVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportVersionParamsWithHTTPClient(client *http.Client) *ExportVersionParams {
	var (
		formatDefault = string("json")
	)
	return &ExportVersionParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*ExportVersionParams contains all the parameters to send to the API endpoint
for the export version operation typically these are written to a http.Request
*/
type ExportVersionParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Format
	  Export format type.

	*/
	Format *string
	/*VersionID
	  The application version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export version params
func (o *ExportVersionParams) WithTimeout(timeout time.Duration) *ExportVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export version params
func (o *ExportVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export version params
func (o *ExportVersionParams) WithContext(ctx context.Context) *ExportVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export version params
func (o *ExportVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export version params
func (o *ExportVersionParams) WithHTTPClient(client *http.Client) *ExportVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export version params
func (o *ExportVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the export version params
func (o *ExportVersionParams) WithAppID(appID string) *ExportVersionParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the export version params
func (o *ExportVersionParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithFormat adds the format to the export version params
func (o *ExportVersionParams) WithFormat(format *string) *ExportVersionParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the export version params
func (o *ExportVersionParams) SetFormat(format *string) {
	o.Format = format
}

// WithVersionID adds the versionID to the export version params
func (o *ExportVersionParams) WithVersionID(versionID string) *ExportVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the export version params
func (o *ExportVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
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
