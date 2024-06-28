// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ClicksReferers struct {
	// The name of the referer. If unknown, this will be `(direct)`
	Referer string `json:"referer"`
	// The number of clicks from this referer
	Clicks float64 `json:"clicks"`
}

func (o *ClicksReferers) GetReferer() string {
	if o == nil {
		return ""
	}
	return o.Referer
}

func (o *ClicksReferers) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}