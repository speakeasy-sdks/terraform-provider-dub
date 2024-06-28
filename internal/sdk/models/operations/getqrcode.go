// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-dub/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

// Level - The level of error correction to use for the QR code. Defaults to `L` if not provided.
type Level string

const (
	LevelL Level = "L"
	LevelM Level = "M"
	LevelQ Level = "Q"
	LevelH Level = "H"
)

func (e Level) ToPointer() *Level {
	return &e
}
func (e *Level) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "L":
		fallthrough
	case "M":
		fallthrough
	case "Q":
		fallthrough
	case "H":
		*e = Level(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Level: %v", v)
	}
}

type GetQRCodeRequest struct {
	// The URL to generate a QR code for.
	URL string `queryParam:"style=form,explode=true,name=url"`
	// The size of the QR code in pixels. Defaults to `600` if not provided.
	Size *float64 `default:"600" queryParam:"style=form,explode=true,name=size"`
	// The level of error correction to use for the QR code. Defaults to `L` if not provided.
	Level *Level `default:"L" queryParam:"style=form,explode=true,name=level"`
	// The foreground color of the QR code in hex format. Defaults to `#000000` if not provided.
	FgColor *string `default:"#000000" queryParam:"style=form,explode=true,name=fgColor"`
	// The background color of the QR code in hex format. Defaults to `#ffffff` if not provided.
	BgColor *string `default:"#FFFFFF" queryParam:"style=form,explode=true,name=bgColor"`
	// Whether to include a margin around the QR code. Defaults to `false` if not provided.
	IncludeMargin *bool `default:"false" queryParam:"style=form,explode=true,name=includeMargin"`
}

func (g GetQRCodeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetQRCodeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetQRCodeRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *GetQRCodeRequest) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetQRCodeRequest) GetLevel() *Level {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *GetQRCodeRequest) GetFgColor() *string {
	if o == nil {
		return nil
	}
	return o.FgColor
}

func (o *GetQRCodeRequest) GetBgColor() *string {
	if o == nil {
		return nil
	}
	return o.BgColor
}

func (o *GetQRCodeRequest) GetIncludeMargin() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeMargin
}

type GetQRCodeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The QR code
	Res *string
	// The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
	BadRequest *shared.BadRequest
	// Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.
	Unauthorized *shared.Unauthorized
	// The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401 Unauthorized, the client's identity is known to the server.
	Forbidden *shared.Forbidden
	// The server cannot find the requested resource.
	NotFound *shared.NotFound
	// This response is sent when a request conflicts with the current state of the server.
	Conflict *shared.Conflict
	// This response is sent when the requested content has been permanently deleted from server, with no forwarding address.
	InviteExpired *shared.InviteExpired
	// The request was well-formed but was unable to be followed due to semantic errors.
	UnprocessableEntity *shared.UnprocessableEntity
	// The user has sent too many requests in a given amount of time ("rate limiting")
	RateLimitExceeded *shared.RateLimitExceeded
	// The server has encountered a situation it does not know how to handle.
	InternalServerError *shared.InternalServerError
}

func (o *GetQRCodeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetQRCodeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetQRCodeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetQRCodeResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}

func (o *GetQRCodeResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *GetQRCodeResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *GetQRCodeResponse) GetForbidden() *shared.Forbidden {
	if o == nil {
		return nil
	}
	return o.Forbidden
}

func (o *GetQRCodeResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *GetQRCodeResponse) GetConflict() *shared.Conflict {
	if o == nil {
		return nil
	}
	return o.Conflict
}

func (o *GetQRCodeResponse) GetInviteExpired() *shared.InviteExpired {
	if o == nil {
		return nil
	}
	return o.InviteExpired
}

func (o *GetQRCodeResponse) GetUnprocessableEntity() *shared.UnprocessableEntity {
	if o == nil {
		return nil
	}
	return o.UnprocessableEntity
}

func (o *GetQRCodeResponse) GetRateLimitExceeded() *shared.RateLimitExceeded {
	if o == nil {
		return nil
	}
	return o.RateLimitExceeded
}

func (o *GetQRCodeResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}
