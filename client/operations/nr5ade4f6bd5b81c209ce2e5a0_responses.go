// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade4f6bd5b81c209ce2e5a0Reader is a Reader for the Nr5ade4f6bd5b81c209ce2e5a0 structure.
type Nr5ade4f6bd5b81c209ce2e5a0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade4f6bd5b81c209ce2e5a0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ade4f6bd5b81c209ce2e5a0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade4f6bd5b81c209ce2e5a0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade4f6bd5b81c209ce2e5a0Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade4f6bd5b81c209ce2e5a0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade4f6bd5b81c209ce2e5a0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade4f6bd5b81c209ce2e5a0OK creates a Nr5ade4f6bd5b81c209ce2e5a0OK with default headers values
func NewNr5ade4f6bd5b81c209ce2e5a0OK() *Nr5ade4f6bd5b81c209ce2e5a0OK {
	return &Nr5ade4f6bd5b81c209ce2e5a0OK{}
}

/*Nr5ade4f6bd5b81c209ce2e5a0OK handles this case with default header values.

A list of Pattern.Any entity model infos.
*/
type Nr5ade4f6bd5b81c209ce2e5a0OK struct {
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0OK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities][%d] 5ade4f6bd5b81c209ce2e5a0OK ", 200)
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade4f6bd5b81c209ce2e5a0BadRequest creates a Nr5ade4f6bd5b81c209ce2e5a0BadRequest with default headers values
func NewNr5ade4f6bd5b81c209ce2e5a0BadRequest() *Nr5ade4f6bd5b81c209ce2e5a0BadRequest {
	return &Nr5ade4f6bd5b81c209ce2e5a0BadRequest{}
}

/*Nr5ade4f6bd5b81c209ce2e5a0BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade4f6bd5b81c209ce2e5a0BadRequest struct {
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0BadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities][%d] 5ade4f6bd5b81c209ce2e5a0BadRequest ", 400)
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade4f6bd5b81c209ce2e5a0Unauthorized creates a Nr5ade4f6bd5b81c209ce2e5a0Unauthorized with default headers values
func NewNr5ade4f6bd5b81c209ce2e5a0Unauthorized() *Nr5ade4f6bd5b81c209ce2e5a0Unauthorized {
	return &Nr5ade4f6bd5b81c209ce2e5a0Unauthorized{}
}

/*Nr5ade4f6bd5b81c209ce2e5a0Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade4f6bd5b81c209ce2e5a0Unauthorized struct {
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0Unauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities][%d] 5ade4f6bd5b81c209ce2e5a0Unauthorized ", 401)
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade4f6bd5b81c209ce2e5a0Forbidden creates a Nr5ade4f6bd5b81c209ce2e5a0Forbidden with default headers values
func NewNr5ade4f6bd5b81c209ce2e5a0Forbidden() *Nr5ade4f6bd5b81c209ce2e5a0Forbidden {
	return &Nr5ade4f6bd5b81c209ce2e5a0Forbidden{}
}

/*Nr5ade4f6bd5b81c209ce2e5a0Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade4f6bd5b81c209ce2e5a0Forbidden struct {
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0Forbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities][%d] 5ade4f6bd5b81c209ce2e5a0Forbidden ", 403)
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade4f6bd5b81c209ce2e5a0TooManyRequests creates a Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests with default headers values
func NewNr5ade4f6bd5b81c209ce2e5a0TooManyRequests() *Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests {
	return &Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests{}
}

/*Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests struct {
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities][%d] 5ade4f6bd5b81c209ce2e5a0TooManyRequests ", 429)
}

func (o *Nr5ade4f6bd5b81c209ce2e5a0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}