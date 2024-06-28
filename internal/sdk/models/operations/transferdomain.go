// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dub/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

type TransferDomainGlobals struct {
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *TransferDomainGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *TransferDomainGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type TransferDomainRequestBody struct {
	// The ID of the new workspace to transfer the domain to.
	NewWorkspaceID string `json:"newWorkspaceId"`
}

func (o *TransferDomainRequestBody) GetNewWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.NewWorkspaceID
}

type TransferDomainRequest struct {
	// The domain name.
	Slug        string                     `pathParam:"style=simple,explode=false,name=slug"`
	RequestBody *TransferDomainRequestBody `request:"mediaType=application/json"`
}

func (o *TransferDomainRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *TransferDomainRequest) GetRequestBody() *TransferDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type TransferDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The domain transfer initiated
	DomainSchema *shared.DomainSchema
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

func (o *TransferDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TransferDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TransferDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TransferDomainResponse) GetDomainSchema() *shared.DomainSchema {
	if o == nil {
		return nil
	}
	return o.DomainSchema
}

func (o *TransferDomainResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *TransferDomainResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *TransferDomainResponse) GetForbidden() *shared.Forbidden {
	if o == nil {
		return nil
	}
	return o.Forbidden
}

func (o *TransferDomainResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *TransferDomainResponse) GetConflict() *shared.Conflict {
	if o == nil {
		return nil
	}
	return o.Conflict
}

func (o *TransferDomainResponse) GetInviteExpired() *shared.InviteExpired {
	if o == nil {
		return nil
	}
	return o.InviteExpired
}

func (o *TransferDomainResponse) GetUnprocessableEntity() *shared.UnprocessableEntity {
	if o == nil {
		return nil
	}
	return o.UnprocessableEntity
}

func (o *TransferDomainResponse) GetRateLimitExceeded() *shared.RateLimitExceeded {
	if o == nil {
		return nil
	}
	return o.RateLimitExceeded
}

func (o *TransferDomainResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}
