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

// CreateDNSRecordReader is a Reader for the CreateDNSRecord structure.
type CreateDNSRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDNSRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDNSRecordCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDNSRecordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDNSRecordCreated creates a CreateDNSRecordCreated with default headers values
func NewCreateDNSRecordCreated() *CreateDNSRecordCreated {
	return &CreateDNSRecordCreated{}
}

/*
CreateDNSRecordCreated handles this case with default header values.

Created
*/
type CreateDNSRecordCreated struct {
	Payload *models.DNSRecord
}

func (o *CreateDNSRecordCreated) Error() string {
	return fmt.Sprintf("[POST /dns_zones/{zone_id}/dns_records][%d] createDnsRecordCreated  %+v", 201, o.Payload)
}

func (o *CreateDNSRecordCreated) GetPayload() *models.DNSRecord {
	return o.Payload
}

func (o *CreateDNSRecordCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDNSRecordDefault creates a CreateDNSRecordDefault with default headers values
func NewCreateDNSRecordDefault(code int) *CreateDNSRecordDefault {
	return &CreateDNSRecordDefault{
		_statusCode: code,
	}
}

/*
CreateDNSRecordDefault handles this case with default header values.

error
*/
type CreateDNSRecordDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create Dns record default response
func (o *CreateDNSRecordDefault) Code() int {
	return o._statusCode
}

func (o *CreateDNSRecordDefault) Error() string {
	return fmt.Sprintf("[POST /dns_zones/{zone_id}/dns_records][%d] createDnsRecord default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDNSRecordDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDNSRecordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
