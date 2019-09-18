package epFulfillment

//Products for list of products
type Products struct {
	Products []Product `json:"products,omitempty"`
}

//Product is for EP fulfillment products
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

//Create this function allows for the creation of a single product
func (product Product) Create() (p Product, err error) {
	err = mainRequest("POST", "products/", product, &p)
	return
}

//RetrieveProduct this is used for retrieving a single Product from the API
func RetrieveProduct(id string) (p Product, err error) {
	err = mainRequest("GET", "products/"+id, nil, &p)
	return
}

//RetrieveAllProducts will retrieve a list of all products on the account
func RetrieveAllProducts() (products Products, err error) {
	err = mainRequest("GET", "products/", nil, &products)
	return
}

//Update allows for updating a single product
func (product Product) Update() (p Product, err error) {
	err = mainRequest("PATCH", "products/"+product.ID, product, &p)
	return
}

//DeleteProduct for deleting a product by id
func DeleteProduct(id string) error {
	return mainRequest("DELETE", "products/"+id, nil, nil)
}

//Delete for deleting a product using an interface on the object
func (product Product) Delete() error {
	return mainRequest("DELETE", "products/"+product.ID, nil, nil)
}
