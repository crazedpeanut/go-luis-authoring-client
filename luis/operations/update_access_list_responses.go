// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateAccessListReader is a Reader for the UpdateAccessList structure.
type UpdateAccessListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccessListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccessListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAccessListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAccessListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAccessListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateAccessListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAccessListOK creates a UpdateAccessListOK with default headers values
func NewUpdateAccessListOK() *UpdateAccessListOK {
	return &UpdateAccessListOK{}
}

/*UpdateAccessListOK handles this case with default header values.

UpdateAccessListOK update access list o k
*/
type UpdateAccessListOK struct {
}

func (o *UpdateAccessListOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/permissions][%d] updateAccessListOK ", 200)
}

func (o *UpdateAccessListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessListBadRequest creates a UpdateAccessListBadRequest with default headers values
func NewUpdateAccessListBadRequest() *UpdateAccessListBadRequest {
	return &UpdateAccessListBadRequest{}
}

/*UpdateAccessListBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type UpdateAccessListBadRequest struct {
}

func (o *UpdateAccessListBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/permissions][%d] updateAccessListBadRequest ", 400)
}

func (o *UpdateAccessListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessListUnauthorized creates a UpdateAccessListUnauthorized with default headers values
func NewUpdateAccessListUnauthorized() *UpdateAccessListUnauthorized {
	return &UpdateAccessListUnauthorized{}
}

/*UpdateAccessListUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type UpdateAccessListUnauthorized struct {
}

func (o *UpdateAccessListUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/permissions][%d] updateAccessListUnauthorized ", 401)
}

func (o *UpdateAccessListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessListForbidden creates a UpdateAccessListForbidden with default headers values
func NewUpdateAccessListForbidden() *UpdateAccessListForbidden {
	return &UpdateAccessListForbidden{}
}

/*UpdateAccessListForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type UpdateAccessListForbidden struct {
}

func (o *UpdateAccessListForbidden) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/permissions][%d] updateAccessListForbidden ", 403)
}

func (o *UpdateAccessListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessListTooManyRequests creates a UpdateAccessListTooManyRequests with default headers values
func NewUpdateAccessListTooManyRequests() *UpdateAccessListTooManyRequests {
	return &UpdateAccessListTooManyRequests{}
}

/*UpdateAccessListTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type UpdateAccessListTooManyRequests struct {
}

func (o *UpdateAccessListTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/permissions][%d] updateAccessListTooManyRequests ", 429)
}

func (o *UpdateAccessListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}