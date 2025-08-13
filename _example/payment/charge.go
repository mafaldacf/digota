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
	"github.com/digota/digota/payment/paymentpb"
	"github.com/digota/digota/sdk"
	"golang.org/x/net/context"
	"log"
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

	// Charge amount
	log.Println(paymentpb.NewPaymentServiceClient(c).NewCharge(context.Background(), &paymentpb.ChargeRequest{
		Total:             10 * 1000,
		Currency:          paymentpb.Currency_USD,
		Email:             "yaron@digota.com",
		Statement:         "Charge statement",
		PaymentProviderId: paymentpb.PaymentProviderId_Stripe,
		Card: &paymentpb.Card{
			Type:        paymentpb.CardType_Visa,
			CVC:         "123",
			ExpireMonth: "12",
			ExpireYear:  "2022",
			LastName:    "Sumel",
			FirstName:   "Yaron",
			Number:      "4111111111111111",
		},
	}))

}
