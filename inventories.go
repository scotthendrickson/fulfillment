package fulfillment

import (
	"net/http"
)

//Inventories consists of the warehouse where the product is stored and relevant information regarding that product
type Inventories struct {
	Warehouse   *Warehouse `json:"warehouse,omitempty"`
	WarehouseID string     `json:"warehouse_id,omitempty"`
	Product     *Product   `json:"product,omitempty"`
	ProductID   string     `json:"product_id,omitempty"`
	Quantity    int64      `json:"quantity,omitempty"`
}

//InventoriesList is a list of inventories
type InventoriesList struct {
	Inventories []*Inventories `json:"inventories,omitempty"`
}

//InventoriesOptions allows for setting product_id's, includes, and list options
type InventoriesOptions struct {
	Limit      int64    `json:"limit,omitempty"`
	Offset     int64    `json:"offset,omitempty"`
	Page       int      `json:"page,omitempty"`
	PerPage    int64    `json:"per_page,omitempty"`
	ProductIDs []string `json:"product_ids,omitempty"`
	Includes   []string `json:"includes,omitempty"`
}

//ListInventories can be used to retrieve all inventories on your products or only specific products
//
//If a list of product id's is not supplied it will return all products and their quantity of units in
//each warehouse. Any product id's supplied in the list will limit the returned information to just
//those products. When no product id's are supplied you can make use of the options per_page, limit,
//page, and offset to take advantage of pagination on the results. If any product id's are included
//though the API will ignore the options per_page, limit, page, and offset.
//
//The includes list can consist of only two string options, either "product" or "warehouse". If you put "product" in the list it
//will return the whole product object and quantity of units in each warehouse. For "warehouse" if you include it in the list it
//will return the full warehouse object but if left out it will only return the warehouse id.
//
//Example calls:
//client := fulfillment.New("YOUR-API-KEY")
//
//This will return a information on a single product, the warehouse id's, product id, and quantity of units in each warehouse
// inventories, err := client.ListInventories(&fulfillment.InventoriesOptions{
// ProductIDs: []string{"PRODUCT-ID"},
// })
//
//This will return information on a single product, the warehouse id's, the product object, and quantity of units in each warehouse
// inventories, err := client.ListInventories(&fulfillment.InventoriesOptions{
// ProductIDs: []string{"PRODUCT-ID"},
// Includes: []string{"product"},
// })
//
//This will return information on a single product, the each warehouse object, product id, and quantity of units in each warehouse
// inventories, err := client.ListInventories(&fulfillment.InventoriesOptions{
// ProductIDs: []string{"PRODUCT-ID"},
// Includes: []string{"warehouse"},
// })
//
//This will return information on a single product, the each warehouse object, the product object, and quantity of units in each warehouse
// inventories, err := client.ListInventories(&fulfillment.InventoriesOptions{
// ProductIDs: []string{"PRODUCT-ID"},
// Includes: []string{"warehouse", "product"},
// })
//
//This will return a paginated list of products limited to the first page of 5 products, the product object and the warehouse object
// inventories, err := client.ListInventories(&fulfillment.InventoriesOptions{
// 	Includes: []string{"warehouse", "product"},
// 	PerPage:  0,
// 	Page: 5,
// })
func (c *Client) ListInventories(opt *InventoriesOptions) (il *InventoriesList, err error) {
	err = c.do(nil, http.MethodGet, "inventories", opt, &il)
	return
}
