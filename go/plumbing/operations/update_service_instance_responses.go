// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// UpdateServiceInstanceReader is a Reader for the UpdateServiceInstance structure.
type UpdateServiceInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateServiceInstanceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateServiceInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateServiceInstanceNoContent creates a UpdateServiceInstanceNoContent with default headers values
func NewUpdateServiceInstanceNoContent() *UpdateServiceInstanceNoContent {
	return &UpdateServiceInstanceNoContent{}
}

/*
UpdateServiceInstanceNoContent handles this case with default header values.

No Content
*/
type UpdateServiceInstanceNoContent struct {
}

func (o *UpdateServiceInstanceNoContent) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/services/{addon}/instances/{instance_id}][%d] updateServiceInstanceNoContent ", 204)
}

func (o *UpdateServiceInstanceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceInstanceDefault creates a UpdateServiceInstanceDefault with default headers values
func NewUpdateServiceInstanceDefault(code int) *UpdateServiceInstanceDefault {
	return &UpdateServiceInstanceDefault{
		_statusCode: code,
	}
}

/*
UpdateServiceInstanceDefault handles this case with default header values.

error
*/
type UpdateServiceInstanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update service instance default response
func (o *UpdateServiceInstanceDefault) Code() int {
	return o._statusCode
}

func (o *UpdateServiceInstanceDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/services/{addon}/instances/{instance_id}][%d] updateServiceInstance default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateServiceInstanceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServiceInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
