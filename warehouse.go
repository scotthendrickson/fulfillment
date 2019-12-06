package epFulfillment

//Warehouse is how to manage where your inventory is
type Warehouse struct {
	ID         string   `json:"id,omitempty"`
	CreatedAt  int64    `json:"created_at,omitempty"`
	UpdatedAt  int64    `json:"updated_at,omitempty"`
	DisabledAt int64    `json:"disabled_at,omitempty"`
	Mode       string   `json:"mode,omitempty"`
	Name       string   `json:"name,omitempty"`
	ShortName  string   `json:"short_name,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

//Warehouses is a list of Warehouse objects
type Warehouses struct {
	Warehouses []*Warehouse `json:"warehouses,omitempty"`
}

//ListWarehouses will allow you to retrieve a list of all warehouses
//epFulfillment.SetAPIKey("YOUR-API-KEY")
//warehouses, err := epFulfillment.RetrieveAllWarehouses()
func (c *Client) ListWarehouses() (warehouses *Warehouses, err error) {
	err = c.mainRequest("GET", "warehouses", nil, &warehouses)
	return
}
