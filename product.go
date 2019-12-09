package fulfillment

import "net/http"

//Products is a list of products used when using ListProducts
type Products struct {
	Products []*Product `json:"products,omitempty"`
}

//Product specifies all uniquely barcoded products in your
//catelog that can be used to describe the contents of a AdvancedShipmentNotice or an Order.
type Product struct {
	ID                     string     `json:"id,omitempty"`
	CreatedAt              int64      `json:"created_at,omitempty"`
	UpdatedAt              int64      `json:"updated_at,omitempty"`
	Mode                   string     `json:"mode,omitempty"`
	Title                  string     `json:"title,omitempty"`
	Barcode                string     `json:"barcode,omitempty"`
	Type                   string     `json:"type,omitempty"`
	OriginCountry          string     `json:"origin_country,omitempty"`
	HsCode                 string     `json:"hs_code,omitempty"`
	RequiresSerialTracking bool       `json:"requires_serial_tracking,omitempty"`
	Length                 *Dimension `json:"length,omitempty"`
	Width                  *Dimension `json:"width,omitempty"`
	Height                 *Dimension `json:"height,omitempty"`
	Weight                 *Dimension `json:"weight,omitempty"`
	Price                  *Dimension `json:"price,omitempty"`
}

//Dimension is used within Product struct to designate a value and unit type for the dimension being specified
type Dimension struct {
	Value float64 `json:"value,omitempty"`
	Unit  string  `json:"unit,omitempty"`
}

//CreateProduct is used to create a new Product object.
// client := epFulfillment.New("YOUR-API-KEY")
// 	product, err := client.CreateProduct(&epFulfillment.Product{
// 		Title:         "Tester McTester Mouse",
// 		Barcode:       "8161616161616",
// 		Type:          "merchandise",
// 		OriginCountry: "US",
// 		HsCode:        "6103.22.0050",
// 		Length:        &epFulfillment.Dimension{Value: 15.0, Unit: "IN"},
// 		Width:         &epFulfillment.Dimension{Value: 7.0, Unit: "IN"},
// 		Height:        &epFulfillment.Dimension{Value: 1.0, Unit: "IN"},
// 		Weight:        &epFulfillment.Dimension{Value: 10.0, Unit: "OZ"},
// 		Price:         &epFulfillment.Dimension{Value: 20.0, Unit: "USD"},
// 	})
// if err != nil {
// 	fmt.Fprintln(os.Stderr, "error creating", err)
// 	os.Exit(1)
// 	return
// }
// fmt.Printf("%+v\n", product)
func (c *Client) CreateProduct(product *Product) (p *Product, err error) {
	err = c.post(nil, "products/", product, &p)
	return
}

//GetProduct this is used for retrieving a single Product from the API
// client := epFulfillment.New("YOUR-API-KEY")
// product, err := client.GetProduct("PRODUCT-ID")
// if err != nil {
// 		fmt.Fprintln(os.Stderr, "error creating", err)
// 		os.Exit(1)
// 		return
// 	}
// fmt.Printf("%+v\n", product)
func (c *Client) GetProduct(id string) (p *Product, err error) {
	err = c.get(nil, "products/"+id, &p)
	return
}

//ListProducts will retrieve a list of all products on the account
// client := epFulfillment.New("YOUR-API-KEY")
// 	products, err := client.ListProducts(&epFulfillment.ListOptions{
// 		PerPage: 3,
// 		Page:    2,
// 	})
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "error creating", err)
// 		os.Exit(1)
// 		return
// 	}
// 	for i := range products.Products {
// 		fmt.Printf("%+v\n", products.Products[i].ID)
// 	}
func (c *Client) ListProducts(opt *ListOptions) (products *Products, err error) {
	err = c.do(nil, http.MethodGet, "products/", opt, &products)
	return
}

//UpdateProduct allows for updating a single product
// client := epFulfillment.New("YOUR-API-KEY")
// product, err := client.GetProduct("PRODUCT-ID")
// product.Title = "New Tester McTester Mouse"
// product, err = client.UpdateProduct(product)
// fmt.Printf("%+v\n", product)
func (c *Client) UpdateProduct(product *Product) (p *Product, err error) {
	err = c.patch(nil, "products/"+product.ID, &product, &p)
	return
}

//DeleteProduct for deleting a product by id
// client := epFulfillment.New("YOUR-API-KEY")
// client.DeleteProduct("PRODUCT-ID")
func (c *Client) DeleteProduct(id string) error {
	return c.del(nil, "products/"+id)
}
