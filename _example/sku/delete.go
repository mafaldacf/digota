package main

import (
	"github.com/digota/digota/sdk"
	"github.com/digota/digota/sku/skupb"
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

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		log.Fatalf("missing required argument: sku ID\nUsage: %s <uuid>", os.Args[0])
	}
	uuid := os.Args[1]

	skupb.NewSkuServiceClient(c).Delete(context.Background(), &skupb.DeleteRequest{
		Id: uuid,
	})

	log.Printf("deleted sku (id=%s)\n", uuid)
}
