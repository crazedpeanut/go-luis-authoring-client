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

// ImportVersionJSONReader is a Reader for the ImportVersionJSON structure.
type ImportVersionJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportVersionJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImportVersionJSONCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportVersionJSONBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportVersionJSONUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportVersionJSONForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewImportVersionJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportVersionJSONCreated creates a ImportVersionJSONCreated with default headers values
func NewImportVersionJSONCreated() *ImportVersionJSONCreated {
	return &ImportVersionJSONCreated{}
}

/*ImportVersionJSONCreated handles this case with default header values.

The created application version.
*/
type ImportVersionJSONCreated struct {
}

func (o *ImportVersionJSONCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import/][%d] importVersionJsonCreated ", 201)
}

func (o *ImportVersionJSONCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportVersionJSONBadRequest creates a ImportVersionJSONBadRequest with default headers values
func NewImportVersionJSONBadRequest() *ImportVersionJSONBadRequest {
	return &ImportVersionJSONBadRequest{}
}

/*ImportVersionJSONBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type ImportVersionJSONBadRequest struct {
	Payload *models.ErrorResponseObject
}

func (o *ImportVersionJSONBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import/][%d] importVersionJsonBadRequest  %+v", 400, o.Payload)
}

func (o *ImportVersionJSONBadRequest) GetPayload() *models.ErrorResponseObject {
	return o.Payload
}

func (o *ImportVersionJSONBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportVersionJSONUnauthorized creates a ImportVersionJSONUnauthorized with default headers values
func NewImportVersionJSONUnauthorized() *ImportVersionJSONUnauthorized {
	return &ImportVersionJSONUnauthorized{}
}

/*ImportVersionJSONUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type ImportVersionJSONUnauthorized struct {
}

func (o *ImportVersionJSONUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import/][%d] importVersionJsonUnauthorized ", 401)
}

func (o *ImportVersionJSONUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportVersionJSONForbidden creates a ImportVersionJSONForbidden with default headers values
func NewImportVersionJSONForbidden() *ImportVersionJSONForbidden {
	return &ImportVersionJSONForbidden{}
}

/*ImportVersionJSONForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type ImportVersionJSONForbidden struct {
}

func (o *ImportVersionJSONForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import/][%d] importVersionJsonForbidden ", 403)
}

func (o *ImportVersionJSONForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportVersionJSONTooManyRequests creates a ImportVersionJSONTooManyRequests with default headers values
func NewImportVersionJSONTooManyRequests() *ImportVersionJSONTooManyRequests {
	return &ImportVersionJSONTooManyRequests{}
}

/*ImportVersionJSONTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type ImportVersionJSONTooManyRequests struct {
}

func (o *ImportVersionJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/import/][%d] importVersionJsonTooManyRequests ", 429)
}

func (o *ImportVersionJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
