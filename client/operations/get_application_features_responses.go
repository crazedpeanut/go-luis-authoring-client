// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetApplicationFeaturesReader is a Reader for the GetApplicationFeatures structure.
type GetApplicationFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationFeaturesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetApplicationFeaturesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationFeaturesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetApplicationFeaturesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationFeaturesOK creates a GetApplicationFeaturesOK with default headers values
func NewGetApplicationFeaturesOK() *GetApplicationFeaturesOK {
	return &GetApplicationFeaturesOK{}
}

/*GetApplicationFeaturesOK handles this case with default header values.

A JSON object containing a list of all phraselist features and a list of all pattern featuers in the application.
*/
type GetApplicationFeaturesOK struct {
}

func (o *GetApplicationFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/features][%d] getApplicationFeaturesOK ", 200)
}

func (o *GetApplicationFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationFeaturesBadRequest creates a GetApplicationFeaturesBadRequest with default headers values
func NewGetApplicationFeaturesBadRequest() *GetApplicationFeaturesBadRequest {
	return &GetApplicationFeaturesBadRequest{}
}

/*GetApplicationFeaturesBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetApplicationFeaturesBadRequest struct {
}

func (o *GetApplicationFeaturesBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/features][%d] getApplicationFeaturesBadRequest ", 400)
}

func (o *GetApplicationFeaturesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationFeaturesUnauthorized creates a GetApplicationFeaturesUnauthorized with default headers values
func NewGetApplicationFeaturesUnauthorized() *GetApplicationFeaturesUnauthorized {
	return &GetApplicationFeaturesUnauthorized{}
}

/*GetApplicationFeaturesUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type GetApplicationFeaturesUnauthorized struct {
}

func (o *GetApplicationFeaturesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/features][%d] getApplicationFeaturesUnauthorized ", 401)
}

func (o *GetApplicationFeaturesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationFeaturesForbidden creates a GetApplicationFeaturesForbidden with default headers values
func NewGetApplicationFeaturesForbidden() *GetApplicationFeaturesForbidden {
	return &GetApplicationFeaturesForbidden{}
}

/*GetApplicationFeaturesForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetApplicationFeaturesForbidden struct {
}

func (o *GetApplicationFeaturesForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/features][%d] getApplicationFeaturesForbidden ", 403)
}

func (o *GetApplicationFeaturesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationFeaturesTooManyRequests creates a GetApplicationFeaturesTooManyRequests with default headers values
func NewGetApplicationFeaturesTooManyRequests() *GetApplicationFeaturesTooManyRequests {
	return &GetApplicationFeaturesTooManyRequests{}
}

/*GetApplicationFeaturesTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetApplicationFeaturesTooManyRequests struct {
}

func (o *GetApplicationFeaturesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/features][%d] getApplicationFeaturesTooManyRequests ", 429)
}

func (o *GetApplicationFeaturesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
