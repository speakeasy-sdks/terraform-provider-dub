// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-dub/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-dub/internal/sdk/models/shared"
	"net/http"
)

type UpdateLinkGlobals struct {
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *UpdateLinkGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *UpdateLinkGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type UpdateLinkTagIdsType string

const (
	UpdateLinkTagIdsTypeStr        UpdateLinkTagIdsType = "str"
	UpdateLinkTagIdsTypeArrayOfStr UpdateLinkTagIdsType = "arrayOfStr"
)

// UpdateLinkTagIds - The unique IDs of the tags assigned to the short link.
type UpdateLinkTagIds struct {
	Str        *string
	ArrayOfStr []string

	Type UpdateLinkTagIdsType
}

func CreateUpdateLinkTagIdsStr(str string) UpdateLinkTagIds {
	typ := UpdateLinkTagIdsTypeStr

	return UpdateLinkTagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpdateLinkTagIdsArrayOfStr(arrayOfStr []string) UpdateLinkTagIds {
	typ := UpdateLinkTagIdsTypeArrayOfStr

	return UpdateLinkTagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *UpdateLinkTagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpdateLinkTagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = UpdateLinkTagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateLinkTagIds", string(data))
}

func (u UpdateLinkTagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateLinkTagIds: all fields are null")
}

type UpdateLinkTagNamesType string

const (
	UpdateLinkTagNamesTypeStr        UpdateLinkTagNamesType = "str"
	UpdateLinkTagNamesTypeArrayOfStr UpdateLinkTagNamesType = "arrayOfStr"
)

// UpdateLinkTagNames - The unique name of the tags assigned to the short link (case insensitive).
type UpdateLinkTagNames struct {
	Str        *string
	ArrayOfStr []string

	Type UpdateLinkTagNamesType
}

func CreateUpdateLinkTagNamesStr(str string) UpdateLinkTagNames {
	typ := UpdateLinkTagNamesTypeStr

	return UpdateLinkTagNames{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpdateLinkTagNamesArrayOfStr(arrayOfStr []string) UpdateLinkTagNames {
	typ := UpdateLinkTagNamesTypeArrayOfStr

	return UpdateLinkTagNames{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *UpdateLinkTagNames) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpdateLinkTagNamesTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = UpdateLinkTagNamesTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateLinkTagNames", string(data))
}

func (u UpdateLinkTagNames) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateLinkTagNames: all fields are null")
}

type UpdateLinkRequestBody struct {
	// The destination URL of the short link.
	URL string `json:"url"`
	// The domain of the short link. If not provided, the primary domain for the workspace will be used (or `dub.sh` if the workspace has no domains).
	Domain *string `json:"domain,omitempty"`
	// The short link slug. If not provided, a random 7-character slug will be generated.
	Key *string `json:"key,omitempty"`
	// This is the ID of the link in your database. If set, it can be used to identify the link in the future. Must be prefixed with `ext_` when passed as a query parameter.
	ExternalID *string `json:"externalId,omitempty"`
	// The prefix of the short link slug for randomly-generated keys (e.g. if prefix is `/c/`, generated keys will be in the `/c/:key` format). Will be ignored if `key` is provided.
	Prefix *string `json:"prefix,omitempty"`
	// Whether to track conversions for the short link.
	TrackConversion *bool `default:"false" json:"trackConversion"`
	// Whether the short link is archived.
	Archived *bool `default:"false" json:"archived"`
	// Whether the short link's stats are publicly accessible.
	PublicStats *bool `default:"false" json:"publicStats"`
	// The unique ID of the tag assigned to the short link. This field is deprecated – use `tagIds` instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TagID *string `json:"tagId,omitempty"`
	// The unique IDs of the tags assigned to the short link.
	TagIds *UpdateLinkTagIds `json:"tagIds,omitempty"`
	// The unique name of the tags assigned to the short link (case insensitive).
	TagNames *UpdateLinkTagNames `json:"tagNames,omitempty"`
	// The comments for the short link.
	Comments *string `json:"comments,omitempty"`
	// The date and time when the short link will expire at.
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// The URL to redirect to when the short link has expired.
	ExpiredURL *string `json:"expiredUrl,omitempty"`
	// The password required to access the destination URL of the short link.
	Password *string `json:"password,omitempty"`
	// Whether the short link uses Custom Social Media Cards feature.
	Proxy *bool `default:"false" json:"proxy"`
	// The title of the short link generated via `api.dub.co/metatags`. Will be used for Custom Social Media Cards if `proxy` is true.
	Title *string `json:"title,omitempty"`
	// The description of the short link generated via `api.dub.co/metatags`. Will be used for Custom Social Media Cards if `proxy` is true.
	Description *string `json:"description,omitempty"`
	// The image of the short link generated via `api.dub.co/metatags`. Will be used for Custom Social Media Cards if `proxy` is true.
	Image *string `json:"image,omitempty"`
	// Whether the short link uses link cloaking.
	Rewrite *bool `default:"false" json:"rewrite"`
	// The iOS destination URL for the short link for iOS device targeting.
	Ios *string `json:"ios,omitempty"`
	// The Android destination URL for the short link for Android device targeting.
	Android *string `json:"android,omitempty"`
	// Geo targeting information for the short link in JSON format `{[COUNTRY]: https://example.com }`.
	Geo *shared.LinkGeoTargeting `json:"geo,omitempty"`
	// Allow search engines to index your short link. Defaults to `false` if not provided. Learn more: https://d.to/noindex
	DoIndex *bool `default:"false" json:"doIndex"`
}

func (u UpdateLinkRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateLinkRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateLinkRequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *UpdateLinkRequestBody) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *UpdateLinkRequestBody) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *UpdateLinkRequestBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *UpdateLinkRequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *UpdateLinkRequestBody) GetTrackConversion() *bool {
	if o == nil {
		return nil
	}
	return o.TrackConversion
}

func (o *UpdateLinkRequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *UpdateLinkRequestBody) GetPublicStats() *bool {
	if o == nil {
		return nil
	}
	return o.PublicStats
}

func (o *UpdateLinkRequestBody) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *UpdateLinkRequestBody) GetTagIds() *UpdateLinkTagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *UpdateLinkRequestBody) GetTagNames() *UpdateLinkTagNames {
	if o == nil {
		return nil
	}
	return o.TagNames
}

