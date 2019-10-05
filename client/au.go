package client

import (
	"context"
	"fmt"
	"github.com/y-tajiri/go-gmo-pg/errors"
	"io/ioutil"
	"net/url"
	"strconv"
)

type (
	EntryTranAuAcceptIdPassResponse struct{
		AccessID   string
		AccessPass string
	}
	EntryTranAuIdPassResponse struct {
		AccessID   string
		AccessPass string
	}
	ExecTranAuAcceptIdPassReq struct {
		AccessID           string
		AccessPass         string
		OrderID            string
		Commodity          string
		ClientField1       string
		ClientField2       string
		ClientField3       string
		RetURL             string
		PaymentTermSec     int
		ServiceName        string
		ServiceTel         string
	}

	ExecTranAuIdPassReq struct {
		AccessID           string
		AccessPass         string
		OrderID            string
		ClientField1       string
		ClientField2       string
		ClientField3       string
		Commodity          string
		ServiceName        string
		ServiceTel         string
		AuAcceptCode       string
	}

	ExecTranAuAcceptIdPassResponse struct {
		AccessID       string
		Token          string
	}
	ExecTranAuIdPassResponse struct {
		OrderID       string
	}
)

func (c *Client) EntryTranAuAcceptIdPass(ctx context.Context, orderID string) (*EntryTranAuIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	resp, err := c.Post(ctx, "/payment/EntryTranAuAccept.idPass", data, false)
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	retVals, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	if retVals.Get("ErrCode") != "" {
		return nil, errors.NewErrorGMOPG(retVals.Get("ErrCode"), retVals.Get("ErrInfo"))
	}
	ret := &EntryTranAuIdPassResponse{}
	ret.AccessID = retVals["AccessID"][0]
	ret.AccessPass = retVals["AccessPass"][0]
	return ret, nil

}

func (c *Client) ExecTranAuAcceptIdPass(ctx context.Context, req *ExecTranAuAcceptIdPassReq) (*ExecTranAuAcceptIdPassResponse, error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTranAuAccept.idPass", data, false)
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	retVals, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	if retVals.Get("ErrCode") != "" {
		return nil, errors.NewErrorGMOPG(retVals.Get("ErrCode"), retVals.Get("ErrInfo"))
	}
	ret := &ExecTranAuAcceptIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.Token = retVals.Get("Token")
	return ret, nil
}

func (c *Client) EntryTranAuIdPass(ctx context.Context, orderID string, amount, tax int) (*EntryTranAuIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	data.Set("Amount", strconv.Itoa(amount))
	data.Set("Tax", strconv.Itoa(tax))
	data.Set("JobCd", "CAPTURE")
	data.Set("PaymentType", "1")
	resp, err := c.Post(ctx, "/payment/EntryTranAu.idPass", data, false)
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	retVals, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	if retVals.Get("ErrCode") != "" {
		return nil, errors.NewErrorGMOPG(retVals.Get("ErrCode"), retVals.Get("ErrInfo"))
	}
	ret := &EntryTranAuIdPassResponse{}
	ret.AccessID = retVals["AccessID"][0]
	ret.AccessPass = retVals["AccessPass"][0]
	return ret, nil

}

func (c *Client) ExecTranAuIdPass(ctx context.Context, req *ExecTranAuIdPassReq) (*ExecTranAuIdPassResponse, error) {
	data := c.initRequestData(req)
	fmt.Printf("%+v\n", req)
	resp, err := c.Post(ctx, "/payment/ExecTranAu.idPass", data, false)
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	retVals, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	if retVals.Get("ErrCode") != "" {
		return nil, errors.NewErrorGMOPG(retVals.Get("ErrCode"), retVals.Get("ErrInfo"))
	}
	ret := &ExecTranAuIdPassResponse{}
	ret.OrderID = retVals.Get("OrderID")
	return ret, nil
}
