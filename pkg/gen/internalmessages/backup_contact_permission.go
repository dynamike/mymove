// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BackupContactPermission Permissions
//
// swagger:model BackupContactPermission
type BackupContactPermission string

func NewBackupContactPermission(value BackupContactPermission) *BackupContactPermission {
	v := value
	return &v
}

const (

	// BackupContactPermissionNONE captures enum value "NONE"
	BackupContactPermissionNONE BackupContactPermission = "NONE"

	// BackupContactPermissionVIEW captures enum value "VIEW"
	BackupContactPermissionVIEW BackupContactPermission = "VIEW"

	// BackupContactPermissionEDIT captures enum value "EDIT"
	BackupContactPermissionEDIT BackupContactPermission = "EDIT"
)

// for schema
var backupContactPermissionEnum []interface{}

func init() {
	var res []BackupContactPermission
	if err := json.Unmarshal([]byte(`["NONE","VIEW","EDIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupContactPermissionEnum = append(backupContactPermissionEnum, v)
	}
}

func (m BackupContactPermission) validateBackupContactPermissionEnum(path, location string, value BackupContactPermission) error {
	if err := validate.EnumCase(path, location, value, backupContactPermissionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup contact permission
func (m BackupContactPermission) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupContactPermissionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup contact permission based on context it is used
func (m BackupContactPermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
