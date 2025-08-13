// Digota <http://digota.com> - eCommerce microservice
// Copyright (c) 2018 Yaron Sumel <yaron@digota.com>
//
// MIT License
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"github.com/digota/digota/sdk"
	"github.com/digota/digota/sku/skupb"
	"github.com/digota/digota/payment/paymentpb"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"time"
	"os"
)

func main() {

	c, err := sdk.NewClient("127.0.0.1:8080", &sdk.ClientOpt{
		InsecureSkipVerify: false,
		ServerName:         "server.com",
		CaCrt:              "out/ca.crt",
		Crt:                "out/client.com.crt",
		Key:                "out/client.com.key",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	rand.Seed(time.Now().Unix())

	if len(os.Args) < 3 {
		log.Fatalf("missing required arguments: sku ID and active flag\nUsage: %s <uuid> {true, false}", os.Args[0])
	}

	uuid := os.Args[1]
	activeArg := os.Args[2]

	var active bool
	switch activeArg {
	case "1", "true", "True", "TRUE":
		active = true
	case "0", "false", "False", "FALSE":
		active = false
	default:
		log.Fatalf("invalid value for active: %s (use 1/0 or true/false or TRUE/FALSE)", activeArg)
	}


	sku, err := skupb.NewSkuServiceClient(c).Update(context.Background(), &skupb.UpdateRequest{
		Id:       uuid,
		Active:   active,
		Price:    uint64(rand.Int31n(10001)),
		Currency: paymentpb.Currency_EUR,
		//name:      fake.Brand(),
		//parent: "cb379ae1-8729-4b32-ba7a-3119dc2bd212",
		//metadata: map[string]string{"key": "val"},
		//image: "http://sadf.124.com",
		//packageDimensions: &skupb.PackageDimensions{weight: 1.1, length: 1.2, height: 1.3, width: 1.4},
		//inventory: &skupb.Inventory{quantity: 1111, type: skupb.Inventory_Finite},
		//attributes: map[string]string{"color": "red"},
	})
	if err != nil {
		log.Fatalf("failed to update sku %s: %v", uuid, err)
	}

	log.Printf("updated sku:\n"+
		"  id: %s\n"+
		"  name: %s\n"+
		"  active: %t\n"+
		"  price: %d %s\n"+
		"  parent product id: %s\n"+
		"  attributes: %v\n"+
		"  metadata: %v\n"+
		"  image: %s\n"+
		"  package dimensions: {weight: %.2f, length: %.2f, height: %.2f, width: %.2f}\n"+
		"  inventory: {quantity: %d, type: %s}\n"+
		"  created: %d\n"+
		"  updated: %d\n",
		sku.Id,
		sku.Name,
		sku.Active,
		sku.Price,
		sku.Currency.String(),
		sku.Parent,
		sku.Attributes,
		sku.Metadata,
		sku.Image,
		sku.PackageDimensions.Weight,
		sku.PackageDimensions.Length,
		sku.PackageDimensions.Height,
		sku.PackageDimensions.Width,
		sku.Inventory.Quantity,
		sku.Inventory.Type.String(),
		sku.Created,
		sku.Updated,
	)
}
