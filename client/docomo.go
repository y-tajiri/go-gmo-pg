package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/y-tajiri/go-gmo-pg/errors"
)

type (
	EntryTranDocomoAcceptIdPassResponse struct{
		AccessID   string
		AccessPass string
	}
	ExecTranDocomoAcceptIdPassReq struct{
		AccessID           string
		AccessPass         string
		OrderID            string
		RetURL             string
		ClientField1       string
		ClientField2       string
		ClientField3       string
		PaymentTermSec     int
		ChargeDay          int
		FirstMonthFreeFlag int
	}
	ExecTranDocomoAcceptIdPassResponse struct{
		AccessID       string
		Token          string
		StartURL   string
	}
	EntryTranDocomoIdPassResponse struct {
		AccessID   string
		AccessPass string
	}
	ExecTranDocomoIdPassReq struct {
		AccessID           string
		AccessPass         string
		OrderID            string
		RetURL             string
		ClientField1       string
		ClientField2       string
		ClientField3       string
		PaymentTermSec     int
	}

	ExecTranDocomoIdPassResponse struct {
		AccessID string
		Token string
		StartURL string
	}
)

func (c *Client) EntryTranDocomoAcceptIdPass(ctx context.Context, orderID string) (*EntryTranDocomoAcceptIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	resp, err := c.Post(ctx, "/payment/EntryTranDocomoAccept.idPass", data, false)
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
	ret := &EntryTranDocomoAcceptIdPassResponse{}
	ret.AccessID = retVals["AccessID"][0]
	ret.AccessPass = retVals["AccessPass"][0]
	return ret, nil

}
func (c *Client)ExecTranDocomoAcceptIdPass(ctx context.Context, req *ExecTranDocomoAcceptIdPassReq) (*ExecTranDocomoAcceptIdPassResponse, error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTranDocomoAccept.idPass", data, false)
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
	ret := &ExecTranDocomoAcceptIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.Token = retVals.Get("Token")
	ret.StartURL = retVals.Get("StartURL")
	return ret, nil
}

func (c *Client) EntryTranDocomoIdPass(ctx context.Context, orderID string, amount, tax int) (*EntryTranDocomoIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	data.Set("Amount", strconv.Itoa(amount))
	data.Set("Tax", strconv.Itoa(tax))
	data.Set("JobCd", "CAPTURE")
	data.Set("PaymentType", "1")
	resp, err := c.Post(ctx, "/payment/EntryTranDocomo.idPass", data, false)
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
	ret := &EntryTranDocomoIdPassResponse{}
	ret.AccessID = retVals["AccessID"][0]
	ret.AccessPass = retVals["AccessPass"][0]
	return ret, nil

}

func (c *Client) ExecTranDocomoIdPass(ctx context.Context, req *ExecTranDocomoIdPassReq) (*ExecTranDocomoIdPassResponse, error) {
	data := c.initRequestData(req)
	fmt.Printf("")
	resp, err := c.Post(ctx, "/payment/ExecTranDocomo.idPass", data, false)
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
	ret := &ExecTranDocomoIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.Token = retVals.Get("Token")
	ret.StartURL = retVals.Get("StartURL")
	return ret, nil
}
