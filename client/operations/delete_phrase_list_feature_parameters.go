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

// NewDeletePhraseListFeatureParams creates a new DeletePhraseListFeatureParams object
// with the default values initialized.
func NewDeletePhraseListFeatureParams() *DeletePhraseListFeatureParams {
	var ()
	return &DeletePhraseListFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePhraseListFeatureParamsWithTimeout creates a new DeletePhraseListFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePhraseListFeatureParamsWithTimeout(timeout time.Duration) *DeletePhraseListFeatureParams {
	var ()
	return &DeletePhraseListFeatureParams{

		timeout: timeout,
	}
}

// NewDeletePhraseListFeatureParamsWithContext creates a new DeletePhraseListFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePhraseListFeatureParamsWithContext(ctx context.Context) *DeletePhraseListFeatureParams {
	var ()
	return &DeletePhraseListFeatureParams{

		Context: ctx,
	}
}

// NewDeletePhraseListFeatureParamsWithHTTPClient creates a new DeletePhraseListFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePhraseListFeatureParamsWithHTTPClient(client *http.Client) *DeletePhraseListFeatureParams {
	var ()
	return &DeletePhraseListFeatureParams{
		HTTPClient: client,
	}
}

/*DeletePhraseListFeatureParams contains all the parameters to send to the API endpoint
for the delete phrase list feature operation typically these are written to a http.Request
*/
type DeletePhraseListFeatureParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*PhraselistID
	  The ID of the feature to be deleted.

	*/
	PhraselistID int64
	/*VersionID
	  The version ID of the task.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) WithTimeout(timeout time.Duration) *DeletePhraseListFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) WithContext(ctx context.Context) *DeletePhraseListFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) WithHTTPClient(client *http.Client) *DeletePhraseListFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) WithAppID(appID string) *DeletePhraseListFeatureParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithPhraselistID adds the phraselistID to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) WithPhraselistID(phraselistID int64) *DeletePhraseListFeatureParams {
	o.SetPhraselistID(phraselistID)
	return o
}

// SetPhraselistID adds the phraselistId to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) SetPhraselistID(phraselistID int64) {
	o.PhraselistID = phraselistID
}

// WithVersionID adds the versionID to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) WithVersionID(versionID string) *DeletePhraseListFeatureParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete phrase list feature params
func (o *DeletePhraseListFeatureParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePhraseListFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param phraselistId
	if err := r.SetPathParam("phraselistId", swag.FormatInt64(o.PhraselistID)); err != nil {
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
