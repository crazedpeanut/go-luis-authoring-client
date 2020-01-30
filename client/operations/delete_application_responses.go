// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteApplicationReader is a Reader for the DeleteApplication structure.
type DeleteApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteApplicationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApplicationOK creates a DeleteApplicationOK with default headers values
func NewDeleteApplicationOK() *DeleteApplicationOK {
	return &DeleteApplicationOK{}
}

/*DeleteApplicationOK handles this case with default header values.

DeleteApplicationOK delete application o k
*/
type DeleteApplicationOK struct {
}

func (o *DeleteApplicationOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] deleteApplicationOK ", 200)
}

func (o *DeleteApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationBadRequest creates a DeleteApplicationBadRequest with default headers values
func NewDeleteApplicationBadRequest() *DeleteApplicationBadRequest {
	return &DeleteApplicationBadRequest{}
}

/*DeleteApplicationBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type DeleteApplicationBadRequest struct {
}

func (o *DeleteApplicationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] deleteApplicationBadRequest ", 400)
}

func (o *DeleteApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationUnauthorized creates a DeleteApplicationUnauthorized with default headers values
func NewDeleteApplicationUnauthorized() *DeleteApplicationUnauthorized {
	return &DeleteApplicationUnauthorized{}
}

/*DeleteApplicationUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type DeleteApplicationUnauthorized struct {
}

func (o *DeleteApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] deleteApplicationUnauthorized ", 401)
}

func (o *DeleteApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationForbidden creates a DeleteApplicationForbidden with default headers values
func NewDeleteApplicationForbidden() *DeleteApplicationForbidden {
	return &DeleteApplicationForbidden{}
}

/*DeleteApplicationForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type DeleteApplicationForbidden struct {
}

func (o *DeleteApplicationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] deleteApplicationForbidden ", 403)
}

func (o *DeleteApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationTooManyRequests creates a DeleteApplicationTooManyRequests with default headers values
func NewDeleteApplicationTooManyRequests() *DeleteApplicationTooManyRequests {
	return &DeleteApplicationTooManyRequests{}
}

/*DeleteApplicationTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type DeleteApplicationTooManyRequests struct {
}

func (o *DeleteApplicationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}][%d] deleteApplicationTooManyRequests ", 429)
}

func (o *DeleteApplicationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}