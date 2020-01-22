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

// NewCreateApplicationParams creates a new CreateApplicationParams object
// with the default values initialized.
func NewCreateApplicationParams() *CreateApplicationParams {
	var ()
	return &CreateApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApplicationParamsWithTimeout creates a new CreateApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateApplicationParamsWithTimeout(timeout time.Duration) *CreateApplicationParams {
	var ()
	return &CreateApplicationParams{

		timeout: timeout,
	}
}

// NewCreateApplicationParamsWithContext creates a new CreateApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateApplicationParamsWithContext(ctx context.Context) *CreateApplicationParams {
	var ()
	return &CreateApplicationParams{

		Context: ctx,
	}
}

// NewCreateApplicationParamsWithHTTPClient creates a new CreateApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateApplicationParamsWithHTTPClient(client *http.Client) *CreateApplicationParams {
	var ()
	return &CreateApplicationParams{
		HTTPClient: client,
	}
}

/*CreateApplicationParams contains all the parameters to send to the API endpoint
for the create application operation typically these are written to a http.Request
*/
type CreateApplicationParams struct {

	/*ApplicationCreateObject
	  Application name length has to be less than 50 characters.

	The culture and tokenizer version of an app cannot be changed after the app is created.

	The default version is "0.1"

	*/
	ApplicationCreateObject *models.ApplicationCreateObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create application params
func (o *CreateApplicationParams) WithTimeout(timeout time.Duration) *CreateApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create application params
func (o *CreateApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create application params
func (o *CreateApplicationParams) WithContext(ctx context.Context) *CreateApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create application params
func (o *CreateApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create application params
func (o *CreateApplicationParams) WithHTTPClient(client *http.Client) *CreateApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create application params
func (o *CreateApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCreateObject adds the applicationCreateObject to the create application params
func (o *CreateApplicationParams) WithApplicationCreateObject(applicationCreateObject *models.ApplicationCreateObject) *CreateApplicationParams {
	o.SetApplicationCreateObject(applicationCreateObject)
	return o
}

// SetApplicationCreateObject adds the applicationCreateObject to the create application params
func (o *CreateApplicationParams) SetApplicationCreateObject(applicationCreateObject *models.ApplicationCreateObject) {
	o.ApplicationCreateObject = applicationCreateObject
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationCreateObject != nil {
		if err := r.SetBodyParam(o.ApplicationCreateObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
