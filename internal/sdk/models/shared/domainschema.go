// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-dub/internal/sdk/internal/utils"
)

type DomainSchema struct {
	// The unique identifier of the domain.
	ID string `json:"id"`
	// The domain name.
	Slug string `json:"slug"`
	// Whether the domain is verified.
	Verified *bool `default:"false" json:"verified"`
	// Whether the domain is the primary domain for the workspace.
	Primary *bool `default:"false" json:"primary"`
	// Whether the domain is archived.
	Archived *bool `default:"false" json:"archived"`
	// Provide context to your teammates in the link creation modal by showing them an example of a link to be shortened.
	Placeholder *string `default:"https://dub.co/help/article/what-is-dub" json:"placeholder"`
	// The URL to redirect to when a link under this domain has expired.
	ExpiredURL *string `json:"expiredUrl"`
	// The date the domain was created.
	CreatedAt string `json:"createdAt"`
	// The date the domain was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (d DomainSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DomainSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DomainSchema) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DomainSchema) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *DomainSchema) GetVerified() *bool {
	if o == nil {
		return nil
	}
	return o.Verified
}

func (o *DomainSchema) GetPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.Primary
}

func (o *DomainSchema) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *DomainSchema) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *DomainSchema) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *DomainSchema) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DomainSchema) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}
