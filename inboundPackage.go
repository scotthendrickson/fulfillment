package epFulfillment

//InboundPackage is added to Advanced Shipment Notifications to indicate what products and in
//what quantities are being delivered to the warehouse
type InboundPackage struct {
	ID           string                    `json:"id,omitempty"`
	CreatedAt    int64                     `json:"created_at,omitempty"`
	UpdatedAt    int64                     `json:"updated_at,omitempty"`
	Mode         string                    `json:"mode,omitempty"`
	Name         string                    `json:"name,omitempty"`
	Comments     string                    `json:"comments,omitempty"`
	TrackingCode string                    `json:"tracking_code,omitempty"`
	LineItems    []*InboundPackageLineItem `json:"line_items,omitempty"`
}

//InboundPackageList when you retrieve a list of IP it goes here
type InboundPackageList struct {
	InboundPackages []*InboundPackage `json:"inbound_packages,omitempty"`
}

//InboundPackageLineItem is to specify a product and line item being sent as well as what has been
//received by the warehouse.
type InboundPackageLineItem struct {
	ID            string   `json:"id,omitempty"`
	CreatedAt     int64    `json:"created_at,omitempty"`
	UpdatedAt     int64    `json:"updated_at,omitempty"`
	Mode          string   `json:"mode,omitempty"`
	Product       *Product `json:"product,omitempty"`
	Units         int64    `json:"units,omitempty"`
	ReceivedUnits int64    `json:"received_units,omitempty"`
}

//CreateInboundPackage will take an Advanced Shipment Notice ID and then create the inbound package
//attached to that particular Advanced Shipment Notice.
// client := epFulfillment.New("YOUR-API-KEY")
// asn, err := client.GetASN("ADVANCED-SHIPMENT-NOTICE-ID")
// product, err := client.GetProduct("PRODUCT-ID")
// epLineItem := &epFulfillment.InboundPackageLineItem{
// 		Product: product,
// 		Units:   121,
// 	}
// 	inboundPackage, err := client.CreateInboundPackage(asn.ID, &epFulfillment.InboundPackage{
// 		Name:         "IP1",
// 		Comments:     "First set of Inbound Packages",
// 		TrackingCode: product.Title,
// 		LineItems:    []*epFulfillment.InboundPackageLineItem{epLineItem},
// 	})
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "error creating", err)
// 		os.Exit(1)
// 		return
// 	}
// 	fmt.Printf("%+v\n", inboundPackage)
func (c *Client) CreateInboundPackage(asnID string, inboundPackage *InboundPackage) (ip InboundPackage, err error) {
	err = c.mainRequest("POST", "advanced_shipment_notices/"+asnID+"/inbound_packages/", inboundPackage, &ip)
	return
}

//UpdateInboundPackage this will take an Advanced Shipment Notice ID as well as the Inbound package and update the Inbound Package.
// client := epFulfillment.New("YOUR-API-KEY")
// inboundPackage, err := client.GetInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
// for i := range inboundPackage.LineItems {
// 		inboundPackage.LineItems[i].Units = inboundPackage.LineItems[i].Units + 5
// 	}
// 	inboundPackage, err = client.UpdateInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", &inboundPackage)
func (c *Client) UpdateInboundPackage(asnID string, inboundPackage *InboundPackage) (ip InboundPackage, err error) {
	err = c.mainRequest("PATCH", "advanced_shipment_notices/"+asnID+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
	return
}

//DeleteInboundPackage this will take an Advanced Shipment Notice ID as well as the Inbound Package ID and delete the specific Inbound Package.
// client := epFulfillment.New("YOUR-API-KEY")
// err := client.DeleteInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
func (c *Client) DeleteInboundPackage(asnID string, inboundPackage string) (err error) {
	return c.mainRequest("DELETE", "advanced_shipment_notices/"+asnID+"/inbound_packages/"+inboundPackage, nil, nil)
}

//ListInboundPackages this func (c *Client) will an Advanced Shipment Notice ID and then return a list of Inbound Packages on that ASN
// client := epFulfillment.New("YOUR-API-KEY")
// inboundPackages, err := client.ListInboundPackages("ADVANCED-SHIPMENT-NOTICE-ID")
// 	for i := range inboundPackages.InboundPackages {
// 		fmt.Printf("%+v\n", inboundPackages.InboundPackages[i].ID)
// 	}
func (c *Client) ListInboundPackages(asnID string) (ipl InboundPackageList, err error) {
	err = c.mainRequest("GET", "advanced_shipment_notices/"+asnID+"/inbound_packages/", nil, &ipl)
	return
}

//GetInboundPackage this func (c *Client) will an Advanced Shipment Notice ID and then return the Inbound Package object for that ID
// client := epFulfillment.New("YOUR-API-KEY")
// inboundPackage, err := client.GetInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
func (c *Client) GetInboundPackage(asnID string, ipID string) (ip InboundPackage, err error) {
	err = c.mainRequest("GET", "advanced_shipment_notices/"+asnID+"/inbound_packages/"+ipID, nil, &ip)
	return
}
