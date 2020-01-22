// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteEntityReader is a Reader for the DeleteEntity structure.
type DeleteEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteEntityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEntityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteEntityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEntityOK creates a DeleteEntityOK with default headers values
func NewDeleteEntityOK() *DeleteEntityOK {
	return &DeleteEntityOK{}
}

/*DeleteEntityOK handles this case with default header values.

DeleteEntityOK delete entity o k
*/
type DeleteEntityOK struct {
}

func (o *DeleteEntityOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] deleteEntityOK ", 200)
}

func (o *DeleteEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntityBadRequest creates a DeleteEntityBadRequest with default headers values
func NewDeleteEntityBadRequest() *DeleteEntityBadRequest {
	return &DeleteEntityBadRequest{}
}

/*DeleteEntityBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type DeleteEntityBadRequest struct {
}

func (o *DeleteEntityBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] deleteEntityBadRequest ", 400)
}

func (o *DeleteEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntityUnauthorized creates a DeleteEntityUnauthorized with default headers values
func NewDeleteEntityUnauthorized() *DeleteEntityUnauthorized {
	return &DeleteEntityUnauthorized{}
}

/*DeleteEntityUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type DeleteEntityUnauthorized struct {
}

func (o *DeleteEntityUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] deleteEntityUnauthorized ", 401)
}

func (o *DeleteEntityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntityForbidden creates a DeleteEntityForbidden with default headers values
func NewDeleteEntityForbidden() *DeleteEntityForbidden {
	return &DeleteEntityForbidden{}
}

/*DeleteEntityForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type DeleteEntityForbidden struct {
}

func (o *DeleteEntityForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] deleteEntityForbidden ", 403)
}

func (o *DeleteEntityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntityTooManyRequests creates a DeleteEntityTooManyRequests with default headers values
func NewDeleteEntityTooManyRequests() *DeleteEntityTooManyRequests {
	return &DeleteEntityTooManyRequests{}
}

/*DeleteEntityTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type DeleteEntityTooManyRequests struct {
}

func (o *DeleteEntityTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] deleteEntityTooManyRequests ", 429)
}

func (o *DeleteEntityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
