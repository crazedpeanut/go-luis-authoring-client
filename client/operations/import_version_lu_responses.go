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

// ImportVersionLuReader is a Reader for the ImportVersionLu structure.
type ImportVersionLuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportVersionLuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImportVersionLuCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportVersionLuBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportVersionLuUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportVersionLuForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewImportVersionLuTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportVersionLuCreated creates a ImportVersionLuCreated with default headers values
func NewImportVersionLuCreated() *ImportVersionLuCreated {
	return &ImportVersionLuCreated{}
}

/*ImportVersionLuCreated handles this case with default header values.

The created application version.
*/
type ImportVersionLuCreated struct {
}

func (o *ImportVersionLuCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import][%d] importVersionLuCreated ", 201)
}

func (o *ImportVersionLuCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportVersionLuBadRequest creates a ImportVersionLuBadRequest with default headers values
func NewImportVersionLuBadRequest() *ImportVersionLuBadRequest {
	return &ImportVersionLuBadRequest{}
}

/*ImportVersionLuBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type ImportVersionLuBadRequest struct {
	Payload *models.ErrorResponseObject
}

func (o *ImportVersionLuBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import][%d] importVersionLuBadRequest  %+v", 400, o.Payload)
}

func (o *ImportVersionLuBadRequest) GetPayload() *models.ErrorResponseObject {
	return o.Payload
}

func (o *ImportVersionLuBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportVersionLuUnauthorized creates a ImportVersionLuUnauthorized with default headers values
func NewImportVersionLuUnauthorized() *ImportVersionLuUnauthorized {
	return &ImportVersionLuUnauthorized{}
}

/*ImportVersionLuUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type ImportVersionLuUnauthorized struct {
}

func (o *ImportVersionLuUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import][%d] importVersionLuUnauthorized ", 401)
}

func (o *ImportVersionLuUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportVersionLuForbidden creates a ImportVersionLuForbidden with default headers values
func NewImportVersionLuForbidden() *ImportVersionLuForbidden {
	return &ImportVersionLuForbidden{}
}

/*ImportVersionLuForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type ImportVersionLuForbidden struct {
}

func (o *ImportVersionLuForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import][%d] importVersionLuForbidden ", 403)
}

func (o *ImportVersionLuForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportVersionLuTooManyRequests creates a ImportVersionLuTooManyRequests with default headers values
func NewImportVersionLuTooManyRequests() *ImportVersionLuTooManyRequests {
	return &ImportVersionLuTooManyRequests{}
}

/*ImportVersionLuTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type ImportVersionLuTooManyRequests struct {
}

func (o *ImportVersionLuTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import][%d] importVersionLuTooManyRequests ", 429)
}

func (o *ImportVersionLuTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
