// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"encoding/json"
	"fmt"
)

// Bar gives you some dumb info
type Bar struct {
	Name *string `json:"name,omitempty"`
}

func (m *Bar) Validate() error {
	return nil
}

type BarValidationError struct {
	errType, jsonField, field, message string
}

func (e *BarValidationError) ErrType() string {
	return e.errType
}

func (e *BarValidationError) JSONField() string {
	return e.jsonField
}

func (e *BarValidationError) Field() string {
	return e.field
}

func (e *BarValidationError) Message() string {
	return e.message
}

func (e *BarValidationError) Error() string {
	return fmt.Sprintf("%v: %v", e.field, e.message)
}

// Barz gives you lots of dumb info
type Barz []Bar

func (m Barz) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte(`[]`), nil
	}
	return json.Marshal([]Bar(m))
}

func (m Barz) Validate() error {
	for i := range m {
		if err := m[i].Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BarzValidationError struct {
	errType, jsonField, field, message string
}

func (e *BarzValidationError) ErrType() string {
	return e.errType
}

func (e *BarzValidationError) JSONField() string {
	return e.jsonField
}

func (e *BarzValidationError) Field() string {
	return e.field
}

func (e *BarzValidationError) Message() string {
	return e.message
}

func (e *BarzValidationError) Error() string {
	return fmt.Sprintf("%v: %v", e.field, e.message)
}
