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

// DeleteZoneOverlayReader is a Reader for the DeleteZoneOverlay structure.
type DeleteZoneOverlayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteZoneOverlayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteZoneOverlayNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteZoneOverlayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteZoneOverlayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteZoneOverlayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteZoneOverlayNoContent creates a DeleteZoneOverlayNoContent with default headers values
func NewDeleteZoneOverlayNoContent() *DeleteZoneOverlayNoContent {
	return &DeleteZoneOverlayNoContent{}
}

/* DeleteZoneOverlayNoContent describes a response with status code 204, with default header values.

Overlay successfully deleted.
*/
type DeleteZoneOverlayNoContent struct {
}

func (o *DeleteZoneOverlayNoContent) Error() string {
	return fmt.Sprintf("[DELETE /homes/{home_id}/zones/{zone_id}/overlay][%d] deleteZoneOverlayNoContent ", 204)
}

func (o *DeleteZoneOverlayNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteZoneOverlayUnauthorized creates a DeleteZoneOverlayUnauthorized with default headers values
func NewDeleteZoneOverlayUnauthorized() *DeleteZoneOverlayUnauthorized {
	return &DeleteZoneOverlayUnauthorized{}
}

/* DeleteZoneOverlayUnauthorized describes a response with status code 401, with default header values.

User authentication failed.
*/
type DeleteZoneOverlayUnauthorized struct {
	Payload *models.ClientErrorModel
}

func (o *DeleteZoneOverlayUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /homes/{home_id}/zones/{zone_id}/overlay][%d] deleteZoneOverlayUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteZoneOverlayUnauthorized) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *DeleteZoneOverlayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZoneOverlayForbidden creates a DeleteZoneOverlayForbidden with default headers values
func NewDeleteZoneOverlayForbidden() *DeleteZoneOverlayForbidden {
	return &DeleteZoneOverlayForbidden{}
}

/* DeleteZoneOverlayForbidden describes a response with status code 403, with default header values.

Authenticated user has no access rights to the requested entity.
*/
type DeleteZoneOverlayForbidden struct {
	Payload *models.ClientErrorModel
}

func (o *DeleteZoneOverlayForbidden) Error() string {
	return fmt.Sprintf("[DELETE /homes/{home_id}/zones/{zone_id}/overlay][%d] deleteZoneOverlayForbidden  %+v", 403, o.Payload)
}
func (o *DeleteZoneOverlayForbidden) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *DeleteZoneOverlayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZoneOverlayNotFound creates a DeleteZoneOverlayNotFound with default headers values
func NewDeleteZoneOverlayNotFound() *DeleteZoneOverlayNotFound {
	return &DeleteZoneOverlayNotFound{}
}

/* DeleteZoneOverlayNotFound describes a response with status code 404, with default header values.

Requested entity not found.
*/
type DeleteZoneOverlayNotFound struct {
	Payload *models.ClientErrorModel
}

func (o *DeleteZoneOverlayNotFound) Error() string {
	return fmt.Sprintf("[DELETE /homes/{home_id}/zones/{zone_id}/overlay][%d] deleteZoneOverlayNotFound  %+v", 404, o.Payload)
}
func (o *DeleteZoneOverlayNotFound) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *DeleteZoneOverlayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
