// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMongoDBServiceReader is a Reader for the AddMongoDBService structure.
type AddMongoDBServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMongoDBServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddMongoDBServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMongoDBServiceOK creates a AddMongoDBServiceOK with default headers values
func NewAddMongoDBServiceOK() *AddMongoDBServiceOK {
	return &AddMongoDBServiceOK{}
}

/*AddMongoDBServiceOK handles this case with default header values.

A successful response.
*/
type AddMongoDBServiceOK struct {
	Payload *AddMongoDBServiceOKBody
}

func (o *AddMongoDBServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMongoDB][%d] addMongoDbServiceOk  %+v", 200, o.Payload)
}

func (o *AddMongoDBServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMongoDBServiceDefault creates a AddMongoDBServiceDefault with default headers values
func NewAddMongoDBServiceDefault(code int) *AddMongoDBServiceDefault {
	return &AddMongoDBServiceDefault{
		_statusCode: code,
	}
}

/*AddMongoDBServiceDefault handles this case with default header values.

An error response.
*/
type AddMongoDBServiceDefault struct {
	_statusCode int

	Payload *AddMongoDBServiceDefaultBody
}

// Code gets the status code for the add mongo DB service default response
func (o *AddMongoDBServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddMongoDBServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddMongoDB][%d] AddMongoDBService default  %+v", o._statusCode, o.Payload)
}

func (o *AddMongoDBServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMongoDBServiceBody add mongo DB service body
swagger:model AddMongoDBServiceBody
*/
type AddMongoDBServiceBody struct {

	// Access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Access port. Required.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add mongo DB service body
func (o *AddMongoDBServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBServiceDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddMongoDBServiceDefaultBody
*/
type AddMongoDBServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add mongo DB service default body
func (o *AddMongoDBServiceDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBServiceOKBody add mongo DB service OK body
swagger:model AddMongoDBServiceOKBody
*/
type AddMongoDBServiceOKBody struct {

	// mongodb
	Mongodb *AddMongoDBServiceOKBodyMongodb `json:"mongodb,omitempty"`
}

// Validate validates this add mongo DB service OK body
func (o *AddMongoDBServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBServiceOKBody) validateMongodb(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	if o.Mongodb != nil {
		if err := o.Mongodb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbServiceOk" + "." + "mongodb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBServiceOKBodyMongodb MongoDBService represents a generic MongoDB instance.
swagger:model AddMongoDBServiceOKBodyMongodb
*/
type AddMongoDBServiceOKBodyMongodb struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add mongo DB service OK body mongodb
func (o *AddMongoDBServiceOKBodyMongodb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBServiceOKBodyMongodb) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBServiceOKBodyMongodb) UnmarshalBinary(b []byte) error {
	var res AddMongoDBServiceOKBodyMongodb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}