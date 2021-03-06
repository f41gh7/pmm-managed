// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIScrapeTargetHealth ScrapeTargetHealth represents Prometheus scrape target health: unknown, down, or up.
// swagger:model apiScrapeTargetHealth
type APIScrapeTargetHealth struct {

	// health
	Health ScrapeTargetHealthHealth `json:"health,omitempty"`

	// "instance" label value, may be different from target due to relabeling
	Instance string `json:"instance,omitempty"`

	// "job" label value, may be different from job_name due to relabeling
	Job string `json:"job,omitempty"`

	// Original scrape job name
	JobName string `json:"job_name,omitempty"`

	// Original target
	Target string `json:"target,omitempty"`
}

// Validate validates this api scrape target health
func (m *APIScrapeTargetHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIScrapeTargetHealth) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if err := m.Health.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("health")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIScrapeTargetHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIScrapeTargetHealth) UnmarshalBinary(b []byte) error {
	var res APIScrapeTargetHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
