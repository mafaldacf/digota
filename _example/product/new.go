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
	"github.com/icrowley/fake"
	"golang.org/x/net/context"
	"log"
)

func main() {

	c, err := sdk.NewClient("127.0.0.1:8080", &sdk.ClientOpt{
		InsecureSkipVerify: true,
		ServerName:         "server.com",
		CaCrt:              "out/ca.crt",
		Crt:                "out/client.com.crt",
		Key:                "out/client.com.key",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	resp, err := productpb.NewProductServiceClient(c).New(context.Background(), &productpb.NewRequest{
		Name:        fake.Product(),
		Active:      true,
		Attributes:  []string{"size"},
		Description: fake.Model(),
		Images: []string{
			"http://digota.com/image1.jpg",
			"http://digota.com/image2.jpg",
		},
		Metadata: map[string]string{
			"mymetadata": "12345",
		},
		Shippable: true,
		Url:       "http://digota.com",
	})
	if err != nil {
		log.Fatalf("failed to create product: %v", err)
	}

	log.Printf("Created Product:\n  ID: %s\n  Name: %s\n  Active: %t\n  Description: %s\n  URL: %s\n  Attributes: %v\n  Images: %v\n  Metadata: %v\n  Shippable: %t\n  Created: %d\n",
		resp.Id,
		resp.Name,
		resp.Active,
		resp.Description,
		resp.Url,
		resp.Attributes,
		resp.Images,
		resp.Metadata,
		resp.Shippable,
		resp.Created,
	)
}
