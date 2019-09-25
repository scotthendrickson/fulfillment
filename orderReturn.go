package epFulfillment

//OrderReturn describes the origin, the destination and the set of returning goods from original order to
//be mailed to that destination
type OrderReturn struct {
	ID                 string                `json:"id,omitempty"`
	CreatedAt          int64                 `json:"created_at,omitempty"`
	UpdatedAt          int64                 `json:"updated_at,omitempty"`
	Mode               string                `json:"mode,omitempty"`
	LineItems          []OrderReturnLineItem `json:"line_items,omitempty"`
	DestinationAddress Address               `json:"destination_address,omitempty"`
	OriginAddress      Address               `json:"origin_address,omitempty"`
	LabelURL           string                `json:"label_url,omitempty"`
	TrackingCode       string                `json:"tracking_code,omitempty"`
	Tracker            Tracker               `json:"tracker,omitempty"`
	Status             string                `json:"status,omitempty"`
	Order              Order                 `json:"order,omitempty"`
	OrderID            string                `json:"order_id,omitempty"`
}

//OrderReturnList is for a list of order returns
type OrderReturnList struct {
	OrderReturns []OrderReturn `json:"order_returns,omitempty"`
}

//OrderReturnLineItem each product to be on the return
type OrderReturnLineItem struct {
	ID         string  `json:"id,omitempty"`
	CreatedAt  int64   `json:"created_at,omitempty"`
	UpdatedAt  int64   `json:"updated_at,omitempty"`
	Mode       string  `json:"mode,omitempty"`
	Product    Product `json:"product,omitempty"`
	ProductID  string  `json:"product_id,omitempty"`
	Units      int64   `json:"units,omitempty"`
	ReasonType string  `json:"reason_type,omitempty"`
}

//Create allows for the creation of an order return
//One thing to note is that both Order ID and LineItems are required by the EasyPostAPI
//epFulfillment.SetAPIKey("YOUR-API-KEY")
// warehouses, err := epFulfillment.RetrieveAllWarehouses()
// warehouse := epFulfillment.Warehouse{}
// for i := range warehouses.Warehouses {
// 	if warehouses.Warehouses[i].ID == "WAREHOUSE-ID" {
// 		warehouse = warehouses.Warehouses[i]
// 	}
// }
// order, err := epFulfillment.RetrieveOrder("ORDER-ID")
// orderReturnLineItem := epFulfillment.OrderReturnLineItem{
// 	ProductID:  product.ID,
// 	ReasonType: "wrong_size",
// 	Units:      1,
// }
// orderReturn := epFulfillment.OrderReturn{
// 	LineItems:          []epFulfillment.OrderReturnLineItem{orderReturnLineItem},
// 	DestinationAddress: warehouse.Address,
// 	OriginAddress:      order.Destination,
//  Order:              order,
////alternatively here you can just use the Order ID instead of the order object
// 	OrderID: order.ID,
// }.Create()
func (orderReturn OrderReturn) Create() (or OrderReturn, err error) {
	err = mainRequest("POST", "order_returns", orderReturn, &or)
	return
}

//Delete is an interface for deleting an orderReturn
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//orderReturn, err := epFulfillment.RetrieveOrderReturn("ORDER-RETURN-ID")
//orderReturn.Delete()
func (orderReturn OrderReturn) Delete() error {
	return mainRequest("DELETE", "order_returns/"+orderReturn.ID, nil, nil)
}

//DeleteOrderReturn is for deleting an Order return by provided id
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//epFulfillment.DeleteOrderReturn("ORDER-RETURN-ID")
func DeleteOrderReturn(id string) error {
	return mainRequest("DELETE", "order_returns/"+id, nil, nil)
}

//RetrieveAllOrderReturns will return a list of all your order returns
//epFulfillment.SetAPIKey("YOUR-API-KEY")
// listOptions := epFulfillment.ListOptions{
// 	Limit:   10,
// 	Offset:  0,
// 	Page:    1,
// 	PerPage: 2,
// }
// orderReturns, err := epFulfillment.RetrieveAllOrderReturns(listOptions)
func RetrieveAllOrderReturns(in ListOptions) (or OrderReturnList, err error) {
	err = mainRequest("GET", "order_returns", in, &or)
	return
}

//RetrieveOrderReturn will retrieve a single OR by id provided
//epFulfillment.SetAPIKey("YOUR-API-KEY")
// orderReturn, err := epFulfillment.RetrieveOrderReturn("ORDER-RETURN-ID")
func RetrieveOrderReturn(id string) (or OrderReturn, err error) {
	err = mainRequest("GET", "order_returns/"+id, nil, &or)
	return
}
