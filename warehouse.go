package epFulfillment

//Warehouse is how to manage where your inventory is
type Warehouse struct {
	ID         string  `json:"id,omitempty"`
	CreatedAt  int64   `json:"created_at,omitempty"`
	UpdatedAt  int64   `json:"updated_at,omitempty"`
	DisabledAt int64   `json:"disabled_at,omitempty"`
	Mode       string  `json:"mode,omitempty"`
	Name       string  `json:"name,omitempty"`
	ShortName  string  `json:"short_name,omitempty"`
	Address    Address `json:"address,omitempty"`
}

//Warehouses is a list of Warehouse objects
type Warehouses struct {
	Warehouses []Warehouse `json:"warehouses,omitempty"`
}

//RetrieveAllWarehouses will allow you to retrieve a list of all warehouses
func RetrieveAllWarehouses() (warehouses Warehouses, err error) {
	err = mainRequest("GET", "warehouses", nil, &warehouses)
	return
}

//Doesn't seem to be a way to retrieve a single warehouse?
//RetrieveWarehouse takes a string ID and returns an EasyPost warehouse object
// func RetrieveWarehouse(id string) (warehouse Warehouse, err error) {
// 	err = mainRequest("GET", "warehouses/"+id, nil, &warehouse)
// 	return
// }
