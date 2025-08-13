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
	"github.com/digota/digota/order/orderpb"
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
		log.Fatalf("missing required argument: order ID\nUsage: %s <uuid>", os.Args[0])
	}
	uuid := os.Args[1]

	// Return order
	order, err := orderpb.NewOrderServiceClient(c).Return(context.Background(), &orderpb.ReturnRequest{
		Id: uuid,
	})
	if err != nil {
		log.Fatalf("failed to return order with id %s: %v", uuid, err)
	}

	log.Printf("returned order:\n"+
		"  id: %s\n"+
		"  amount: %d\n"+
		"  currency: %s\n"+
		"  status: %s\n"+
		"  email: %s\n"+
		"  charge id: %s\n"+
		"  metadata: %v\n"+
		"  created: %d\n"+
		"  updated: %d\n"+
		"  items:\n",
		order.Id,
		order.Amount,
		order.Currency.String(),
		order.Status.String(),
		order.Email,
		order.ChargeId,
		order.Metadata,
		order.Created,
		order.Updated,
	)

	for i, order := range order.Items {
		log.Printf("    [%d] parent: %s | qty: %d | type: %s | amount: %d %s",
			i+1,
			order.Parent,
			order.Quantity,
			order.Type.String(),
			order.Amount,
			order.Currency.String(),
		)
	}

}
