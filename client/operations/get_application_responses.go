// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/crazedpeanut/luis/models"
)

// GetApplicationReader is a Reader for the GetApplication structure.
type GetApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetApplicationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationOK creates a GetApplicationOK with default headers values
func NewGetApplicationOK() *GetApplicationOK {
	return &GetApplicationOK{}
}

/*GetApplicationOK handles this case with default header values.

A JSON object containing the application info.
*/
type GetApplicationOK struct {
	Payload *models.ApplicationGetObject
}

func (o *GetApplicationOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}][%d] getApplicationOK  %+v", 200, o.Payload)
}

func (o *GetApplicationOK) GetPayload() *models.ApplicationGetObject {
	return o.Payload
}

func (o *GetApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationGetObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationBadRequest creates a GetApplicationBadRequest with default headers values
func NewGetApplicationBadRequest() *GetApplicationBadRequest {
	return &GetApplicationBadRequest{}
}

/*GetApplicationBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetApplicationBadRequest struct {
}

func (o *GetApplicationBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}][%d] getApplicationBadRequest ", 400)
}

func (o *GetApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationUnauthorized creates a GetApplicationUnauthorized with default headers values
func NewGetApplicationUnauthorized() *GetApplicationUnauthorized {
	return &GetApplicationUnauthorized{}
}

/*GetApplicationUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type GetApplicationUnauthorized struct {
}

func (o *GetApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}][%d] getApplicationUnauthorized ", 401)
}

func (o *GetApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationForbidden creates a GetApplicationForbidden with default headers values
func NewGetApplicationForbidden() *GetApplicationForbidden {
	return &GetApplicationForbidden{}
}

/*GetApplicationForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetApplicationForbidden struct {
}

func (o *GetApplicationForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}][%d] getApplicationForbidden ", 403)
}

func (o *GetApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationTooManyRequests creates a GetApplicationTooManyRequests with default headers values
func NewGetApplicationTooManyRequests() *GetApplicationTooManyRequests {
	return &GetApplicationTooManyRequests{}
}

/*GetApplicationTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetApplicationTooManyRequests struct {
}

func (o *GetApplicationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}][%d] getApplicationTooManyRequests ", 429)
}

func (o *GetApplicationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
