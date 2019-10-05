package main

import (
	"context"
	"fmt"
	"github.com/y-tajiri/go-gmo-pg/client"
	"github.com/y-tajiri/go-gmo-pg/config"
	"strings"
)

func main() {
	cnf := config.Config{
		EndPoint: "https://pt01.mul-pay.jp",
		ShopID: "tshop00036680",
		SiteID: "tsite00029389",
		ShopPass: "nbe6kqqk",
		SitePass: "m5h81x2u",
	}
	cli,err := client.NewClient(cnf)
	if err != nil {
		panic(err)
	}
	orderID := "x12	"
	ctx := context.Background()
	e, err := cli.EntryTranDocomoAcceptIdPass(ctx, orderID)
	if err != nil {
		panic(err)
	}
	req := &client.ExecTranDocomoAcceptIdPassReq{
		AccessID: e.AccessID,
		AccessPass: e.AccessPass,
		OrderID: orderID,
		RetURL: "https://test.careclub.jp/acceptid",
		PaymentTermSec: 3600,
	}
	z, err := cli.ExecTranDocomoAcceptIdPass(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", z.AccessID)
	fmt.Printf("%s\n", strings.Replace(z.Token, " ", "+", -1))
}
