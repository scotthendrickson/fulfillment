# fulfillment- Go Client Library for EasyPost's Fulfillment API

EasyPost offers an API to their fulfillment customers that allows them to manage all their products, inventories, and orders. Their documentation can be found here https://gist.github.com/att14/ff68a0f2684c711444864dcb1ebf6030 and their website is https://easypost.com.

## Requirements

This package should work with any recent version of Go

## Installation

```go get -u https://github.com/scotthendrickson/fulfillment```

## Example

Examples can be found on each file but here is an example of how to create a product using this library.

### Product Creation

```go
package main

import (
    "fmt"
    "os"

    "github.com/scotthendrickson/fulfillment"
)

func main() {

    client := fulfillment.New("YOUR-API-KEY")

    product, err := client.CreateProduct(&fulfillment.Product{
        Title:         "Tester McTester Mouse",
        Barcode:       "8161616161616",
        Type:          "merchandise",
        OriginCountry: "US",
        HsCode:        "6103.22.0050",
        Length:        &fulfillment.Dimension{Value: 15.0, Unit: "IN"},
        Width:         &fulfillment.Dimension{Value: 7.0, Unit: "IN"},
        Height:        &fulfillment.Dimension{Value: 1.0, Unit: "IN"},
        Weight:        &fulfillment.Dimension{Value: 10.0, Unit: "OZ"},
        Price:         &fulfillment.Dimension{Value: 20.0, Unit: "USD"},
    })

    if err != nil {
        fmt.Fprintln(os.Stderr, "error creating", err)
        os.Exit(1)
        return
    }

    fmt.Printf("%+v\n", product)

}
```
