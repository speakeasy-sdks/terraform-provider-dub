// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LeadsCount struct {
	// The total number of leads
	Leads float64 `json:"leads"`
}

func (o *LeadsCount) GetLeads() float64 {
	if o == nil {
		return 0.0
	}
	return o.Leads
}