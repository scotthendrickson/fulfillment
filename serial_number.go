package fulfillment

//SerialNumber is for serial numbers on Products. As EP defines "A unique identifier for an instance of a Product, aka an Inventory."
type SerialNumber struct {
	ID        string   `json:"id,omitempty"`
	CreatedAt int64    `json:"created_at,omitempty"`
	UpdatedAt int64    `json:"updated_at,omitempty"`
	Mode      string   `json:"mode,omitempty"`
	Value     string   `json:"value,omitempty"`
	Product   *Product `json:"product,omitempty"`
}

//CreateSerialNumber is for creating a serial number on a product.
// client := fulfillment.New("YOUR-API-KEY")
// product, err := client.GetProduct("PRODUCT-ID")
// serialNumber, err := client.CreateSerialNumber(&fulfillment.SerialNumber{
// 		Value:   "SERIAL-NUMBER-VALUE",
// 		Product: product,
// 	})
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "error creating", err)
// 		os.Exit(1)
// 		return
// 	}
// 	fmt.Println(serialNumber.ID)
func (c *Client) CreateSerialNumber(serialNumber *SerialNumber) (s *SerialNumber, err error) {
	err = c.post(nil, "products/"+serialNumber.Product.ID+"/serial_numbers", serialNumber, &s)
	return
}
