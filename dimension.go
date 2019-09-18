package epFulfillment

//Dimension is used within Product struct to designate a value and unit type for the dimension being specified
type Dimension struct {
	Value float64 `json:"value,omitempty"`
	Unit  string  `json:"unit,omitempty"`
}
