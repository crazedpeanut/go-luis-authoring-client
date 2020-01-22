// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddPrebuildEntitiesReader is a Reader for the AddPrebuildEntities structure.
type AddPrebuildEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPrebuildEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddPrebuildEntitiesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddPrebuildEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddPrebuildEntitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddPrebuildEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddPrebuildEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddPrebuildEntitiesCreated creates a AddPrebuildEntitiesCreated with default headers values
func NewAddPrebuildEntitiesCreated() *AddPrebuildEntitiesCreated {
	return &AddPrebuildEntitiesCreated{}
}

/*AddPrebuildEntitiesCreated handles this case with default header values.

An array of the created prebuilt extractor infos.
*/
type AddPrebuildEntitiesCreated struct {
}

func (o *AddPrebuildEntitiesCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/prebuilts][%d] addPrebuildEntitiesCreated ", 201)
}

func (o *AddPrebuildEntitiesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPrebuildEntitiesBadRequest creates a AddPrebuildEntitiesBadRequest with default headers values
func NewAddPrebuildEntitiesBadRequest() *AddPrebuildEntitiesBadRequest {
	return &AddPrebuildEntitiesBadRequest{}
}

/*AddPrebuildEntitiesBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type AddPrebuildEntitiesBadRequest struct {
}

func (o *AddPrebuildEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/prebuilts][%d] addPrebuildEntitiesBadRequest ", 400)
}

func (o *AddPrebuildEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPrebuildEntitiesUnauthorized creates a AddPrebuildEntitiesUnauthorized with default headers values
func NewAddPrebuildEntitiesUnauthorized() *AddPrebuildEntitiesUnauthorized {
	return &AddPrebuildEntitiesUnauthorized{}
}

/*AddPrebuildEntitiesUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type AddPrebuildEntitiesUnauthorized struct {
}

func (o *AddPrebuildEntitiesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/prebuilts][%d] addPrebuildEntitiesUnauthorized ", 401)
}

func (o *AddPrebuildEntitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPrebuildEntitiesForbidden creates a AddPrebuildEntitiesForbidden with default headers values
func NewAddPrebuildEntitiesForbidden() *AddPrebuildEntitiesForbidden {
	return &AddPrebuildEntitiesForbidden{}
}

/*AddPrebuildEntitiesForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type AddPrebuildEntitiesForbidden struct {
}

func (o *AddPrebuildEntitiesForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/prebuilts][%d] addPrebuildEntitiesForbidden ", 403)
}

func (o *AddPrebuildEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPrebuildEntitiesTooManyRequests creates a AddPrebuildEntitiesTooManyRequests with default headers values
func NewAddPrebuildEntitiesTooManyRequests() *AddPrebuildEntitiesTooManyRequests {
	return &AddPrebuildEntitiesTooManyRequests{}
}

/*AddPrebuildEntitiesTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type AddPrebuildEntitiesTooManyRequests struct {
}

func (o *AddPrebuildEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/prebuilts][%d] addPrebuildEntitiesTooManyRequests ", 429)
}

func (o *AddPrebuildEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
