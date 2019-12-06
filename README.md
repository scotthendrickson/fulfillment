# epFulfillment- Go Client Library for EasyPost's Fulfillment API

EasyPost offers an API to their fulfillment customers that allows them to manage all their products, inventories, and orders. Their documentation can be found here https://gist.github.com/att14/ff68a0f2684c711444864dcb1ebf6030 and their website is https://easypost.com.

## Requirements

This package should work with any recent version of Go

## Installation

```go get -u https://github.com/scotthendrickson/epFulfillment```

## Example

Examples can be found on each file but here is an example of how to create a product using this library.

### Product Creation

```go
package main

import (
    "fmt"
    "os"

    "github.com/scotthendrickson/epFulfillment"
)

func main() {

    client := epFulfillment.New("YOUR-API-KEY")

    product, err := client.CreateProduct(&epFulfillment.Product{
        Title:         "Tester McTester Mouse",
        Barcode:       "8161616161616",
        Type:          "merchandise",
        OriginCountry: "US",
        HsCode:        "6103.22.0050",
        Length:        &epFulfillment.Dimension{Value: 15.0, Unit: "IN"},
        Width:         &epFulfillment.Dimension{Value: 7.0, Unit: "IN"},
        Height:        &epFulfillment.Dimension{Value: 1.0, Unit: "IN"},
        Weight:        &epFulfillment.Dimension{Value: 10.0, Unit: "OZ"},
        Price:         &epFulfillment.Dimension{Value: 20.0, Unit: "USD"},
    })
if err != nil {
    fmt.Fprintln(os.Stderr, "error creating", err)
    os.Exit(1)
    return
}
fmt.Printf("%+v\n", product)

    if err != nil {
        fmt.Fprintln(os.Stderr, "error creating", err)
        os.Exit(1)
        return
    }

    fmt.Printf("%+v\n", product)

}
```
