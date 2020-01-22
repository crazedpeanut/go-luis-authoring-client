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

// NewPublishApplicationParams creates a new PublishApplicationParams object
// with the default values initialized.
func NewPublishApplicationParams() *PublishApplicationParams {
	var ()
	return &PublishApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublishApplicationParamsWithTimeout creates a new PublishApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublishApplicationParamsWithTimeout(timeout time.Duration) *PublishApplicationParams {
	var ()
	return &PublishApplicationParams{

		timeout: timeout,
	}
}

// NewPublishApplicationParamsWithContext creates a new PublishApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublishApplicationParamsWithContext(ctx context.Context) *PublishApplicationParams {
	var ()
	return &PublishApplicationParams{

		Context: ctx,
	}
}

// NewPublishApplicationParamsWithHTTPClient creates a new PublishApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublishApplicationParamsWithHTTPClient(client *http.Client) *PublishApplicationParams {
	var ()
	return &PublishApplicationParams{
		HTTPClient: client,
	}
}

/*PublishApplicationParams contains all the parameters to send to the API endpoint
for the publish application operation typically these are written to a http.Request
*/
type PublishApplicationParams struct {

	/*AppID
	  Format - guid. The application version ID.

	*/
	AppID string
	/*ApplicationPublishObject
	  A JSON object containing the publish Properties.<br><br>

	The flag "isStaging" should be set to true in case you want to publish to the STAGING slot, otherwise the application version will be published to the PRODUCTION slot.
	In case you do not want to publish to either the PRODUCTION or STAGING slots, you can set the flag "directVersionPublish" to true and query the endpoint [directly using the versionId](https://westus.dev.cognitive.microsoft.com/docs/services/luis-endpoint-api-v3-0-preview/operations/5cb0a9459a1fe8fa44c28dd8).

	The [publish location](https://docs.microsoft.com/azure/cognitive-services/LUIS/luis-reference-regions#publishing-regions) is determined from the creation location.

	*/
	ApplicationPublishObject *models.ApplicationPublishObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the publish application params
func (o *PublishApplicationParams) WithTimeout(timeout time.Duration) *PublishApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publish application params
func (o *PublishApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publish application params
func (o *PublishApplicationParams) WithContext(ctx context.Context) *PublishApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publish application params
func (o *PublishApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publish application params
func (o *PublishApplicationParams) WithHTTPClient(client *http.Client) *PublishApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publish application params
func (o *PublishApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the publish application params
func (o *PublishApplicationParams) WithAppID(appID string) *PublishApplicationParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the publish application params
func (o *PublishApplicationParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithApplicationPublishObject adds the applicationPublishObject to the publish application params
func (o *PublishApplicationParams) WithApplicationPublishObject(applicationPublishObject *models.ApplicationPublishObject) *PublishApplicationParams {
	o.SetApplicationPublishObject(applicationPublishObject)
	return o
}

// SetApplicationPublishObject adds the applicationPublishObject to the publish application params
func (o *PublishApplicationParams) SetApplicationPublishObject(applicationPublishObject *models.ApplicationPublishObject) {
	o.ApplicationPublishObject = applicationPublishObject
}

// WriteToRequest writes these params to a swagger request
func (o *PublishApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.ApplicationPublishObject != nil {
		if err := r.SetBodyParam(o.ApplicationPublishObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
