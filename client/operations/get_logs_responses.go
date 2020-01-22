// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLogsReader is a Reader for the GetLogs structure.
type GetLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLogsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsOK creates a GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {
	return &GetLogsOK{}
}

/*GetLogsOK handles this case with default header values.

A CSV file containing the query logs for the past month.
*/
type GetLogsOK struct {
}

func (o *GetLogsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogs/][%d] getLogsOK ", 200)
}

func (o *GetLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsBadRequest creates a GetLogsBadRequest with default headers values
func NewGetLogsBadRequest() *GetLogsBadRequest {
	return &GetLogsBadRequest{}
}

/*GetLogsBadRequest handles this case with default header values.

The app ID is missing or invalid.
The includeResponse is invalid.
*/
type GetLogsBadRequest struct {
}

func (o *GetLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogs/][%d] getLogsBadRequest ", 400)
}

func (o *GetLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsUnauthorized creates a GetLogsUnauthorized with default headers values
func NewGetLogsUnauthorized() *GetLogsUnauthorized {
	return &GetLogsUnauthorized{}
}

/*GetLogsUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type GetLogsUnauthorized struct {
}

func (o *GetLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogs/][%d] getLogsUnauthorized ", 401)
}

func (o *GetLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsForbidden creates a GetLogsForbidden with default headers values
func NewGetLogsForbidden() *GetLogsForbidden {
	return &GetLogsForbidden{}
}

/*GetLogsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetLogsForbidden struct {
}

func (o *GetLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogs/][%d] getLogsForbidden ", 403)
}

func (o *GetLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsTooManyRequests creates a GetLogsTooManyRequests with default headers values
func NewGetLogsTooManyRequests() *GetLogsTooManyRequests {
	return &GetLogsTooManyRequests{}
}

/*GetLogsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetLogsTooManyRequests struct {
}

func (o *GetLogsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogs/][%d] getLogsTooManyRequests ", 429)
}

func (o *GetLogsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
