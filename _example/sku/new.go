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

	if len(os.Args) < 2 {
		log.Fatalf("missing required argument: product ID\nUsage: %s <uuid>", os.Args[0])
	}
	uuid := os.Args[1]

	rand.Seed(time.Now().UnixNano())

	resp, err := skupb.NewSkuServiceClient(c).New(context.Background(), &skupb.NewRequest{
		Name:     "Snaptags Portable Audible Amplifier",
		Active:   true,
		Price:    uint64(rand.Int31n(10001)),
		Currency: paymentpb.Currency_EUR,
		Parent:   uuid,
		//Metadata: map[string]string{
		//	"key": "val",
		//},
		Image: "http://sadf.124.com",
		PackageDimensions: &skupb.PackageDimensions{
			Weight: 1.1,
			Length: 1.2,
			Height: 1.3,
			Width:  1.4,
		},
		Inventory: &skupb.Inventory{
			Quantity: 1111,
			Type:     skupb.Inventory_Finite,
		},
		Attributes: map[string]string{
			"color": "red",
		},
	})
	if err != nil {
		log.Fatalf("failed to create SKU: %v", err)
	}

	// human-readable output
	log.Printf("Created SKU:\n"+
		"  ID: %s\n"+
		"  Name: %s\n"+
		"  Active: %t\n"+
		"  Price: %d %s\n"+
		"  Parent Product ID: %s\n"+
		"  Attributes: %v\n"+
		"  Metadata: %v\n"+
		"  Image: %s\n"+
		"  Package Dimensions: {Weight: %.2f, Length: %.2f, Height: %.2f, Width: %.2f}\n"+
		"  Inventory: {Quantity: %d, Type: %s}\n"+
		"  Created: %d\n",
		resp.Id,
		resp.Name,
		resp.Active,
		resp.Price,
		resp.Currency.String(),
		resp.Parent,
		resp.Attributes,
		resp.Metadata,
		resp.Image,
		resp.PackageDimensions.Weight,
		resp.PackageDimensions.Length,
		resp.PackageDimensions.Height,
		resp.PackageDimensions.Width,
		resp.Inventory.Quantity,
		resp.Inventory.Type.String(),
		resp.Created,
	)
}
