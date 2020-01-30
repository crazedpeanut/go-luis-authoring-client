// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr593889745aca2f0e24581850Reader is a Reader for the Nr593889745aca2f0e24581850 structure.
type Nr593889745aca2f0e24581850Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr593889745aca2f0e24581850Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNr593889745aca2f0e24581850Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr593889745aca2f0e24581850BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr593889745aca2f0e24581850Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr593889745aca2f0e24581850Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr593889745aca2f0e24581850TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr593889745aca2f0e24581850Created creates a Nr593889745aca2f0e24581850Created with default headers values
func NewNr593889745aca2f0e24581850Created() *Nr593889745aca2f0e24581850Created {
	return &Nr593889745aca2f0e24581850Created{}
}

/*Nr593889745aca2f0e24581850Created handles this case with default header values.

The ID of the created model.
*/
type Nr593889745aca2f0e24581850Created struct {
}

func (o *Nr593889745aca2f0e24581850Created) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/children][%d] 593889745aca2f0e24581850Created ", 201)
}

func (o *Nr593889745aca2f0e24581850Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr593889745aca2f0e24581850BadRequest creates a Nr593889745aca2f0e24581850BadRequest with default headers values
func NewNr593889745aca2f0e24581850BadRequest() *Nr593889745aca2f0e24581850BadRequest {
	return &Nr593889745aca2f0e24581850BadRequest{}
}

/*Nr593889745aca2f0e24581850BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.

This error can be returned if you are attempting to create a composite entity child when the max child count is exceeded.
```
  "error": {
    "code": "BadArgument",
    "message": "Number of entity children cannot exceed 10"
  }
```
*/
type Nr593889745aca2f0e24581850BadRequest struct {
}

func (o *Nr593889745aca2f0e24581850BadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/children][%d] 593889745aca2f0e24581850BadRequest ", 400)
}

func (o *Nr593889745aca2f0e24581850BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr593889745aca2f0e24581850Unauthorized creates a Nr593889745aca2f0e24581850Unauthorized with default headers values
func NewNr593889745aca2f0e24581850Unauthorized() *Nr593889745aca2f0e24581850Unauthorized {
	return &Nr593889745aca2f0e24581850Unauthorized{}
}

/*Nr593889745aca2f0e24581850Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr593889745aca2f0e24581850Unauthorized struct {
}

func (o *Nr593889745aca2f0e24581850Unauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/children][%d] 593889745aca2f0e24581850Unauthorized ", 401)
}

func (o *Nr593889745aca2f0e24581850Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr593889745aca2f0e24581850Forbidden creates a Nr593889745aca2f0e24581850Forbidden with default headers values
func NewNr593889745aca2f0e24581850Forbidden() *Nr593889745aca2f0e24581850Forbidden {
	return &Nr593889745aca2f0e24581850Forbidden{}
}

/*Nr593889745aca2f0e24581850Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr593889745aca2f0e24581850Forbidden struct {
}

func (o *Nr593889745aca2f0e24581850Forbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/children][%d] 593889745aca2f0e24581850Forbidden ", 403)
}

func (o *Nr593889745aca2f0e24581850Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr593889745aca2f0e24581850TooManyRequests creates a Nr593889745aca2f0e24581850TooManyRequests with default headers values
func NewNr593889745aca2f0e24581850TooManyRequests() *Nr593889745aca2f0e24581850TooManyRequests {
	return &Nr593889745aca2f0e24581850TooManyRequests{}
}

/*Nr593889745aca2f0e24581850TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr593889745aca2f0e24581850TooManyRequests struct {
}

func (o *Nr593889745aca2f0e24581850TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/children][%d] 593889745aca2f0e24581850TooManyRequests ", 429)
}

func (o *Nr593889745aca2f0e24581850TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}