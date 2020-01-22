// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetApplicationVersionKeyReader is a Reader for the GetApplicationVersionKey structure.
type GetApplicationVersionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationVersionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationVersionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationVersionKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetApplicationVersionKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetApplicationVersionKeyGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetApplicationVersionKeyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationVersionKeyOK creates a GetApplicationVersionKeyOK with default headers values
func NewGetApplicationVersionKeyOK() *GetApplicationVersionKeyOK {
	return &GetApplicationVersionKeyOK{}
}

/*GetApplicationVersionKeyOK handles this case with default header values.

The subscription key for this application's endpoint.
*/
type GetApplicationVersionKeyOK struct {
}

func (o *GetApplicationVersionKeyOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/assignedkey][%d] getApplicationVersionKeyOK ", 200)
}

func (o *GetApplicationVersionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationVersionKeyBadRequest creates a GetApplicationVersionKeyBadRequest with default headers values
func NewGetApplicationVersionKeyBadRequest() *GetApplicationVersionKeyBadRequest {
	return &GetApplicationVersionKeyBadRequest{}
}

/*GetApplicationVersionKeyBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetApplicationVersionKeyBadRequest struct {
}

func (o *GetApplicationVersionKeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/assignedkey][%d] getApplicationVersionKeyBadRequest ", 400)
}

func (o *GetApplicationVersionKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationVersionKeyUnauthorized creates a GetApplicationVersionKeyUnauthorized with default headers values
func NewGetApplicationVersionKeyUnauthorized() *GetApplicationVersionKeyUnauthorized {
	return &GetApplicationVersionKeyUnauthorized{}
}

/*GetApplicationVersionKeyUnauthorized handles this case with default header values.

You do not have access.
<br>
Reasons can include:<br>
<ul>
<li>invalid, malformed, or empty subscription key
<li>you are not the owner or collaborator
<li>invalid order of API calls
<li>subscription key doesn't match region
</ul>
*/
type GetApplicationVersionKeyUnauthorized struct {
}

func (o *GetApplicationVersionKeyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/assignedkey][%d] getApplicationVersionKeyUnauthorized ", 401)
}

func (o *GetApplicationVersionKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationVersionKeyGone creates a GetApplicationVersionKeyGone with default headers values
func NewGetApplicationVersionKeyGone() *GetApplicationVersionKeyGone {
	return &GetApplicationVersionKeyGone{}
}

/*GetApplicationVersionKeyGone handles this case with default header values.

THIS API IS DEPRECATED.
*/
type GetApplicationVersionKeyGone struct {
}

func (o *GetApplicationVersionKeyGone) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/assignedkey][%d] getApplicationVersionKeyGone ", 410)
}

func (o *GetApplicationVersionKeyGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationVersionKeyTooManyRequests creates a GetApplicationVersionKeyTooManyRequests with default headers values
func NewGetApplicationVersionKeyTooManyRequests() *GetApplicationVersionKeyTooManyRequests {
	return &GetApplicationVersionKeyTooManyRequests{}
}

/*GetApplicationVersionKeyTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetApplicationVersionKeyTooManyRequests struct {
}

func (o *GetApplicationVersionKeyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/assignedkey][%d] getApplicationVersionKeyTooManyRequests ", 429)
}

func (o *GetApplicationVersionKeyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
