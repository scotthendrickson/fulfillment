package epFulfillment

//Orders a list of orders used for when a list is retrieved from the API
type Orders struct {
	Orders []Order `json:"orders,omitempty"`
}

//Order describes the destination and the set of purchased goods to be mailed to that destination.
type Order struct {
	ID              string          `json:"id,omitempty"`
	CreatedAt       int             `json:"created_at,omitempty"`
	UpdatedAt       int             `json:"updated_at,omitempty"`
	Mode            string          `json:"mode,omitempty"`
	LatestDelivery  string          `json:"latest_delivery,omitempty"`
	LineItems       []OrderLineItem `json:"line_items,omitempty"`
	Inserts         []Insert        `json:"inserts,omitempty"`
	Destination     Address         `json:"destination,omitempty"`
	Description     string          `json:"description,omitempty"`
	ShipmentOptions Options         `json:"shipment_options,omitempty"`
	Status          string          `json:"status,omitempty"`
	Picks           []Pick          `json:"picks,omitempty"`
	Fees            []Fee           `json:"fees,omitempty"`
	Service         string          `json:"service,omitempty"`
}

//OrderLineItem for individual products on orders
type OrderLineItem struct {
	ID        string  `json:"id,omitempty"`
	CreatedAt int     `json:"created_at,omitempty"`
	UpdatedAt int     `json:"updated_at,omitempty"`
	Mode      string  `json:"mode,omitempty"`
	Product   Product `json:"product,omitempty"`
	Units     int     `json:"units,omitempty"`
}

//Currently ignoring Metadata on Inserts but I'll have to work that out later

//Insert is for specific documents that could be put inside of the orders shipment
type Insert struct {
	ID        string `json:"id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	Mode      string `json:"mode,omitempty"`
	URL       string `json:"url,omitempty"`
	ACL       string `json:"acl,omitempty"`
}

//Pick after an order has been received by the warehouse a pick will be initialized which will describe how many
// items are sent in each package as well as the tracker for that package
type Pick struct {
	ID          string      `json:"id,omitempty"`
	CreatedAt   int         `json:"created_at,omitempty"`
	UpdatedAt   int         `json:"updated_at,omitempty"`
	Mode        string      `json:"mode,omitempty"`
	Tracker     Tracker     `json:"tracker,omitempty"`
	Inventories []Inventory `json:"inventories,omitempty"`
}

//Inventory list of products and quantities
type Inventory struct {
	ID        string  `json:"id,omitempty"`
	CreatedAt int     `json:"created_at,omitempty"`
	UpdatedAt int     `json:"updated_at,omitempty"`
	Mode      string  `json:"mode,omitempty"`
	Product   Product `json:"product,omitempty"`
}

//Options being set on the order
type Options struct {
	DeliveryConfirmation string `json:"delivery_confirmation,omitempty"`
	Insurance            bool   `json:"insurance,omitempty"`
}

//Fee fee type and information
type Fee struct {
	ID         string `json:"id,omitempty"`
	CreatedAt  int    `json:"created_at,omitempty"`
	UpdatedAt  int    `json:"updated_at,omitempty"`
	DisabledAt int    `json:"disabled_at,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Amount     string `json:"amount,omitempty"`
	Charged    bool   `json:"charged,omitempty"`
	Refunded   bool   `json:"refunded,omitempty"`
	Type       string `json:"type"`
}

//Address used to set addresses on orders and such
type Address struct {
	ID           string `json:"id,omitempty"`
	CreatedAt    int    `json:"created_at,omitempty"`
	UpdatedAt    int    `json:"updated_at,omitempty"`
	Mode         string `json:"mode,omitempty"`
	Name         string `json:"name,omitempty"`
	Company      string `json:"company,omitempty"`
	Street1      string `json:"street1,omitempty"`
	Street2      string `json:"street2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Country      string `json:"country,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	StateTaxID   string `json:"state_tax_id,omitempty"`
	FederalTaxID string `json:"federal_tax_id,omitempty"`
	Residential  bool   `json:"residential,omitempty"`
}

//Create is for creating a single product
//epFulfillment.SetAPIKey("YOUR-API-KEY")
// product, err := epFulfillment.RetrieveProduct("PRODUCT-ID")
// orderLineItem := epFulfillment.OrderLineItem{
// 	Product: product,
// 	Units:   1,
// }
// order, err := epFulfillment.Order{
// 	Service:     "Standard",
// 	LineItems:   []epFulfillment.OrderLineItem{orderLineItem},
// 	Description: "PO#12345",
// 	Destination: epFulfillment.Address{
// 		Name:    "Scott Hendrickson",
// 		Street1: "417 MONTGOMERY ST FL 5",
// 		City:    "San Francisco",
// 		State:   "CA",
// 		Zip:     "94104",
// 		Country: "US",
// 	},
// 	ShipmentOptions: epFulfillment.Options{
// 		DeliveryConfirmation: "SIGNATURE",
// 		Insurance:            false,
// 	},
// }.Create()
// if err != nil {
// 	fmt.Fprintln(os.Stderr, "error creating", err)
// 	os.Exit(1)
// 	return
// }
// fmt.Printf("%+v\n", order)
func (order Order) Create() (o Order, err error) {
	err = mainRequest("POST", "orders/", order, &o)
	return
}

//RetrieveAllOrders will retrieve a list of all orders on the account
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//orders, err := epFulfillment.RetrieveAllOrders()
func RetrieveAllOrders() (orders Orders, err error) {
	err = mainRequest("GET", "orders/", nil, &orders)
	return
}

//RetrieveOrder will retrieve an order by it's ID
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//order, err := epFulfillment.RetrieveOrder("ORDER-ID")
func RetrieveOrder(id string) (o Order, err error) {
	err = mainRequest("GET", "orders/"+id, nil, &o)
	return
}

//Update can be used to patch an order once retrieved
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//order, err := epFulfillment.RetrieveOrder("ORDER-ID")
//	order.Destination = epFulfillment.Address{
// 	Name:    "Scott Hendrickson",
// 	Street1: "417 MONTGOMERY ST FL 5",
// 	City:    "San Francisco",
// 	State:   "CA",
// 	Zip:     "94104",
// 	Country: "US",
// }
// order.Update()
func (order Order) Update() (o Order, err error) {
	err = mainRequest("PATCH", "orders/"+order.ID, order, &o)
	return
}

//Delete can be used to delete an order once retrieved
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//order, err := epFulfillment.RetrieveOrder("ORDER-ID")
//order.Delete()
func (order Order) Delete() error {
	return mainRequest("DELETE", "orders/"+order.ID, order, nil)
}

//DeleteOrder can be used to delete an order using just the ID without retrieving the order first
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//epFulfillment.DeleteOrder("ORDER-ID")
func DeleteOrder(id string) error {
	return mainRequest("DELETE", "orders/"+id, nil, nil)
}
