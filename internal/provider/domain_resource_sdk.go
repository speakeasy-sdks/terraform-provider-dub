// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/dub/terraform-provider-dub/internal/sdk/models/operations"
	"github.com/dub/terraform-provider-dub/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DomainResourceModel) ToOperationsCreateDomainRequestBody() *operations.CreateDomainRequestBody {
	archived := new(bool)
	if !r.Archived.IsUnknown() && !r.Archived.IsNull() {
		*archived = r.Archived.ValueBool()
	} else {
		archived = nil
	}
	expiredURL := new(string)
	if !r.ExpiredURL.IsUnknown() && !r.ExpiredURL.IsNull() {
		*expiredURL = r.ExpiredURL.ValueString()
	} else {
		expiredURL = nil
	}
	placeholder := new(string)
	if !r.Placeholder.IsUnknown() && !r.Placeholder.IsNull() {
		*placeholder = r.Placeholder.ValueString()
	} else {
		placeholder = nil
	}
	slug := r.Slug.ValueString()
	out := operations.CreateDomainRequestBody{
		Archived:    archived,
		ExpiredURL:  expiredURL,
		Placeholder: placeholder,
		Slug:        slug,
	}
	return &out
}

func (r *DomainResourceModel) RefreshFromSharedDomainSchema(resp *shared.DomainSchema) {
	if resp != nil {
		r.Archived = types.BoolPointerValue(resp.Archived)
		r.CreatedAt = types.StringValue(resp.CreatedAt)
		r.ExpiredURL = types.StringPointerValue(resp.ExpiredURL)
		r.ID = types.StringValue(resp.ID)
		r.Placeholder = types.StringPointerValue(resp.Placeholder)
		r.Primary = types.BoolPointerValue(resp.Primary)
		r.Slug = types.StringValue(resp.Slug)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt)
		r.Verified = types.BoolPointerValue(resp.Verified)
	}
}

func (r *DomainResourceModel) ToOperationsUpdateDomainRequestBody() *operations.UpdateDomainRequestBody {
	archived := new(bool)
	if !r.Archived.IsUnknown() && !r.Archived.IsNull() {
		*archived = r.Archived.ValueBool()
	} else {
		archived = nil
	}
	expiredURL := new(string)
	if !r.ExpiredURL.IsUnknown() && !r.ExpiredURL.IsNull() {
		*expiredURL = r.ExpiredURL.ValueString()
	} else {
		expiredURL = nil
	}
	placeholder := new(string)
	if !r.Placeholder.IsUnknown() && !r.Placeholder.IsNull() {
		*placeholder = r.Placeholder.ValueString()
	} else {
		placeholder = nil
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	out := operations.UpdateDomainRequestBody{
		Archived:    archived,
		ExpiredURL:  expiredURL,
		Placeholder: placeholder,
		Slug:        slug,
	}
	return &out
}