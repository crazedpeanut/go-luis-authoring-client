// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade5b5dd5b81c209ce2e5b4Reader is a Reader for the Nr5ade5b5dd5b81c209ce2e5b4 structure.
type Nr5ade5b5dd5b81c209ce2e5b4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade5b5dd5b81c209ce2e5b4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ade5b5dd5b81c209ce2e5b4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade5b5dd5b81c209ce2e5b4BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade5b5dd5b81c209ce2e5b4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade5b5dd5b81c209ce2e5b4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade5b5dd5b81c209ce2e5b4TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade5b5dd5b81c209ce2e5b4OK creates a Nr5ade5b5dd5b81c209ce2e5b4OK with default headers values
func NewNr5ade5b5dd5b81c209ce2e5b4OK() *Nr5ade5b5dd5b81c209ce2e5b4OK {
	return &Nr5ade5b5dd5b81c209ce2e5b4OK{}
}

/*Nr5ade5b5dd5b81c209ce2e5b4OK handles this case with default header values.

Nr5ade5b5dd5b81c209ce2e5b4OK 5ade5b5dd5b81c209ce2e5b4 o k
*/
type Nr5ade5b5dd5b81c209ce2e5b4OK struct {
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4OK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/closedlists/{entityId}/roles/{roleId}][%d] 5ade5b5dd5b81c209ce2e5b4OK ", 200)
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5b5dd5b81c209ce2e5b4BadRequest creates a Nr5ade5b5dd5b81c209ce2e5b4BadRequest with default headers values
func NewNr5ade5b5dd5b81c209ce2e5b4BadRequest() *Nr5ade5b5dd5b81c209ce2e5b4BadRequest {
	return &Nr5ade5b5dd5b81c209ce2e5b4BadRequest{}
}

/*Nr5ade5b5dd5b81c209ce2e5b4BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade5b5dd5b81c209ce2e5b4BadRequest struct {
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4BadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/closedlists/{entityId}/roles/{roleId}][%d] 5ade5b5dd5b81c209ce2e5b4BadRequest ", 400)
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5b5dd5b81c209ce2e5b4Unauthorized creates a Nr5ade5b5dd5b81c209ce2e5b4Unauthorized with default headers values
func NewNr5ade5b5dd5b81c209ce2e5b4Unauthorized() *Nr5ade5b5dd5b81c209ce2e5b4Unauthorized {
	return &Nr5ade5b5dd5b81c209ce2e5b4Unauthorized{}
}

/*Nr5ade5b5dd5b81c209ce2e5b4Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade5b5dd5b81c209ce2e5b4Unauthorized struct {
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/closedlists/{entityId}/roles/{roleId}][%d] 5ade5b5dd5b81c209ce2e5b4Unauthorized ", 401)
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5b5dd5b81c209ce2e5b4Forbidden creates a Nr5ade5b5dd5b81c209ce2e5b4Forbidden with default headers values
func NewNr5ade5b5dd5b81c209ce2e5b4Forbidden() *Nr5ade5b5dd5b81c209ce2e5b4Forbidden {
	return &Nr5ade5b5dd5b81c209ce2e5b4Forbidden{}
}

/*Nr5ade5b5dd5b81c209ce2e5b4Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade5b5dd5b81c209ce2e5b4Forbidden struct {
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4Forbidden) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/closedlists/{entityId}/roles/{roleId}][%d] 5ade5b5dd5b81c209ce2e5b4Forbidden ", 403)
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5b5dd5b81c209ce2e5b4TooManyRequests creates a Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests with default headers values
func NewNr5ade5b5dd5b81c209ce2e5b4TooManyRequests() *Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests {
	return &Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests{}
}

/*Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests struct {
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/closedlists/{entityId}/roles/{roleId}][%d] 5ade5b5dd5b81c209ce2e5b4TooManyRequests ", 429)
}

func (o *Nr5ade5b5dd5b81c209ce2e5b4TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}