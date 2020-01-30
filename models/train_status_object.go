// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrainStatusObject train status object
// swagger:model TrainStatusObject
type TrainStatusObject []*TrainStatusObjectItems0

// Validate validates this train status object
func (m TrainStatusObject) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// TrainStatusObjectItems0 train status object items0
// swagger:model TrainStatusObjectItems0
type TrainStatusObjectItems0 struct {

	// details
	Details *TrainStatusObjectItems0Details `json:"details,omitempty"`

	// model Id
	ModelID string `json:"modelId,omitempty"`
}

// Validate validates this train status object items0
func (m *TrainStatusObjectItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrainStatusObjectItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrainStatusObjectItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrainStatusObjectItems0) UnmarshalBinary(b []byte) error {
	var res TrainStatusObjectItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TrainStatusObjectItems0Details train status object items0 details
// swagger:model TrainStatusObjectItems0Details
type TrainStatusObjectItems0Details struct {

	// example count
	ExampleCount int64 `json:"exampleCount,omitempty"`

	// failure reason
	FailureReason string `json:"failureReason,omitempty"`

	// status
	// Enum: [InProgress Success UptoDate Fail]
	Status string `json:"status,omitempty"`

	// status Id
	StatusID string `json:"statusId,omitempty"`
}

// Validate validates this train status object items0 details
func (m *TrainStatusObjectItems0Details) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var trainStatusObjectItems0DetailsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InProgress","Success","UptoDate","Fail"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trainStatusObjectItems0DetailsTypeStatusPropEnum = append(trainStatusObjectItems0DetailsTypeStatusPropEnum, v)
	}
}

const (

	// TrainStatusObjectItems0DetailsStatusInProgress captures enum value "InProgress"
	TrainStatusObjectItems0DetailsStatusInProgress string = "InProgress"

	// TrainStatusObjectItems0DetailsStatusSuccess captures enum value "Success"
	TrainStatusObjectItems0DetailsStatusSuccess string = "Success"

	// TrainStatusObjectItems0DetailsStatusUptoDate captures enum value "UptoDate"
	TrainStatusObjectItems0DetailsStatusUptoDate string = "UptoDate"

	// TrainStatusObjectItems0DetailsStatusFail captures enum value "Fail"
	TrainStatusObjectItems0DetailsStatusFail string = "Fail"
)

// prop value enum
func (m *TrainStatusObjectItems0Details) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trainStatusObjectItems0DetailsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrainStatusObjectItems0Details) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("details"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrainStatusObjectItems0Details) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrainStatusObjectItems0Details) UnmarshalBinary(b []byte) error {
	var res TrainStatusObjectItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
