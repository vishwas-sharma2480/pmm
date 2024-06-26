// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/*
GetRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRoleOK struct {
	Payload *GetRoleOKBody
}

func (o *GetRoleOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Role/Get][%d] getRoleOk  %+v", 200, o.Payload)
}

func (o *GetRoleOK) GetPayload() *GetRoleOKBody {
	return o.Payload
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetRoleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleDefault creates a GetRoleDefault with default headers values
func NewGetRoleDefault(code int) *GetRoleDefault {
	return &GetRoleDefault{
		_statusCode: code,
	}
}

/*
GetRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRoleDefault struct {
	_statusCode int

	Payload *GetRoleDefaultBody
}

// Code gets the status code for the get role default response
func (o *GetRoleDefault) Code() int {
	return o._statusCode
}

func (o *GetRoleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Role/Get][%d] GetRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleDefault) GetPayload() *GetRoleDefaultBody {
	return o.Payload
}

func (o *GetRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetRoleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetRoleBody get role body
swagger:model GetRoleBody
*/
type GetRoleBody struct {
	// role id
	RoleID int64 `json:"role_id,omitempty"`
}

// Validate validates this get role body
func (o *GetRoleBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get role body based on context it is used
func (o *GetRoleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleBody) UnmarshalBinary(b []byte) error {
	var res GetRoleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRoleDefaultBody get role default body
swagger:model GetRoleDefaultBody
*/
type GetRoleDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetRoleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get role default body
func (o *GetRoleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get role default body based on the context it is used
func (o *GetRoleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetRoleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRoleDefaultBodyDetailsItems0 get role default body details items0
swagger:model GetRoleDefaultBodyDetailsItems0
*/
type GetRoleDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get role default body details items0
func (o *GetRoleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get role default body details items0 based on context it is used
func (o *GetRoleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetRoleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRoleOKBody get role OK body
swagger:model GetRoleOKBody
*/
type GetRoleOKBody struct {
	// role id
	RoleID int64 `json:"role_id,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this get role OK body
func (o *GetRoleOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get role OK body based on context it is used
func (o *GetRoleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleOKBody) UnmarshalBinary(b []byte) error {
	var res GetRoleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
