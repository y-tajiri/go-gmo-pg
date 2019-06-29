package client

import (
	"context"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/y-tajiri/go-gmo-pg/errors"
)

type (
	EntryTranSbIdPassResponse struct {
		AccessID   string
		AccessPass string
	}
	ExecTranSbIdPassReq struct {
		AccessID           string
		AccessPass         string
		OrderID            string
		ClientField1       string
		ClientField2       string
		ClientField3       string
		RetURL             string
		PaymentTermSec     int
		ChargeDay          int
		FirstMonthFreeFlag int
	}

	ExecTranSbIdPassResponse struct {
		AccessID       string
		Token          string
		StartURL       string
		StartLimitDate string
	}
)

func (c *Client) EntryTranSbIdPass(ctx context.Context, orderID string, amount, tax int) (*EntryTranSbIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	data.Set("Amount", strconv.Itoa(amount))
	data.Set("Tax", strconv.Itoa(tax))
	data.Set("JobCd", "CAPTURE")
	resp, err := c.Post(ctx, "/payment/EntryTranSb.idPass", data, false)
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
	ret := &EntryTranSbIdPassResponse{}
	ret.AccessID = retVals["AccessID"][0]
	ret.AccessPass = retVals["AccessPass"][0]
	return ret, nil

}

func (c *Client) ExecTranSbIdPass(ctx context.Context, req *ExecTranSbIdPassReq) (*ExecTranSbIdPassResponse, error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTranSb.idPass", data, false)
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
	ret := &ExecTranSbIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.StartLimitDate = retVals.Get("StartLimitDate")
	ret.StartURL = retVals.Get("StartURL")
	ret.Token = retVals.Get("Token")
	return ret, nil
}
