// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dub/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

type DeleteLinkGlobals struct {
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *DeleteLinkGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *DeleteLinkGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type DeleteLinkRequest struct {
	// The id of the link to delete. You may use either `linkId` (obtained via `/links/info` endpoint) or `externalId` prefixed with `ext_`.
	LinkID string `pathParam:"style=simple,explode=false,name=linkId"`
}

func (o *DeleteLinkRequest) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

// DeleteLinkResponseBody - The deleted link
type DeleteLinkResponseBody struct {
	// The ID of the link.
	ID string `json:"id"`
}

func (o *DeleteLinkResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The deleted link
	Object *DeleteLinkResponseBody
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

func (o *DeleteLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteLinkResponse) GetObject() *DeleteLinkResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteLinkResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *DeleteLinkResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *DeleteLinkResponse) GetForbidden() *shared.Forbidden {
	if o == nil {
		return nil
	}
	return o.Forbidden
}

func (o *DeleteLinkResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *DeleteLinkResponse) GetConflict() *shared.Conflict {
	if o == nil {
		return nil
	}
	return o.Conflict
}

func (o *DeleteLinkResponse) GetInviteExpired() *shared.InviteExpired {
	if o == nil {
		return nil
	}
	return o.InviteExpired
}

func (o *DeleteLinkResponse) GetUnprocessableEntity() *shared.UnprocessableEntity {
	if o == nil {
		return nil
	}
	return o.UnprocessableEntity
}

func (o *DeleteLinkResponse) GetRateLimitExceeded() *shared.RateLimitExceeded {
	if o == nil {
		return nil
	}
	return o.RateLimitExceeded
}

func (o *DeleteLinkResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}
