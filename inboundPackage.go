package epFulfillment

import "fmt"

//InboundPackage ...
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

//InboundPackageLineItem ...
type InboundPackageLineItem struct {
	ID            string  `json:"id,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
	UpdatedAt     int64   `json:"updated_at,omitempty"`
	Mode          string  `json:"mode,omitempty"`
	Product       Product `json:"product,omitempty"`
	Units         int64   `json:"units,omitempty"`
	ReceivedUnits int64   `json:"received_units,omitempty"`
}

//CreateWithASNID will take an ASN's id as a string and create the Inbound Package
func (inboundPackage InboundPackage) CreateWithASNID(asnID string) (ip InboundPackage, err error) {
	err = mainRequest("POST", "advanced_shipment_notices/"+asnID+"/inbound_packages/", inboundPackage, &ip)
	return
}

//CreateWithASN will take an ASN and create the Inbound Package
func (inboundPackage InboundPackage) CreateWithASN(asn AdvancedShipmentNotice) (ip InboundPackage, err error) {
	err = mainRequest("POST", "advanced_shipment_notices/"+asn.ID+"/inbound_packages/", inboundPackage, &ip)
	return
}

//Create ideally this should take either a string or an ASN and create the Inbound Package under it.
func (inboundPackage InboundPackage) Create(in interface{}) (ip InboundPackage, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("POST", "advanced_shipment_notices/"+in+"/inbound_packages/", inboundPackage, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("POST", "advanced_shipment_notices/"+in.ID+"/inbound_packages/", inboundPackage, &ip)
		return
	case nil:
		fmt.Println("Nothing was provided. Inbound Packages must be created as part of Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	default:
		fmt.Print("Inbound Packages must be created as part of Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//Update ideally this should take either a string or an ASN and update the specific Inbound Package under it.
func (inboundPackage InboundPackage) Update(in interface{}) (ip InboundPackage, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("PATCH", "advanced_shipment_notices/"+in+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("PATCH", "advanced_shipment_notices/"+in.ID+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	case nil:
		fmt.Println("Nothing was provided. Inbound Packages must be updated using an Advanced Shipment Notice. Please provide either an ID or Advanced Shipment Notice to continue.")
	default:
		fmt.Print("Inbound Packages must be updated using an Advanced Shipment Notice. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//Delete ideally this should take either a string or an ASN and delete the specific Inbound Package under it.
func (inboundPackage InboundPackage) Delete(in interface{}) (ip InboundPackage, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("DELETE", "advanced_shipment_notices/"+in+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("DELETE", "advanced_shipment_notices/"+in.ID+"/inbound_packages/"+inboundPackage.ID, inboundPackage, &ip)
		return
	case nil:
		fmt.Println("Nothing was provided. Inbound Packages must be updated using an Advanced Shipment Notice. Please provide either an ID or Advanced Shipment Notice to continue.")
	default:
		fmt.Print("Inbound Packages must be updated using an Advanced Shipment Notice. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//DeleteIP ideally this should take either a string or an ASN and delete the specific Inbound Package under it.
func DeleteIP(asnIn interface{}, ipID string) (err error) {
	switch asnIn := asnIn.(type) {
	case string:
		return mainRequest("DELETE", "advanced_shipment_notices/"+asnIn+"/inbound_packages/"+ipID, nil, nil)
	case AdvancedShipmentNotice:
		return mainRequest("DELETE", "advanced_shipment_notices/"+asnIn.ID+"/inbound_packages/"+ipID, nil, nil)
	case nil:
		fmt.Println("Nothing was provided. Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	default:
		fmt.Print("Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//RetrieveAllIP this func will take either an ASN ID or ASN object and return a list of Inbound Packages on that ASN
func RetrieveAllIP(in interface{}) (ipl InboundPackageList, err error) {
	switch in := in.(type) {
	case string:
		err = mainRequest("GET", "advanced_shipment_notices/"+in+"/inbound_packages/", nil, &ipl)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("GET", "advanced_shipment_notices/"+in.ID+"/inbound_packages/", nil, &ipl)
		return
	case nil:
		fmt.Println("Nothing was provided. Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	default:
		fmt.Print("Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}

//RetrieveIP this func will take either an ASN ID or ASN object and return a list of Inbound Packages on that ASN
func RetrieveIP(asnIn interface{}, ipID string) (ip InboundPackage, err error) {
	switch asnIn := asnIn.(type) {
	case string:
		err = mainRequest("GET", "advanced_shipment_notices/"+asnIn+"/inbound_packages/"+ipID, nil, &ip)
		return
	case AdvancedShipmentNotice:
		err = mainRequest("GET", "advanced_shipment_notices/"+asnIn.ID+"/inbound_packages/"+ipID, nil, &ip)
		return
	case nil:
		fmt.Println("Nothing was provided. Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	default:
		fmt.Print("Inbound Packages must be retrieved using an Advanced Shipment Notices. Please provide either an ID or Advanced Shipment Notice to continue.")
	}
	return
}
