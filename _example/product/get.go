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
	"github.com/digota/digota/product/productpb"
	"github.com/digota/digota/sdk"
	"golang.org/x/net/context"
	"log"
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

	resp, err := productpb.NewProductServiceClient(c).Get(context.Background(), &productpb.GetRequest{
		Id: uuid,
	})
	if err != nil {
		log.Fatalf("failed to get product with ID %s: %v", uuid, err)
	}

	log.Printf("Fetched Product:\n"+
		"  ID: %s\n"+
		"  Name: %s\n"+
		"  Active: %t\n"+
		"  Attributes: %v\n"+
		"  Description: %s\n"+
		"  Images: %v\n"+
		"  Metadata: %v\n"+
		"  Shippable: %t\n"+
		"  URL: %s\n"+
		"  SKUs: %d\n"+
		"  Created: %d\n"+
		"  Updated: %d\n",
		resp.Id,
		resp.Name,
		resp.Active,
		resp.Attributes,
		resp.Description,
		resp.Images,
		resp.Metadata,
		resp.Shippable,
		resp.Url,
		len(resp.Skus),
		resp.Created,
		resp.Updated,
	)

}
