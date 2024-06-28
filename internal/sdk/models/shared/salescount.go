// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SalesCount struct {
	// The total number of sales
	Sales float64 `json:"sales"`
	// The total amount of sales
	Amount float64 `json:"amount"`
}

func (o *SalesCount) GetSales() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sales
}

func (o *SalesCount) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}