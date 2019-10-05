package main

import (
	"context"
	"fmt"
	"github.com/y-tajiri/go-gmo-pg/client"
	"github.com/y-tajiri/go-gmo-pg/config"
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
	orderID := "x13"
	ctx := context.Background()
	e, err := cli.EntryTranDocomoIdPass(ctx, orderID,1000,100)
	if err != nil {
		panic(err)
	}
	req := &client.ExecTranDocomoIdPassReq{
		AccessID: e.AccessID,
		AccessPass: e.AccessPass,
		OrderID: orderID,
		RetURL: "https://test.careclub.jp/acceptid",
		PaymentTermSec: 3600,
		DocomoAcceptCode: "T19278091749",
	}
	fmt.Printf("xxx")
	z, err := cli.ExecTranDocomoIdPass(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", z)
}
