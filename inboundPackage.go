package epFulfillment

import "fmt"

//InboundPackage is added to Advanced Shipment Notifications to indicate what products and in
//what quantities are being delivered to the warehouse
type InboundPackage struct {
	ID           string                   `json:"id,omitempty"`
	CreatedAt    int64                    `json:"created_at,omitempty"`
	UpdatedAt    int64                    `json:"updated_at,omitempty"`
	Mode         string                   `json:"mode,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Comments     string                   `json:"comments,omitempty"`
	TrackingCode string                   `json:"tracking_code,omitempty"`
	LineItems    []InboundPackageLineItem `json:"line_items,omitempty"`
}

//InboundPackageList when you retrieve a list of IP it goes here
type InboundPackageList struct {
	InboundPackages []InboundPackage `json:"inbound_packages,omitempty"`
}

//InboundPackageLineItem is to specify a product and line item being sent as well as what has been
//received by the warehouse.
type InboundPackageLineItem struct {
	ID            string  `json:"id,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
	UpdatedAt     int64   `json:"updated_at,omitempty"`
	Mode          string  `json:"mode,omitempty"`
	Product       Product `json:"product,omitempty"`
	Units         int64   `json:"units,omitempty"`
	ReceivedUnits int64   `json:"received_units,omitempty"`
}

//Create will take either an Advanced Shipment Notice ID or object and then create the inbound package
//attached to that particular Advanced Shipment Notice.
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// asn, err := epFulfillment.RetrieveASN("ADVANCED-SHIPMENT-NOTICE-ID")
// product, err := epFulfillment.RetrieveProduct("PRODUCT-ID")
// epLineItem := epFulfillment.InboundPackageLineItem{
// 	Product: product,
// 	Units:   10,
// }
// inboundPackage, err := epFulfillment.InboundPackage{
// 	Name:         "IP1",
// 	Comments:     "First set of Inbound Packages",
// 	TrackingCode: product.Title,
// 	LineItems:    []epFulfillment.InboundPackageLineItem{epLineItem},
// }.Create(asn)
// if err != nil {
// 	fmt.Fprintln(os.Stderr, "error creating", err)
// 	os.Exit(1)
// 	return
// }
// fmt.Printf("%+v\n", inboundPackage)
func (inboundPackage InboundPackage) Create(in interface{}) (ip InboundPackage, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("POST", "advanced_shipment_notices/"+in+"/inbound_packages/", inboundPackage, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("POST", "advanced_shipment_notices/"+in.ID+"/inbound_packages/", inboundPackage, &ip)
		return
	default:
		fmt.Print("Inbound Packages must be created as part of Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//Update this will take either a string or an ASN object and update the specific Inbound Package under it.
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// asn, err := epFulfillment.RetrieveASN("ADVANCED-SHIPMENT-NOTICE-ID")
// inboundPackage, err := epFulfillment.RetrieveInboundPackage(asn, "INBOUND-PACKAGE-ID")
// for i := range inboundPackage.LineItems {
// 	inboundPackage.LineItems[i].Units = inboundPackage.LineItems[i].Units + 5
// }
// inboundPackage.Update(asn)
//
// alternativly:
//
// inboundPackage, err := epFulfillment.RetrieveInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
// for i := range inboundPackage.LineItems {
// 	inboundPackage.LineItems[i].Units = inboundPackage.LineItems[i].Units + 5
// }
// inboundPackage.Update("ADVANCED-SHIPMENT-NOTICE-ID")
func (inboundPackage InboundPackage) Update(in interface{}) (ip InboundPackage, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("PATCH", "advanced_shipment_notices/"+in+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("PATCH", "advanced_shipment_notices/"+in.ID+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	default:
		fmt.Print("Inbound Packages must be updated using an Advanced Shipment Notice. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//Delete this will take either an ASN string id or an ASN object and delete the specific Inbound Package under it.
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// asn, err := epFulfillment.RetrieveASN("ADVANCED-SHIPMENT-NOTICE-ID")
// inboundPackage, err := epFulfillment.RetrieveInboundPackage(asn, "INBOUND-PACKAGE-ID")
// inboundPackage.Delete()
//
// alternativly:
//
// inboundPackage, err := epFulfillment.RetrieveInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
// inboundPackage.Delete()
func (inboundPackage InboundPackage) Delete(in interface{}) (ip InboundPackage, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("DELETE", "advanced_shipment_notices/"+in+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("DELETE", "advanced_shipment_notices/"+in.ID+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	default:
		fmt.Print("Inbound Packages must be updated using an Advanced Shipment Notice. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//DeleteInboundPackage this will take either a string or an ASN object as well as the Inbound Package ID and delete the specific Inbound Package.
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// asn, err := epFulfillment.RetrieveASN("ADVANCED-SHIPMENT-NOTICE-ID")
// inboundPackage, err := epFulfillment.DeleteInboundPackage(asn, "INBOUND-PACKAGE-ID")//
// alternativly:
//
// inboundPackage, err := epFulfillment.DeleteInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
func DeleteInboundPackage(asnIn interface{}, ipID string) (err error) {
	switch asnIn := asnIn.(type) {
	case string:
		return mainRequest("DELETE", "advanced_shipment_notices/"+asnIn+"/inbound_packages/"+ipID, nil, nil)
	case AdvancedShipmentNotice:
		return mainRequest("DELETE", "advanced_shipment_notices/"+asnIn.ID+"/inbound_packages/"+ipID, nil, nil)
	default:
		fmt.Print("Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//RetrieveAllInboundPackages this func will take either an ASN ID or ASN object and return a list of Inbound Packages on that ASN
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// asn, err := epFulfillment.RetrieveASN("ADVANCED-SHIPMENT-NOTICE-ID")
// inboundPackages, err := epFulfillment.RetrieveAllInboundPackages(asn)
//
// alternativly:
//
// inboundPackages, err := epFulfillment.RetrieveAllInboundPackages("ADVANCED-SHIPMENT-NOTICE-ID")
func RetrieveAllInboundPackages(in interface{}) (ipl InboundPackageList, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("GET", "advanced_shipment_notices/"+in+"/inbound_packages/", nil, &ipl)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("GET", "advanced_shipment_notices/"+in.ID+"/inbound_packages/", nil, &ipl)
		return
	default:
		fmt.Print("Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//RetrieveInboundPackage this func will take either an ASN ID or ASN object and return the Inbound Package object for that ID
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// asn, err := epFulfillment.RetrieveASN("ADVANCED-SHIPMENT-NOTICE-ID")
// inboundPackage, err := epFulfillment.RetrieveInboundPackage(asn, "INBOUND-PACKAGE-ID")
//
// alternativly:
//
// inboundPackage, err := epFulfillment.RetrieveInboundPackage("ADVANCED-SHIPMENT-NOTICE-ID", "INBOUND-PACKAGE-ID")
func RetrieveInboundPackage(asnIn interface{}, ipID string) (ip InboundPackage, err error) {
	switch asnIn := asnIn.(type) {
	case string:
		err = mainRequest("GET", "advanced_shipment_notices/"+asnIn+"/inbound_packages/"+ipID, nil, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("GET", "advanced_shipment_notices/"+asnIn.ID+"/inbound_packages/"+ipID, nil, &ip)
		return
	default:
		fmt.Print("Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}
