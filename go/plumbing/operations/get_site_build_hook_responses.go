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

// GetSiteBuildHookReader is a Reader for the GetSiteBuildHook structure.
type GetSiteBuildHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteBuildHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteBuildHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSiteBuildHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteBuildHookOK creates a GetSiteBuildHookOK with default headers values
func NewGetSiteBuildHookOK() *GetSiteBuildHookOK {
	return &GetSiteBuildHookOK{}
}

/*
GetSiteBuildHookOK handles this case with default header values.

OK
*/
type GetSiteBuildHookOK struct {
	Payload *models.BuildHook
}

func (o *GetSiteBuildHookOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/build_hooks/{id}][%d] getSiteBuildHookOK  %+v", 200, o.Payload)
}

func (o *GetSiteBuildHookOK) GetPayload() *models.BuildHook {
	return o.Payload
}

func (o *GetSiteBuildHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteBuildHookDefault creates a GetSiteBuildHookDefault with default headers values
func NewGetSiteBuildHookDefault(code int) *GetSiteBuildHookDefault {
	return &GetSiteBuildHookDefault{
		_statusCode: code,
	}
}

/*
GetSiteBuildHookDefault handles this case with default header values.

error
*/
type GetSiteBuildHookDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get site build hook default response
func (o *GetSiteBuildHookDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteBuildHookDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/build_hooks/{id}][%d] getSiteBuildHook default  %+v", o._statusCode, o.Payload)
}

func (o *GetSiteBuildHookDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSiteBuildHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
