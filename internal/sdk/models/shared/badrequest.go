// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Code - A short code indicating the error code returned.
type Code string

const (
	CodeBadRequest Code = "bad_request"
)

func (e Code) ToPointer() *Code {
	return &e
}
func (e *Code) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bad_request":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

type Error struct {
	// A short code indicating the error code returned.
	Code Code `json:"code"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
}

func (o *Error) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}

func (o *Error) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type BadRequest struct {
	Error Error `json:"error"`
}

func (o *BadRequest) GetError() Error {
	if o == nil {
		return Error{}
	}
	return o.Error
}
