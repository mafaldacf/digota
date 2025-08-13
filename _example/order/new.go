// Digota <http://digota.com> - eCommerce microservice
// Copyright (c) 2018 Yaron Sumel <yaron@digota.com>
//
// MIT License
// Permission is hereby granted, free of charge, to any person obtaining arg copy
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
	"github.com/digota/digota/order/orderpb"
	"github.com/digota/digota/payment/paymentpb"
	"github.com/digota/digota/sdk"
	"golang.org/x/net/context"
	"log"
	"os"
	"strings"
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
		log.Fatalf("missing required arguments: at least one sku ID\nusage: %s <sku-uuid> [<sku-uuid> ...]", os.Args[0])
	}

	var uuid []string
	for _, arg := range os.Args[1:] {
		uuid = append(uuid, arg)
	}

	quantity := make(map[string]int)
	uuids := make([]string, 0, len(os.Args)-1)
	for _, arg := range os.Args[1:] {
		uuid := strings.TrimSpace(arg)
		if uuid == "" {
			continue
		}
		if _, seen := quantity[uuid]; !seen {
			uuids = append(uuids, uuid)
		}
		quantity[uuid]++
	}

	items := make([]*orderpb.OrderItem, 0, len(quantity))
	for _, uuid := range uuids {
		items = append(items, &orderpb.OrderItem{
			Parent:   uuid,
			Quantity: int64(quantity[uuid]),
			Type:     orderpb.OrderItem_sku,
		})
	}

	order, err := orderpb.NewOrderServiceClient(c).New(context.Background(), &orderpb.NewRequest{
		Currency: paymentpb.Currency_EUR,
		Items:    items,
		Email:    "yaron@digota.com",
		Shipping: &orderpb.Shipping{
			Name:  "Yaron Sumel",
			Phone: "+972 000 000 000",
			Address: &orderpb.Shipping_Address{
				Line1:      "Loren ipsum",
				City:       "San Jose",
				Country:    "USA",
				Line2:      "",
				PostalCode: "12345",
				State:      "CA",
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to create order: %v", err)
	}

	log.Printf("created order:\n  id: %s\n  currency: %s\n  items:", order.Id, paymentpb.Currency_EUR.String())
	for i, item := range items {
		log.Printf("    [%d] parent: %s | qty: %d | type: %s", i+1, item.Parent, item.Quantity, orderpb.OrderItem_sku.String())
	}
}
