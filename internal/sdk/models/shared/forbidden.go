// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ForbiddenCode - A short code indicating the error code returned.
type ForbiddenCode string

const (
	ForbiddenCodeForbidden ForbiddenCode = "forbidden"
)

func (e ForbiddenCode) ToPointer() *ForbiddenCode {
	return &e
}
func (e *ForbiddenCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "forbidden":
		*e = ForbiddenCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ForbiddenCode: %v", v)
	}
}

type ForbiddenError struct {
	// A short code indicating the error code returned.
	Code ForbiddenCode `json:"code"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
}

func (o *ForbiddenError) GetCode() ForbiddenCode {
	if o == nil {
		return ForbiddenCode("")
	}
	return o.Code
}

func (o *ForbiddenError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ForbiddenError) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

type Forbidden struct {
	Error ForbiddenError `json:"error"`
}

func (o *Forbidden) GetError() ForbiddenError {
	if o == nil {
		return ForbiddenError{}
	}
	return o.Error
}