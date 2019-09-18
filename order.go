package epFulfillment

//Orders a list of orders used for when a list is retrieved from the API
type Orders struct {
	Orders []Order `json:"orders,omitempty"`
}

//Order for orders to teh EP API
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

//OrderLineItem for individual line items on orders
type OrderLineItem struct {
	ID        string  `json:"id,omitempty"`
	CreatedAt int     `json:"created_at,omitempty"`
	UpdatedAt int     `json:"updated_at,omitempty"`
	Mode      string  `json:"mode,omitempty"`
	Product   Product `json:"product,omitempty"`
	Units     int     `json:"units,omitempty"`
}

//Currently ignoring Metadata on Inserts but I'll have to work that out later

//Insert is for notes that could be put inside of the orders shipment
type Insert struct {
	ID        string `json:"id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	Mode      string `json:"mode,omitempty"`
	URL       string `json:"url,omitempty"`
	ACL       string `json:"acl,omitempty"`
}

//Pick species which inventory will be part of the pick
type Pick struct {
	ID        string `json:"id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	Mode      string `json:"mode,omitempty"`
	// Tracker Tracker `json:"tracker,omitempty"`
	// Inventories Inventory[] `json:"inventories,omitempty"`
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
func (order Order) Create() (o Order, err error) {
	err = mainRequest("POST", "orders/", order, &o)
	return
}

//RetrieveAllOrders will retrieve a list of all orders on the account
func RetrieveAllOrders() (orders Orders, err error) {
	err = mainRequest("GET", "orders/", nil, &orders)
	return
}

//RetrieveOrder will retrieve an order by it's ID
func RetrieveOrder(id string) (o Order, err error) {
	err = mainRequest("GET", "orders/"+id, nil, &o)
	return
	// return o, Get(fmt.Sprintf("orders/%s", id), &o)
}

//Update can be used to patch an order once retrieved
func (order Order) Update() (o Order, err error) {
	err = mainRequest("PATCH", "orders/"+order.ID, order, &o)
	return
}

//Delete can be used to delete an order once retrieved
func (order Order) Delete() error {
	return mainRequest("DELETE", "orders/"+order.ID, order, nil)
}

//DeleteOrder can be used to delete an order using just the ID without retrieving the order first
func DeleteOrder(id string) error {
	return mainRequest("DELETE", "orders/"+id, nil, nil)
}
