package epFulfillment

//AdvancedShipmentNotice helps us track inbound shipments to for inventory replenishment.
type AdvancedShipmentNotice struct {
	ID                string `json:"id,omitempty"`
	CreatedAt         int64  `json:"created_at,omitempty"`
	UpdatedAt         int64  `json:"updated_at,omitempty"`
	Mode              string `json:"mode,omitempty"`
	EstimatedDelivery string `json:"estimated_delivery,omitempty"`
	Comments          string `json:"comments,omitempty"`
	URL               string `json:"url,omitempty"`
	Status            string `json:"status,omitempty"`
	//On the retrun response you'll get a warehouse obj
	Warehouse Warehouse `json:"warehouse,omitempty"`
	//For creating an AdvancedShipmentNotice you can pass an id instead of a warehouse obj
	WarehouseID string `json:"warehouse_id,omitempty"`
}

//ASNList is for when you retrieve a list of ASN's on the account
type ASNList struct {
	AdvancedShipmentNotices []AdvancedShipmentNotice `json:"advanced_shipment_notices,omitempty"`
}

//ASNListOptions is for specifying what options you want set on the get all request
type ASNListOptions struct {
	Limit   int64 `json:"limit,omitempty"`
	Offset  int64 `json:"offset,omitempty"`
	Page    int64 `json:"page,omitempty"`
	PerPage int64 `json:"per_page,omitempty"`
}

//Create is for creating an ASN with the EP fulfillment API
func (asn AdvancedShipmentNotice) Create() (a AdvancedShipmentNotice, err error) {
	err = mainRequest("POST", "advanced_shipment_notices/", asn, &a)
	return
}

//Update is for updating an ASN with any necessary changes
func (asn AdvancedShipmentNotice) Update() (a AdvancedShipmentNotice, err error) {
	err = mainRequest("PATCH", "advanced_shipment_notices/"+asn.ID, asn, &a)
	return
}

//MarkComplete is for updating an ASN as completed
func (asn AdvancedShipmentNotice) MarkComplete() (a AdvancedShipmentNotice, err error) {
	err = mainRequest("PATCH", "advanced_shipment_notices/"+asn.ID+"/complete", asn, &a)
	return
}

//Delete is used to delete an ASN
func (asn AdvancedShipmentNotice) Delete() error {
	return mainRequest("DELETE", "advanced_shipment_notices/"+asn.ID, nil, nil)
}

//DeleteASN is used to delete an ASN
func DeleteASN(id string) error {
	return mainRequest("DELETE", "advanced_shipment_notices/"+id, nil, nil)
}

//RetrieveAllASN will retrieve a list of all Advanced Shipment Notices on the account
func RetrieveAllASN(in ASNListOptions) (asnlist ASNList, err error) {
	err = mainRequest("GET", "advanced_shipment_notices", in, &asnlist)
	return
}

//RetrieveASN will retrieve a single ASN by id provided
func RetrieveASN(id string) (a AdvancedShipmentNotice, err error) {
	err = mainRequest("GET", "advanced_shipment_notices/"+id, nil, &a)
	return
}
