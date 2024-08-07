// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LeadsOS struct {
	// The number of leads from this OS
	Leads float64 `json:"leads"`
	// The name of the OS
	Os string `json:"os"`
}

func (o *LeadsOS) GetLeads() float64 {
	if o == nil {
		return 0.0
	}
	return o.Leads
}

func (o *LeadsOS) GetOs() string {
	if o == nil {
		return ""
	}
	return o.Os
}