func (o *UpdateLinkRequestBody) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *UpdateLinkRequestBody) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *UpdateLinkRequestBody) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *UpdateLinkRequestBody) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *UpdateLinkRequestBody) GetProxy() *bool {
	if o == nil {
		return nil
	}
	return o.Proxy
}

func (o *UpdateLinkRequestBody) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UpdateLinkRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateLinkRequestBody) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *UpdateLinkRequestBody) GetRewrite() *bool {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *UpdateLinkRequestBody) GetIos() *string {
	if o == nil {
		return nil
	}
	return o.Ios
}

func (o *UpdateLinkRequestBody) GetAndroid() *string {
	if o == nil {
		return nil
	}
	return o.Android
}

func (o *UpdateLinkRequestBody) GetGeo() *shared.LinkGeoTargeting {
	if o == nil {
		return nil
	}
	return o.Geo
}

func (o *UpdateLinkRequestBody) GetDoIndex() *bool {
	if o == nil {
		return nil
	}
	return o.DoIndex
}

type UpdateLinkRequest struct {
	// The id of the link to update. You may use either `linkId` (obtained via `/links/info` endpoint) or `externalId` prefixed with `ext_`.
	LinkID      string                 `pathParam:"style=simple,explode=false,name=linkId"`
	RequestBody *UpdateLinkRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateLinkRequest) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

func (o *UpdateLinkRequest) GetRequestBody() *UpdateLinkRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The updated link
	LinkSchema *shared.LinkSchema
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

func (o *UpdateLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLinkResponse) GetLinkSchema() *shared.LinkSchema {
	if o == nil {
		return nil
	}
	return o.LinkSchema
}

func (o *UpdateLinkResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateLinkResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateLinkResponse) GetForbidden() *shared.Forbidden {
	if o == nil {
		return nil
	}
	return o.Forbidden
}

func (o *UpdateLinkResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateLinkResponse) GetConflict() *shared.Conflict {
	if o == nil {
		return nil
	}
	return o.Conflict
}

func (o *UpdateLinkResponse) GetInviteExpired() *shared.InviteExpired {
	if o == nil {
		return nil
	}
	return o.InviteExpired
}

func (o *UpdateLinkResponse) GetUnprocessableEntity() *shared.UnprocessableEntity {
	if o == nil {
		return nil
	}
	return o.UnprocessableEntity
}

func (o *UpdateLinkResponse) GetRateLimitExceeded() *shared.RateLimitExceeded {
	if o == nil {
		return nil
	}
	return o.RateLimitExceeded
}

func (o *UpdateLinkResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}