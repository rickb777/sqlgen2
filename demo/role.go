package demo

import (
	"fmt"
	"database/sql/driver"
)

type Role int8

const (
	Unknown   Role = iota
	UserRole
	AdminRole
)

func (v Role) String() string {
	switch v {
	case UserRole:
		return "user"
	case AdminRole:
		return "admin"
	}
	return ""
}

func (v *Role) Parse(s string) error {
	switch s {
	case "user":
		*v = UserRole
	case "admin":
		*v = AdminRole
	case "":
	default:
		panic(s)
	}
	return nil
}

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case string:
		r.Parse(value.(string))
	case []byte:
		r.Parse(string(value.([]byte)))
	case nil:
		return nil
	default:
		return fmt.Errorf("Expected a string but got %T %+v", value, value)
	}
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return r.String(), nil
}

