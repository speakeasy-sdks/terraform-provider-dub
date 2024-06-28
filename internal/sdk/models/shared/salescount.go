// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SalesCount struct {
	// The total amount of sales
	Amount float64 `json:"amount"`
	// The total number of sales
	Sales float64 `json:"sales"`
}

func (o *SalesCount) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *SalesCount) GetSales() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sales
}
