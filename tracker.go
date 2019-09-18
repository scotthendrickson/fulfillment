package epFulfillment

//Tracker holds all tracking details
type Tracker struct {
	ID              string           `json:"id,omitempty"`
	CreatedAt       int64            `json:"created_at,omitempty"`
	UpdatedAt       int64            `json:"updated_at,omitempty"`
	Mode            string           `json:"mode,omitempty"`
	DisabledAt      string           `json:"disabled_at,omitempty"`
	TrackingCode    string           `json:"tracking_code,omitempty"`
	Status          string           `json:"status,omitempty"`
	StatusDetail    string           `json:"status_detail,omitempty"`
	SignedBy        string           `json:"signed_by,omitempty"`
	Weight          int64            `json:"weight,omitempty"`
	EstDeliveryDate string           `json:"est_delivery_date,omitempty"`
	Carrier         string           `json:"carrier,omitempty"`
	TrackingDetails []TrackingDetail `json:"tracking_details,omitempty"`
}

//TrackingDetail ...
type TrackingDetail struct {
	Message          string           `json:"message,omitempty"`
	Description      string           `json:"description,omitempty"`
	Status           string           `json:"status,omitempty"`
	StatusDetail     string           `json:"status_detail,omitempty"`
	DateTime         string           `json:"datetime,omitempty"`
	Source           string           `json:"source,omitempty"`
	CarrierCode      string           `json:"carrier_code,omitempty"`
	TrackingLocation TrackingLocation `json:"tracking_location,omitempty"`
}

//TrackingLocation ...
type TrackingLocation struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Zip     string `json:"zip,omitempty"`
}
