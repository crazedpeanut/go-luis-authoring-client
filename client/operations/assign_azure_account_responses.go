// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AssignAzureAccountReader is a Reader for the AssignAzureAccount structure.
type AssignAzureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignAzureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAssignAzureAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignAzureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAssignAzureAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignAzureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAssignAzureAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAssignAzureAccountCreated creates a AssignAzureAccountCreated with default headers values
func NewAssignAzureAccountCreated() *AssignAzureAccountCreated {
	return &AssignAzureAccountCreated{}
}

/*AssignAzureAccountCreated handles this case with default header values.

A message indicating success of operation.
*/
type AssignAzureAccountCreated struct {
}

func (o *AssignAzureAccountCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/azureaccounts][%d] assignAzureAccountCreated ", 201)
}

func (o *AssignAzureAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignAzureAccountBadRequest creates a AssignAzureAccountBadRequest with default headers values
func NewAssignAzureAccountBadRequest() *AssignAzureAccountBadRequest {
	return &AssignAzureAccountBadRequest{}
}

/*AssignAzureAccountBadRequest handles this case with default header values.

The app ID is missing or invalid.
The includeResponse is invalid.
*/
type AssignAzureAccountBadRequest struct {
}

func (o *AssignAzureAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/azureaccounts][%d] assignAzureAccountBadRequest ", 400)
}

func (o *AssignAzureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignAzureAccountUnauthorized creates a AssignAzureAccountUnauthorized with default headers values
func NewAssignAzureAccountUnauthorized() *AssignAzureAccountUnauthorized {
	return &AssignAzureAccountUnauthorized{}
}

/*AssignAzureAccountUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type AssignAzureAccountUnauthorized struct {
}

func (o *AssignAzureAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/azureaccounts][%d] assignAzureAccountUnauthorized ", 401)
}

func (o *AssignAzureAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignAzureAccountForbidden creates a AssignAzureAccountForbidden with default headers values
func NewAssignAzureAccountForbidden() *AssignAzureAccountForbidden {
	return &AssignAzureAccountForbidden{}
}

/*AssignAzureAccountForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type AssignAzureAccountForbidden struct {
}

func (o *AssignAzureAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/azureaccounts][%d] assignAzureAccountForbidden ", 403)
}

func (o *AssignAzureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignAzureAccountTooManyRequests creates a AssignAzureAccountTooManyRequests with default headers values
func NewAssignAzureAccountTooManyRequests() *AssignAzureAccountTooManyRequests {
	return &AssignAzureAccountTooManyRequests{}
}

/*AssignAzureAccountTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type AssignAzureAccountTooManyRequests struct {
}

func (o *AssignAzureAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/azureaccounts][%d] assignAzureAccountTooManyRequests ", 429)
}

func (o *AssignAzureAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
