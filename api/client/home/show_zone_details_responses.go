// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gonzolino/gotado/api/models"
)

// ShowZoneDetailsReader is a Reader for the ShowZoneDetails structure.
type ShowZoneDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowZoneDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowZoneDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewShowZoneDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowZoneDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowZoneDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowZoneDetailsOK creates a ShowZoneDetailsOK with default headers values
func NewShowZoneDetailsOK() *ShowZoneDetailsOK {
	return &ShowZoneDetailsOK{}
}

/* ShowZoneDetailsOK describes a response with status code 200, with default header values.

Zone details.
*/
type ShowZoneDetailsOK struct {
	Payload *models.Zone
}

func (o *ShowZoneDetailsOK) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/details][%d] showZoneDetailsOK  %+v", 200, o.Payload)
}
func (o *ShowZoneDetailsOK) GetPayload() *models.Zone {
	return o.Payload
}

func (o *ShowZoneDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneDetailsUnauthorized creates a ShowZoneDetailsUnauthorized with default headers values
func NewShowZoneDetailsUnauthorized() *ShowZoneDetailsUnauthorized {
	return &ShowZoneDetailsUnauthorized{}
}

/* ShowZoneDetailsUnauthorized describes a response with status code 401, with default header values.

User authentication failed.
*/
type ShowZoneDetailsUnauthorized struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/details][%d] showZoneDetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowZoneDetailsUnauthorized) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneDetailsForbidden creates a ShowZoneDetailsForbidden with default headers values
func NewShowZoneDetailsForbidden() *ShowZoneDetailsForbidden {
	return &ShowZoneDetailsForbidden{}
}

/* ShowZoneDetailsForbidden describes a response with status code 403, with default header values.

Authenticated user has no access rights to the requested entity.
*/
type ShowZoneDetailsForbidden struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/details][%d] showZoneDetailsForbidden  %+v", 403, o.Payload)
}
func (o *ShowZoneDetailsForbidden) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneDetailsNotFound creates a ShowZoneDetailsNotFound with default headers values
func NewShowZoneDetailsNotFound() *ShowZoneDetailsNotFound {
	return &ShowZoneDetailsNotFound{}
}

/* ShowZoneDetailsNotFound describes a response with status code 404, with default header values.

Requested entity not found.
*/
type ShowZoneDetailsNotFound struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/details][%d] showZoneDetailsNotFound  %+v", 404, o.Payload)
}
func (o *ShowZoneDetailsNotFound) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
