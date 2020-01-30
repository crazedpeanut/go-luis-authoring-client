// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/crazedpeanut/go-luis-authoring-client/models"
)

// Nr5ae02f7ed5b81c092c6cf2c3Reader is a Reader for the Nr5ae02f7ed5b81c092c6cf2c3 structure.
type Nr5ae02f7ed5b81c092c6cf2c3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ae02f7ed5b81c092c6cf2c3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ae02f7ed5b81c092c6cf2c3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewNr5ae02f7ed5b81c092c6cf2c3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNr5ae02f7ed5b81c092c6cf2c3Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ae02f7ed5b81c092c6cf2c3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ae02f7ed5b81c092c6cf2c3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ae02f7ed5b81c092c6cf2c3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ae02f7ed5b81c092c6cf2c3OK creates a Nr5ae02f7ed5b81c092c6cf2c3OK with default headers values
func NewNr5ae02f7ed5b81c092c6cf2c3OK() *Nr5ae02f7ed5b81c092c6cf2c3OK {
	return &Nr5ae02f7ed5b81c092c6cf2c3OK{}
}

/*Nr5ae02f7ed5b81c092c6cf2c3OK handles this case with default header values.

A CSV file containing the query logs for the past month.
*/
type Nr5ae02f7ed5b81c092c6cf2c3OK struct {
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3OK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogsasync/][%d] 5ae02f7ed5b81c092c6cf2c3OK ", 200)
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ae02f7ed5b81c092c6cf2c3Created creates a Nr5ae02f7ed5b81c092c6cf2c3Created with default headers values
func NewNr5ae02f7ed5b81c092c6cf2c3Created() *Nr5ae02f7ed5b81c092c6cf2c3Created {
	return &Nr5ae02f7ed5b81c092c6cf2c3Created{}
}

/*Nr5ae02f7ed5b81c092c6cf2c3Created handles this case with default header values.

A string mentions that the download process is finished.
*/
type Nr5ae02f7ed5b81c092c6cf2c3Created struct {
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Created) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogsasync/][%d] 5ae02f7ed5b81c092c6cf2c3Created ", 201)
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ae02f7ed5b81c092c6cf2c3Accepted creates a Nr5ae02f7ed5b81c092c6cf2c3Accepted with default headers values
func NewNr5ae02f7ed5b81c092c6cf2c3Accepted() *Nr5ae02f7ed5b81c092c6cf2c3Accepted {
	return &Nr5ae02f7ed5b81c092c6cf2c3Accepted{}
}

/*Nr5ae02f7ed5b81c092c6cf2c3Accepted handles this case with default header values.

A string mentions that the download process is still in progress.
*/
type Nr5ae02f7ed5b81c092c6cf2c3Accepted struct {
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Accepted) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogsasync/][%d] 5ae02f7ed5b81c092c6cf2c3Accepted ", 202)
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ae02f7ed5b81c092c6cf2c3BadRequest creates a Nr5ae02f7ed5b81c092c6cf2c3BadRequest with default headers values
func NewNr5ae02f7ed5b81c092c6cf2c3BadRequest() *Nr5ae02f7ed5b81c092c6cf2c3BadRequest {
	return &Nr5ae02f7ed5b81c092c6cf2c3BadRequest{}
}

/*Nr5ae02f7ed5b81c092c6cf2c3BadRequest handles this case with default header values.

The app ID is missing or invalid.
*/
type Nr5ae02f7ed5b81c092c6cf2c3BadRequest struct {
	Payload *models.ErrorResponseObject
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3BadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogsasync/][%d] 5ae02f7ed5b81c092c6cf2c3BadRequest  %+v", 400, o.Payload)
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3BadRequest) GetPayload() *models.ErrorResponseObject {
	return o.Payload
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNr5ae02f7ed5b81c092c6cf2c3Unauthorized creates a Nr5ae02f7ed5b81c092c6cf2c3Unauthorized with default headers values
func NewNr5ae02f7ed5b81c092c6cf2c3Unauthorized() *Nr5ae02f7ed5b81c092c6cf2c3Unauthorized {
	return &Nr5ae02f7ed5b81c092c6cf2c3Unauthorized{}
}

/*Nr5ae02f7ed5b81c092c6cf2c3Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ae02f7ed5b81c092c6cf2c3Unauthorized struct {
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogsasync/][%d] 5ae02f7ed5b81c092c6cf2c3Unauthorized ", 401)
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ae02f7ed5b81c092c6cf2c3Forbidden creates a Nr5ae02f7ed5b81c092c6cf2c3Forbidden with default headers values
func NewNr5ae02f7ed5b81c092c6cf2c3Forbidden() *Nr5ae02f7ed5b81c092c6cf2c3Forbidden {
	return &Nr5ae02f7ed5b81c092c6cf2c3Forbidden{}
}

/*Nr5ae02f7ed5b81c092c6cf2c3Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ae02f7ed5b81c092c6cf2c3Forbidden struct {
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Forbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/querylogsasync/][%d] 5ae02f7ed5b81c092c6cf2c3Forbidden ", 403)
}

func (o *Nr5ae02f7ed5b81c092c6cf2c3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}