// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UnauthorizedCode - A short code indicating the error code returned.
type UnauthorizedCode string

const (
	UnauthorizedCodeUnauthorized UnauthorizedCode = "unauthorized"
)

func (e UnauthorizedCode) ToPointer() *UnauthorizedCode {
	return &e
}
func (e *UnauthorizedCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unauthorized":
		*e = UnauthorizedCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UnauthorizedCode: %v", v)
	}
}

type UnauthorizedError struct {
	// A short code indicating the error code returned.
	Code UnauthorizedCode `json:"code"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
}

func (o *UnauthorizedError) GetCode() UnauthorizedCode {
	if o == nil {
		return UnauthorizedCode("")
	}
	return o.Code
}

func (o *UnauthorizedError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *UnauthorizedError) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

type Unauthorized struct {
	Error UnauthorizedError `json:"error"`
}

func (o *Unauthorized) GetError() UnauthorizedError {
	if o == nil {
		return UnauthorizedError{}
	}
	return o.Error
}