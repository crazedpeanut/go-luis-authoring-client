// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteVersionReader is a Reader for the DeleteVersion structure.
type DeleteVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVersionOK creates a DeleteVersionOK with default headers values
func NewDeleteVersionOK() *DeleteVersionOK {
	return &DeleteVersionOK{}
}

/*DeleteVersionOK handles this case with default header values.

DeleteVersionOK delete version o k
*/
type DeleteVersionOK struct {
}

func (o *DeleteVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/][%d] deleteVersionOK ", 200)
}

func (o *DeleteVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVersionBadRequest creates a DeleteVersionBadRequest with default headers values
func NewDeleteVersionBadRequest() *DeleteVersionBadRequest {
	return &DeleteVersionBadRequest{}
}

/*DeleteVersionBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type DeleteVersionBadRequest struct {
}

func (o *DeleteVersionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/][%d] deleteVersionBadRequest ", 400)
}

func (o *DeleteVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVersionUnauthorized creates a DeleteVersionUnauthorized with default headers values
func NewDeleteVersionUnauthorized() *DeleteVersionUnauthorized {
	return &DeleteVersionUnauthorized{}
}

/*DeleteVersionUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type DeleteVersionUnauthorized struct {
}

func (o *DeleteVersionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/][%d] deleteVersionUnauthorized ", 401)
}

func (o *DeleteVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVersionForbidden creates a DeleteVersionForbidden with default headers values
func NewDeleteVersionForbidden() *DeleteVersionForbidden {
	return &DeleteVersionForbidden{}
}

/*DeleteVersionForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type DeleteVersionForbidden struct {
}

func (o *DeleteVersionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/][%d] deleteVersionForbidden ", 403)
}

func (o *DeleteVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVersionTooManyRequests creates a DeleteVersionTooManyRequests with default headers values
func NewDeleteVersionTooManyRequests() *DeleteVersionTooManyRequests {
	return &DeleteVersionTooManyRequests{}
}

/*DeleteVersionTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type DeleteVersionTooManyRequests struct {
}

func (o *DeleteVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/][%d] deleteVersionTooManyRequests ", 429)
}

func (o *DeleteVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
