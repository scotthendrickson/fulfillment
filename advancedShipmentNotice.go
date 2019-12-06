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
	Warehouse *Warehouse `json:"warehouse,omitempty"`
	//For creating an AdvancedShipmentNotice you must pass a warehouse id
	WarehouseID string `json:"warehouse_id,omitempty"`
}

//ASNList is for when you retrieve a list of ASN's on the account
type ASNList struct {
	AdvancedShipmentNotices []*AdvancedShipmentNotice `json:"advanced_shipment_notices,omitempty"`
}

//ListOptions is for specifying what options you want set on the get all request
type ListOptions struct {
	Limit   int64 `json:"limit,omitempty"`
	Offset  int64 `json:"offset,omitempty"`
	Page    int   `json:"page,omitempty"`
	PerPage int64 `json:"per_page,omitempty"`
}

//CreateASN is for creating an ASN with the EP fulfillment API
// client := epFulfillment.New("YOUR-API-KEY")
// asn, err := client.CreateASN(&epFulfillment.AdvancedShipmentNotice{
// 	Comments: "PO#123456",
// 	WarehouseID: "wh_303046cf435e416ba334618730dd7c2b",
// })
// if err != nil {
// 	fmt.Fprintln(os.Stderr, "error creating", err)
// 	os.Exit(1)
// 	return
// }
// fmt.Printf("%+v\n", asn)
func (c *Client) CreateASN(asn *AdvancedShipmentNotice) (a *AdvancedShipmentNotice, err error) {
	err = c.mainRequest("POST", "advanced_shipment_notices/", asn, &a)
	return
}

//UpdateASN is for updating an ASN with any necessary changes
// client := epFulfillment.New("YOUR-API-KEY")
// asn, err := client.GetASN("ADVANCED-SHIPMENT-NOTICE-ID")
// 	asn.Comments = "PO#555555"
// 	asn, err = client.UpdateASN(asn)
func (c *Client) UpdateASN(asn *AdvancedShipmentNotice) (a *AdvancedShipmentNotice, err error) {
	err = c.mainRequest("PATCH", "advanced_shipment_notices/"+asn.ID, asn, &a)
	return
}

//MarkASNComplete is for updating an ASN as completed
// client := epFulfillment.New("YOUR-API-KEY")
// asn, err := client.MarkASNComplete("ADVANCED-SHIPMENT-NOTICE-ID")
func (c *Client) MarkASNComplete(id string) (a *AdvancedShipmentNotice, err error) {
	err = c.mainRequest("PATCH", "advanced_shipment_notices/"+id+"/complete", nil, &a)
	return
}

//DeleteASN is used to delete an ASN
// client := epFulfillment.New("YOUR-API-KEY")
// err := client.DeleteASN("ADVANCED-SHIPMENT-NOTICE-ID")
func (c *Client) DeleteASN(id string) error {
	return c.mainRequest("DELETE", "advanced_shipment_notices/"+id, nil, nil)
}

//ListASNs will retrieve a list of all Advanced Shipment Notices on the account
// client := epFulfillment.New("YOUR-API-KEY")
// asnList, err := client.ListASNs(&epFulfillment.ListOptions{
// 		PerPage: 30,
// 	})
func (c *Client) ListASNs(in *ListOptions) (asnlist *ASNList, err error) {
	err = c.mainRequest("GET", "advanced_shipment_notices", in, &asnlist)
	return
}

//GetASN will retrieve a single ASN by id provided
// client := epFulfillment.New("YOUR-API-KEY")
// asn, err := client.GetASN("ADVANCED-SHIPMENT-NOTICE-ID")
func (c *Client) GetASN(id string) (a *AdvancedShipmentNotice, err error) {
	err = c.mainRequest("GET", "advanced_shipment_notices/"+id, nil, &a)
	return
}
