// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LeadsTimeseries struct {
	// The starting timestamp of the interval
	Start string `json:"start"`
	// The number of leads in the interval
	Leads float64 `json:"leads"`
}

func (o *LeadsTimeseries) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *LeadsTimeseries) GetLeads() float64 {
	if o == nil {
		return 0.0
	}
	return o.Leads
}
