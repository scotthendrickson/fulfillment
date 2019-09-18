package epFulfillment

//Inventories consists of the warehouse where the product is stored and relevant information regarding that product
type Inventories struct {
	Warehouse   Warehouse `json:"warehouse,omitempty"`
	WarehouseID string    `json:"warehouse_id,omitempty"`
	Product     Product   `json:"product,omitempty"`
	ProductID   string    `json:"product_id,omitempty"`
	Quantity    int64     `json:"quantity,omitempty"`
}

//InventoriesList is a list of inventories
type InventoriesList struct {
	Inventories []Inventories `json:"inventories,omitempty"`
}

//RetrieveInventories can be used to retrieve all inventories on your products or only specific products
//If you want to retrieve information on only specific products just supply a list of all the product id's you want information on
//The includes option should be limited to a list of one or two string and the only possible values are "product" or "warehouse".
//If product is supplied the whole product object will be returned, and if not included just the products ID will be returned
//If warehouse is supplied the whole warehouse object will be returned, and if not just the warehouse ID will be returned
//If no includes or products are passed into the function than a list of all product id's, the warehouse id's where they are stored
//and their quantity will be returned
func RetrieveInventories(productIds []string, includes []string) (il InventoriesList, err error) {
	url := ""
	if includes != nil || productIds != nil {
		url = "?"
	}
	if productIds != nil {
		for i := range productIds {
			url = url + "&product_ids[]=" + productIds[i]
		}
	}
	if includes != nil {
		for i := range includes {
			if includes[i] == "product" || includes[i] == "warehouse" {
				url = url + "&includes[]=" + includes[i]
			}
		}
	}
	err = mainRequest("GET", "inventories"+url, nil, &il)
	return
}
