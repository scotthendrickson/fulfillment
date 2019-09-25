package epFulfillment

//Products is a list of products used when using RetrieveAllProducts
type Products struct {
	Products []Product `json:"products,omitempty"`
}

//Product specifies all uniquely barcoded products in your
//catelog that can be used to describe the contents of a AdvancedShipmentNotice or an Order.
type Product struct {
	ID                     string    `json:"id,omitempty"`
	CreatedAt              int64     `json:"created_at,omitempty"`
	UpdatedAt              int64     `json:"updated_at,omitempty"`
	Mode                   string    `json:"mode,omitempty"`
	Title                  string    `json:"title,omitempty"`
	Barcode                string    `json:"barcode,omitempty"`
	Type                   string    `json:"type,omitempty"`
	OriginCountry          string    `json:"origin_country,omitempty"`
	HsCode                 string    `json:"hs_code,omitempty"`
	RequiresSerialTracking bool      `json:"requires_serial_tracking,omitempty"`
	Length                 Dimension `json:"length,omitempty"`
	Width                  Dimension `json:"width,omitempty"`
	Height                 Dimension `json:"height,omitempty"`
	Weight                 Dimension `json:"weight,omitempty"`
	Price                  Dimension `json:"price,omitempty"`
}

//Create is used to create a new Product object.
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// product, err := epFulfillment.Product{
// 	Title:         "T-Shirt Small",
// 	Barcode:       "12345",
// 	Type:          "merchandise",
// 	OriginCountry: "US",
// 	HsCode:        "6103.22.0050",
// 	Length:        epFulfillment.Dimension{Value: 15.0, Unit: "IN"},
// 	Width:         epFulfillment.Dimension{Value: 7.0, Unit: "IN"},
// 	Height:        epFulfillment.Dimension{Value: 1.0, Unit: "IN"},
// 	Weight:        epFulfillment.Dimension{Value: 10.0, Unit: "OZ"},
// 	Price:         epFulfillment.Dimension{Value: 20.0, Unit: "USD"},
// }.Create()
// if err != nil {
// 	fmt.Fprintln(os.Stderr, "error creating", err)
// 	os.Exit(1)
// 	return
// }
// fmt.Printf("%+v\n", product)
func (product Product) Create() (p Product, err error) {
	err = mainRequest("POST", "products/", product, &p)
	return
}

//RetrieveProduct this is used for retrieving a single Product from the API
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// product, err := epFulfillment.RetrieveProduct("PRODUCT-ID")
func RetrieveProduct(id string) (p Product, err error) {
	err = mainRequest("GET", "products/"+id, nil, &p)
	return
}

//RetrieveAllProducts will retrieve a list of all products on the account
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// products, err := epFulfillment.RetrieveAllProducts()
func RetrieveAllProducts() (products Products, err error) {
	err = mainRequest("GET", "products/", nil, &products)
	return
}

//Update allows for updating a single product
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// product, err := epFulfillment.RetrieveProduct("PRODUCT-ID")
// product.Title = "NEW-TITLE"
// product, err = product.Update()
func (product Product) Update() (p Product, err error) {
	err = mainRequest("PATCH", "products/"+product.ID, product, &p)
	return
}

//DeleteProduct for deleting a product by id
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// epFulfillment.DeleteProduct("PRODUCT-ID")
func DeleteProduct(id string) error {
	return mainRequest("DELETE", "products/"+id, nil, nil)
}

//Delete for deleting a product using an interface on the object
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// product, err := epFulfillment.RetrieveProduct("PRODUCT-ID")
// product.Delete()
func (product Product) Delete() error {
	return mainRequest("DELETE", "products/"+product.ID, nil, nil)
}
