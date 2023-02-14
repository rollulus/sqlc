// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

type FooTypeUserRole string

const (
	FooTypeUserRoleAdmin FooTypeUserRole = "admin"
	FooTypeUserRoleUser  FooTypeUserRole = "user"
)

func (e *FooTypeUserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooTypeUserRole(s)
	case string:
		*e = FooTypeUserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for FooTypeUserRole: %T", src)
	}
	return nil
}

type NullFooTypeUserRole struct {
	FooTypeUserRole FooTypeUserRole
	Valid           bool // Valid is true if FooTypeUserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFooTypeUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.FooTypeUserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FooTypeUserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFooTypeUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FooTypeUserRole), nil
}

type FooUser struct {
	Role NullFooTypeUserRole
}
