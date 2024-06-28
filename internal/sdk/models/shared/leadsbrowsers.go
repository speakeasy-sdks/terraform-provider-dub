// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LeadsBrowsers struct {
	// The name of the browser
	Browser string `json:"browser"`
	// The number of leads from this browser
	Leads float64 `json:"leads"`
}

func (o *LeadsBrowsers) GetBrowser() string {
	if o == nil {
		return ""
	}
	return o.Browser
}

func (o *LeadsBrowsers) GetLeads() float64 {
	if o == nil {
		return 0.0
	}
	return o.Leads
}
