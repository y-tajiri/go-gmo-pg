package client

import (
	"context"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/y-tajiri/go-gmo-pg/errors"
)

type (
	EntryTranDocomoIdPassResponse struct {
		AccessID   string
		AccessPass string
	}
	ExecTranDocomoIdPassReq struct {
		AccessID           string
		AccessPass         string
		OrderID            string
		ClientField1       string
		ClientField2       string
		ClientField3       string
 		PaymentTermSec     int
		ChargeDay          int
		FirstMonthFreeFlag int
	}

	ExecTranDocomoIdPassResponse struct {
		AccessID       string
		Token          string
	}
)

func (c *Client) EntryTranDocomoIdPass(ctx context.Context, orderID string, amount, tax int) (*EntryTranDocomoIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	data.Set("Amount", strconv.Itoa(amount))
	data.Set("Tax", strconv.Itoa(tax))
	data.Set("JobCd", "CAPTURE")
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
	return ret, nil
}
