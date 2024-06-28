// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

type CreateWorkspaceRequestBody struct {
	Name   string  `json:"name"`
	Slug   string  `json:"slug"`
	Domain *string `json:"domain,omitempty"`
}

func (o *CreateWorkspaceRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateWorkspaceRequestBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreateWorkspaceRequestBody) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

type CreateWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The created workspace
	WorkspaceSchema *shared.WorkspaceSchema
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

func (o *CreateWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWorkspaceResponse) GetWorkspaceSchema() *shared.WorkspaceSchema {
	if o == nil {
		return nil
	}
	return o.WorkspaceSchema
}

func (o *CreateWorkspaceResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateWorkspaceResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateWorkspaceResponse) GetForbidden() *shared.Forbidden {
	if o == nil {
		return nil
	}
	return o.Forbidden
}

func (o *CreateWorkspaceResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateWorkspaceResponse) GetConflict() *shared.Conflict {
	if o == nil {
		return nil
	}
	return o.Conflict
}

func (o *CreateWorkspaceResponse) GetInviteExpired() *shared.InviteExpired {
	if o == nil {
		return nil
	}
	return o.InviteExpired
}

func (o *CreateWorkspaceResponse) GetUnprocessableEntity() *shared.UnprocessableEntity {
	if o == nil {
		return nil
	}
	return o.UnprocessableEntity
}

func (o *CreateWorkspaceResponse) GetRateLimitExceeded() *shared.RateLimitExceeded {
	if o == nil {
		return nil
	}
	return o.RateLimitExceeded
}

func (o *CreateWorkspaceResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}
