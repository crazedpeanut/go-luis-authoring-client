// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade5137d5b81c209ce2e5a2Reader is a Reader for the Nr5ade5137d5b81c209ce2e5a2 structure.
type Nr5ade5137d5b81c209ce2e5a2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade5137d5b81c209ce2e5a2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ade5137d5b81c209ce2e5a2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade5137d5b81c209ce2e5a2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade5137d5b81c209ce2e5a2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade5137d5b81c209ce2e5a2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade5137d5b81c209ce2e5a2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade5137d5b81c209ce2e5a2OK creates a Nr5ade5137d5b81c209ce2e5a2OK with default headers values
func NewNr5ade5137d5b81c209ce2e5a2OK() *Nr5ade5137d5b81c209ce2e5a2OK {
	return &Nr5ade5137d5b81c209ce2e5a2OK{}
}

/*Nr5ade5137d5b81c209ce2e5a2OK handles this case with default header values.

A list of the roles.
*/
type Nr5ade5137d5b81c209ce2e5a2OK struct {
}

func (o *Nr5ade5137d5b81c209ce2e5a2OK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles][%d] 5ade5137d5b81c209ce2e5a2OK ", 200)
}

func (o *Nr5ade5137d5b81c209ce2e5a2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5137d5b81c209ce2e5a2BadRequest creates a Nr5ade5137d5b81c209ce2e5a2BadRequest with default headers values
func NewNr5ade5137d5b81c209ce2e5a2BadRequest() *Nr5ade5137d5b81c209ce2e5a2BadRequest {
	return &Nr5ade5137d5b81c209ce2e5a2BadRequest{}
}

/*Nr5ade5137d5b81c209ce2e5a2BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade5137d5b81c209ce2e5a2BadRequest struct {
}

func (o *Nr5ade5137d5b81c209ce2e5a2BadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles][%d] 5ade5137d5b81c209ce2e5a2BadRequest ", 400)
}

func (o *Nr5ade5137d5b81c209ce2e5a2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5137d5b81c209ce2e5a2Unauthorized creates a Nr5ade5137d5b81c209ce2e5a2Unauthorized with default headers values
func NewNr5ade5137d5b81c209ce2e5a2Unauthorized() *Nr5ade5137d5b81c209ce2e5a2Unauthorized {
	return &Nr5ade5137d5b81c209ce2e5a2Unauthorized{}
}

/*Nr5ade5137d5b81c209ce2e5a2Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade5137d5b81c209ce2e5a2Unauthorized struct {
}

func (o *Nr5ade5137d5b81c209ce2e5a2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles][%d] 5ade5137d5b81c209ce2e5a2Unauthorized ", 401)
}

func (o *Nr5ade5137d5b81c209ce2e5a2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5137d5b81c209ce2e5a2Forbidden creates a Nr5ade5137d5b81c209ce2e5a2Forbidden with default headers values
func NewNr5ade5137d5b81c209ce2e5a2Forbidden() *Nr5ade5137d5b81c209ce2e5a2Forbidden {
	return &Nr5ade5137d5b81c209ce2e5a2Forbidden{}
}

/*Nr5ade5137d5b81c209ce2e5a2Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade5137d5b81c209ce2e5a2Forbidden struct {
}

func (o *Nr5ade5137d5b81c209ce2e5a2Forbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles][%d] 5ade5137d5b81c209ce2e5a2Forbidden ", 403)
}

func (o *Nr5ade5137d5b81c209ce2e5a2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5137d5b81c209ce2e5a2TooManyRequests creates a Nr5ade5137d5b81c209ce2e5a2TooManyRequests with default headers values
func NewNr5ade5137d5b81c209ce2e5a2TooManyRequests() *Nr5ade5137d5b81c209ce2e5a2TooManyRequests {
	return &Nr5ade5137d5b81c209ce2e5a2TooManyRequests{}
}

/*Nr5ade5137d5b81c209ce2e5a2TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade5137d5b81c209ce2e5a2TooManyRequests struct {
}

func (o *Nr5ade5137d5b81c209ce2e5a2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles][%d] 5ade5137d5b81c209ce2e5a2TooManyRequests ", 429)
}

func (o *Nr5ade5137d5b81c209ce2e5a2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}