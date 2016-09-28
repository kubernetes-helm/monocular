package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChartVersion chart version

swagger:model chartVersion
*/
type ChartVersion struct {

	/* checksum

	Required: true
	Min Length: 1
	*/
	Checksum *string `json:"checksum"`

	/* created

	Required: true
	Min Length: 1
	*/
	Created *string `json:"created"`

	/* description

	Required: true
	Min Length: 1
	*/
	Description *string `json:"description"`

	/* home

	Required: true
	Min Length: 1
	*/
	Home *string `json:"home"`

	/* name

	Required: true
	Min Length: 1
	*/
	Name *string `json:"name"`

	/* url

	Required: true
	Min Length: 1
	*/
	URL *string `json:"url"`

	/* version

	Required: true
	Min Length: 1
	*/
	Version *string `json:"version"`
}

// Validate validates this chart version
func (m *ChartVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecksum(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHome(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartVersion) validateChecksum(formats strfmt.Registry) error {

	if err := validate.Required("checksum", "body", m.Checksum); err != nil {
		return err
	}

	if err := validate.MinLength("checksum", "body", string(*m.Checksum), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.MinLength("created", "body", string(*m.Created), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateHome(formats strfmt.Registry) error {

	if err := validate.Required("home", "body", m.Home); err != nil {
		return err
	}

	if err := validate.MinLength("home", "body", string(*m.Home), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", string(*m.URL), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}