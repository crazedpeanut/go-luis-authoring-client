// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPatternFeaturesReader is a Reader for the GetPatternFeatures structure.
type GetPatternFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatternFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatternFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPatternFeaturesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPatternFeaturesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPatternFeaturesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPatternFeaturesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatternFeaturesOK creates a GetPatternFeaturesOK with default headers values
func NewGetPatternFeaturesOK() *GetPatternFeaturesOK {
	return &GetPatternFeaturesOK{}
}

/*GetPatternFeaturesOK handles this case with default header values.

A JSON object containing a list of all pattern features.
*/
type GetPatternFeaturesOK struct {
}

func (o *GetPatternFeaturesOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patterns][%d] getPatternFeaturesOK ", 200)
}

func (o *GetPatternFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatternFeaturesBadRequest creates a GetPatternFeaturesBadRequest with default headers values
func NewGetPatternFeaturesBadRequest() *GetPatternFeaturesBadRequest {
	return &GetPatternFeaturesBadRequest{}
}

/*GetPatternFeaturesBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetPatternFeaturesBadRequest struct {
}

func (o *GetPatternFeaturesBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patterns][%d] getPatternFeaturesBadRequest ", 400)
}

func (o *GetPatternFeaturesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatternFeaturesUnauthorized creates a GetPatternFeaturesUnauthorized with default headers values
func NewGetPatternFeaturesUnauthorized() *GetPatternFeaturesUnauthorized {
	return &GetPatternFeaturesUnauthorized{}
}

/*GetPatternFeaturesUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type GetPatternFeaturesUnauthorized struct {
}

func (o *GetPatternFeaturesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patterns][%d] getPatternFeaturesUnauthorized ", 401)
}

func (o *GetPatternFeaturesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatternFeaturesForbidden creates a GetPatternFeaturesForbidden with default headers values
func NewGetPatternFeaturesForbidden() *GetPatternFeaturesForbidden {
	return &GetPatternFeaturesForbidden{}
}

/*GetPatternFeaturesForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetPatternFeaturesForbidden struct {
}

func (o *GetPatternFeaturesForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patterns][%d] getPatternFeaturesForbidden ", 403)
}

func (o *GetPatternFeaturesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPatternFeaturesTooManyRequests creates a GetPatternFeaturesTooManyRequests with default headers values
func NewGetPatternFeaturesTooManyRequests() *GetPatternFeaturesTooManyRequests {
	return &GetPatternFeaturesTooManyRequests{}
}

/*GetPatternFeaturesTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetPatternFeaturesTooManyRequests struct {
}

func (o *GetPatternFeaturesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/patterns][%d] getPatternFeaturesTooManyRequests ", 429)
}

func (o *GetPatternFeaturesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
