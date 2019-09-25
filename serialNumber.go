package epFulfillment

//SerialNumber is for serial numbers on Products. As EP defines "A unique identifier for an instance of a Product, aka an Inventory."
type SerialNumber struct {
	ID        string  `json:"id,omitempty"`
	CreatedAt int64   `json:"created_at,omitempty"`
	UpdatedAt int64   `json:"updated_at,omitempty"`
	Mode      string  `json:"mode,omitempty"`
	Value     string  `json:"value,omitempty"`
	Product   Product `json:"product,omitempty"`
}

//Create is for creating a serial number on a product.
// epFulfillment.SetAPIKey("YOUR-API-KEY")
// product, err := epFulfillment.RetrieveProduct("PRODUCT-ID")
// serialNumber, err := epFulfillment.SerialNumber{
// 	Product: product,
// 	Value:   "abcd123456",
// }.Create()
// if err != nil {
// 	fmt.Fprintln(os.Stderr, "error creating", err)
// 	os.Exit(1)
// 	return
// }
// fmt.Printf("%+v\n", serialNumber)
func (serialNumber SerialNumber) Create() (s SerialNumber, err error) {
	err = mainRequest("POST", "products/"+serialNumber.Product.ID+"/serial_numbers", serialNumber, &s)
	return
}
