// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCompositeEntityReader is a Reader for the DeleteCompositeEntity structure.
type DeleteCompositeEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCompositeEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCompositeEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCompositeEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCompositeEntityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCompositeEntityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCompositeEntityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCompositeEntityOK creates a DeleteCompositeEntityOK with default headers values
func NewDeleteCompositeEntityOK() *DeleteCompositeEntityOK {
	return &DeleteCompositeEntityOK{}
}

/*DeleteCompositeEntityOK handles this case with default header values.

DeleteCompositeEntityOK delete composite entity o k
*/
type DeleteCompositeEntityOK struct {
}

func (o *DeleteCompositeEntityOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] deleteCompositeEntityOK ", 200)
}

func (o *DeleteCompositeEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCompositeEntityBadRequest creates a DeleteCompositeEntityBadRequest with default headers values
func NewDeleteCompositeEntityBadRequest() *DeleteCompositeEntityBadRequest {
	return &DeleteCompositeEntityBadRequest{}
}

/*DeleteCompositeEntityBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type DeleteCompositeEntityBadRequest struct {
}

func (o *DeleteCompositeEntityBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] deleteCompositeEntityBadRequest ", 400)
}

func (o *DeleteCompositeEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCompositeEntityUnauthorized creates a DeleteCompositeEntityUnauthorized with default headers values
func NewDeleteCompositeEntityUnauthorized() *DeleteCompositeEntityUnauthorized {
	return &DeleteCompositeEntityUnauthorized{}
}

/*DeleteCompositeEntityUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type DeleteCompositeEntityUnauthorized struct {
}

func (o *DeleteCompositeEntityUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] deleteCompositeEntityUnauthorized ", 401)
}

func (o *DeleteCompositeEntityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCompositeEntityForbidden creates a DeleteCompositeEntityForbidden with default headers values
func NewDeleteCompositeEntityForbidden() *DeleteCompositeEntityForbidden {
	return &DeleteCompositeEntityForbidden{}
}

/*DeleteCompositeEntityForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type DeleteCompositeEntityForbidden struct {
}

func (o *DeleteCompositeEntityForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] deleteCompositeEntityForbidden ", 403)
}

func (o *DeleteCompositeEntityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCompositeEntityTooManyRequests creates a DeleteCompositeEntityTooManyRequests with default headers values
func NewDeleteCompositeEntityTooManyRequests() *DeleteCompositeEntityTooManyRequests {
	return &DeleteCompositeEntityTooManyRequests{}
}

/*DeleteCompositeEntityTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type DeleteCompositeEntityTooManyRequests struct {
}

func (o *DeleteCompositeEntityTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] deleteCompositeEntityTooManyRequests ", 429)
}

func (o *DeleteCompositeEntityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
