// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dub/terraform-provider-dub/internal/sdk/internal/utils"
	"github.com/dub/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

type GetLinksGlobals struct {
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *GetLinksGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *GetLinksGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type QueryParamTagIdsType string

const (
	QueryParamTagIdsTypeStr        QueryParamTagIdsType = "str"
	QueryParamTagIdsTypeArrayOfStr QueryParamTagIdsType = "arrayOfStr"
)

// QueryParamTagIds - The tag IDs to filter the links by.
type QueryParamTagIds struct {
	Str        *string
	ArrayOfStr []string

	Type QueryParamTagIdsType
}

func CreateQueryParamTagIdsStr(str string) QueryParamTagIds {
	typ := QueryParamTagIdsTypeStr

	return QueryParamTagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateQueryParamTagIdsArrayOfStr(arrayOfStr []string) QueryParamTagIds {
	typ := QueryParamTagIdsTypeArrayOfStr

	return QueryParamTagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *QueryParamTagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = QueryParamTagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = QueryParamTagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for QueryParamTagIds", string(data))
}

func (u QueryParamTagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type QueryParamTagIds: all fields are null")
}

type QueryParamTagNamesType string

const (
	QueryParamTagNamesTypeStr        QueryParamTagNamesType = "str"
	QueryParamTagNamesTypeArrayOfStr QueryParamTagNamesType = "arrayOfStr"
)

// QueryParamTagNames - The unique name of the tags assigned to the short link (case insensitive).
type QueryParamTagNames struct {
	Str        *string
	ArrayOfStr []string

	Type QueryParamTagNamesType
}

func CreateQueryParamTagNamesStr(str string) QueryParamTagNames {
	typ := QueryParamTagNamesTypeStr

	return QueryParamTagNames{
		Str:  &str,
		Type: typ,
	}
}

func CreateQueryParamTagNamesArrayOfStr(arrayOfStr []string) QueryParamTagNames {
	typ := QueryParamTagNamesTypeArrayOfStr

	return QueryParamTagNames{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *QueryParamTagNames) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = QueryParamTagNamesTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = QueryParamTagNamesTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for QueryParamTagNames", string(data))
}

func (u QueryParamTagNames) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type QueryParamTagNames: all fields are null")
}

// Sort - The field to sort the links by. The default is `createdAt`, and sort order is always descending.
type Sort string

const (
	SortCreatedAt   Sort = "createdAt"
	SortClicks      Sort = "clicks"
	SortLastClicked Sort = "lastClicked"
)

func (e Sort) ToPointer() *Sort {
	return &e
}
func (e *Sort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "createdAt":
		fallthrough
	case "clicks":
		fallthrough
	case "lastClicked":
		*e = Sort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sort: %v", v)
	}
}

type GetLinksRequest struct {
	// The domain to filter the links by. E.g. `ac.me`. If not provided, all links for the workspace will be returned.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The tag ID to filter the links by. This field is deprecated – use `tagIds` instead.
	TagID *string `queryParam:"style=form,explode=true,name=tagId"`
	// The tag IDs to filter the links by.
	TagIds *QueryParamTagIds `queryParam:"style=form,explode=true,name=tagIds"`
	// The unique name of the tags assigned to the short link (case insensitive).
	TagNames *QueryParamTagNames `queryParam:"style=form,explode=true,name=tagNames"`
	// The search term to filter the links by. The search term will be matched against the short link slug and the destination url.
	Search *string `queryParam:"style=form,explode=true,name=search"`
	// The user ID to filter the links by.
	UserID *string `queryParam:"style=form,explode=true,name=userId"`
	// Whether to include archived links in the response. Defaults to `false` if not provided.
	ShowArchived *bool `default:"false" queryParam:"style=form,explode=true,name=showArchived"`
	// Whether to include tags in the response. Defaults to `false` if not provided.
	WithTags *bool `default:"false" queryParam:"style=form,explode=true,name=withTags"`
	// The field to sort the links by. The default is `createdAt`, and sort order is always descending.
	Sort *Sort `default:"createdAt" queryParam:"style=form,explode=true,name=sort"`
	// The page number for pagination (each page contains 100 links).
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
}

func (g GetLinksRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLinksRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLinksRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetLinksRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *GetLinksRequest) GetTagIds() *QueryParamTagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *GetLinksRequest) GetTagNames() *QueryParamTagNames {
	if o == nil {
		return nil
	}
	return o.TagNames
}

func (o *GetLinksRequest) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *GetLinksRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *GetLinksRequest) GetShowArchived() *bool {
	if o == nil {
		return nil
	}
	return o.ShowArchived
}

func (o *GetLinksRequest) GetWithTags() *bool {
	if o == nil {
		return nil
	}
	return o.WithTags
}

func (o *GetLinksRequest) GetSort() *Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetLinksRequest) GetPage() *float64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetLinksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A list of links
	LinkSchemas []shared.LinkSchema
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

func (o *GetLinksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLinksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLinksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLinksResponse) GetLinkSchemas() []shared.LinkSchema {
	if o == nil {
		return nil
	}
	return o.LinkSchemas
}

func (o *GetLinksResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *GetLinksResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *GetLinksResponse) GetForbidden() *shared.Forbidden {
	if o == nil {
		return nil
	}
	return o.Forbidden
}

func (o *GetLinksResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *GetLinksResponse) GetConflict() *shared.Conflict {
	if o == nil {
		return nil
	}
	return o.Conflict
}

func (o *GetLinksResponse) GetInviteExpired() *shared.InviteExpired {
	if o == nil {
		return nil
	}
	return o.InviteExpired
}

func (o *GetLinksResponse) GetUnprocessableEntity() *shared.UnprocessableEntity {
	if o == nil {
		return nil
	}
	return o.UnprocessableEntity
}

func (o *GetLinksResponse) GetRateLimitExceeded() *shared.RateLimitExceeded {
	if o == nil {
		return nil
	}
	return o.RateLimitExceeded
}

func (o *GetLinksResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}
