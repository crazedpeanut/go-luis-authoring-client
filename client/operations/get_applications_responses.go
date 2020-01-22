// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/crazedpeanut/go-luis-authoring-client/models"
)

// GetApplicationsReader is a Reader for the GetApplications structure.
type GetApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetApplicationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationsOK creates a GetApplicationsOK with default headers values
func NewGetApplicationsOK() *GetApplicationsOK {
	return &GetApplicationsOK{}
}

/*GetApplicationsOK handles this case with default header values.

A list of the user applications.
*/
type GetApplicationsOK struct {
	Payload []*models.ApplicationGetObject
}

func (o *GetApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /apps/][%d] getApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsOK) GetPayload() []*models.ApplicationGetObject {
	return o.Payload
}

func (o *GetApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsBadRequest creates a GetApplicationsBadRequest with default headers values
func NewGetApplicationsBadRequest() *GetApplicationsBadRequest {
	return &GetApplicationsBadRequest{}
}

/*GetApplicationsBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetApplicationsBadRequest struct {
}

func (o *GetApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/][%d] getApplicationsBadRequest ", 400)
}

func (o *GetApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationsUnauthorized creates a GetApplicationsUnauthorized with default headers values
func NewGetApplicationsUnauthorized() *GetApplicationsUnauthorized {
	return &GetApplicationsUnauthorized{}
}

/*GetApplicationsUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type GetApplicationsUnauthorized struct {
}

func (o *GetApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/][%d] getApplicationsUnauthorized ", 401)
}

func (o *GetApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationsForbidden creates a GetApplicationsForbidden with default headers values
func NewGetApplicationsForbidden() *GetApplicationsForbidden {
	return &GetApplicationsForbidden{}
}

/*GetApplicationsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetApplicationsForbidden struct {
}

func (o *GetApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/][%d] getApplicationsForbidden ", 403)
}

func (o *GetApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationsTooManyRequests creates a GetApplicationsTooManyRequests with default headers values
func NewGetApplicationsTooManyRequests() *GetApplicationsTooManyRequests {
	return &GetApplicationsTooManyRequests{}
}

/*GetApplicationsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetApplicationsTooManyRequests struct {
}

func (o *GetApplicationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/][%d] getApplicationsTooManyRequests ", 429)
}

func (o *GetApplicationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
