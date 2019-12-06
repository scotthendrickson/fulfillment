package epFulfillment

//OrderReturn describes the origin, the destination and the set of returning goods from original order to
//be mailed to that destination
type OrderReturn struct {
	ID                 string                 `json:"id,omitempty"`
	CreatedAt          int64                  `json:"created_at,omitempty"`
	UpdatedAt          int64                  `json:"updated_at,omitempty"`
	Mode               string                 `json:"mode,omitempty"`
	LineItems          []*OrderReturnLineItem `json:"line_items,omitempty"`
	DestinationAddress *Address               `json:"destination_address,omitempty"`
	OriginAddress      *Address               `json:"origin_address,omitempty"`
	LabelURL           string                 `json:"label_url,omitempty"`
	TrackingCode       string                 `json:"tracking_code,omitempty"`
	Tracker            *Tracker               `json:"tracker,omitempty"`
	Status             string                 `json:"status,omitempty"`
	Order              *Order                 `json:"order,omitempty"`
	OrderID            string                 `json:"order_id,omitempty"`
}

//OrderReturnList is for a list of order returns
type OrderReturnList struct {
	OrderReturns []*OrderReturn `json:"order_returns,omitempty"`
}

//OrderReturnLineItem each product to be on the return
type OrderReturnLineItem struct {
	ID         string   `json:"id,omitempty"`
	CreatedAt  int64    `json:"created_at,omitempty"`
	UpdatedAt  int64    `json:"updated_at,omitempty"`
	Mode       string   `json:"mode,omitempty"`
	Product    *Product `json:"product,omitempty"`
	ProductID  string   `json:"product_id,omitempty"`
	Units      int64    `json:"units,omitempty"`
	ReasonType string   `json:"reason_type,omitempty"`
}

//CreateOrderReturn allows for the creation of an order return
//One thing to note is that both Order ID and LineItems are required by the EasyPostAPI
//client := epFulfillment.New("YOUR-API-KEY")
// warehouses, err := client.ListWarehouses()
// warehouse := &epFulfillment.Warehouse{}
// for i := range warehouses.Warehouses {
// 	if warehouses.Warehouses[i].ID == "WAREHOUSE-ID" {
// 		warehouse = warehouses.Warehouses[i]
// 	}
// }
// orderReturn, err := client.CreateOrderReturn(&epFulfillment.OrderReturn{
// 	LineItems: []*epFulfillment.OrderReturnLineItem{&epFulfillment.OrderReturnLineItem{
// 		ProductID:  "PRODUCT-ID",
// 		ReasonType: "wrong_size",
// 		Units:      1,
// 	}},
// 	DestinationAddress: warehouse.Address,
// 	OrderID: "ORDER-ID",
// })
func (c *Client) CreateOrderReturn(orderReturn *OrderReturn) (or *OrderReturn, err error) {
	err = c.mainRequest("POST", "order_returns", orderReturn, &or)
	return
}

//DeleteOrderReturn is for deleting an Order return by provided id
//client := epFulfillment.New("YOUR-API-KEY")
//epFulfillment.DeleteOrderReturn("ORDER-RETURN-ID")
func (c *Client) DeleteOrderReturn(id string) error {
	return c.mainRequest("DELETE", "order_returns/"+id, nil, nil)
}

//ListOrderReturns will return a list of all your order returns
//client := epFulfillment.New("YOUR-API-KEY")
// orderReturns, err := client.ListOrderReturns(&epFulfillment.ListOptions{
// 	PerPage: 3,
// 	Page:    0,
// })
func (c *Client) ListOrderReturns(in *ListOptions) (or *OrderReturnList, err error) {
	err = c.mainRequest("GET", "order_returns", in, &or)
	return
}

//GetOrderReturn will retrieve a single OR by id provided
//client := epFulfillment.New("YOUR-API-KEY")
// orderReturn, err := client.GetOrderReturn("ORDER-RETURN-ID")
func (c *Client) GetOrderReturn(id string) (or *OrderReturn, err error) {
	err = c.mainRequest("GET", "order_returns/"+id, nil, &or)
	return
}
