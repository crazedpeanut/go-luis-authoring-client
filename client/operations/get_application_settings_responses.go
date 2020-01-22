// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetApplicationSettingsReader is a Reader for the GetApplicationSettings structure.
type GetApplicationSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetApplicationSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetApplicationSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationSettingsOK creates a GetApplicationSettingsOK with default headers values
func NewGetApplicationSettingsOK() *GetApplicationSettingsOK {
	return &GetApplicationSettingsOK{}
}

/*GetApplicationSettingsOK handles this case with default header values.

A JSON object containing the application settings.
*/
type GetApplicationSettingsOK struct {
}

func (o *GetApplicationSettingsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/settings][%d] getApplicationSettingsOK ", 200)
}

func (o *GetApplicationSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationSettingsBadRequest creates a GetApplicationSettingsBadRequest with default headers values
func NewGetApplicationSettingsBadRequest() *GetApplicationSettingsBadRequest {
	return &GetApplicationSettingsBadRequest{}
}

/*GetApplicationSettingsBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetApplicationSettingsBadRequest struct {
}

func (o *GetApplicationSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/settings][%d] getApplicationSettingsBadRequest ", 400)
}

func (o *GetApplicationSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationSettingsUnauthorized creates a GetApplicationSettingsUnauthorized with default headers values
func NewGetApplicationSettingsUnauthorized() *GetApplicationSettingsUnauthorized {
	return &GetApplicationSettingsUnauthorized{}
}

/*GetApplicationSettingsUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type GetApplicationSettingsUnauthorized struct {
}

func (o *GetApplicationSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/settings][%d] getApplicationSettingsUnauthorized ", 401)
}

func (o *GetApplicationSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationSettingsForbidden creates a GetApplicationSettingsForbidden with default headers values
func NewGetApplicationSettingsForbidden() *GetApplicationSettingsForbidden {
	return &GetApplicationSettingsForbidden{}
}

/*GetApplicationSettingsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetApplicationSettingsForbidden struct {
}

func (o *GetApplicationSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/settings][%d] getApplicationSettingsForbidden ", 403)
}

func (o *GetApplicationSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationSettingsTooManyRequests creates a GetApplicationSettingsTooManyRequests with default headers values
func NewGetApplicationSettingsTooManyRequests() *GetApplicationSettingsTooManyRequests {
	return &GetApplicationSettingsTooManyRequests{}
}

/*GetApplicationSettingsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetApplicationSettingsTooManyRequests struct {
}

func (o *GetApplicationSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/settings][%d] getApplicationSettingsTooManyRequests ", 429)
}

func (o *GetApplicationSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
