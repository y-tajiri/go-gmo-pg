package client

import (
	"context"
	"github.com/y-tajiri/go-gmo-pg/errors"
	"io/ioutil"
	"net/url"
)

type (
	EntryTranAuContinuanceIdPassResponse struct {
		AccessID   string
		AccessPass string
	}
	EntryTranAuContinuanceIdPassReq struct {
		OrderID     string
		Amount      int
		Tax         int
		FirstAmount int
		FirstTax    int
	}

	ExecTranAuContinuanceIdPassReq struct {
		AccessID         string
		AccessPass       string
		OrderID          string
		ClientField1     string
		ClientField2     string
		ClientField3     string
		RetURL           string
		Commodity        string
		AccountTimingKbn string
		AccountTiming    int
		FirstAccountDate string
		PaymentTermSec   int
		ServiceName      string
		ServiceTel       string
	}

	ExecTranAuContinuanceIdPassResponse struct {
		AccessID       string
		Token          string
		StartURL       string
		StartLimitDate string
	}
)

const (
	AccountTimingKbnCustom  = "01"
	AccountTimingKbnLastDay = "02"
)

func (c *Client) EntryTranAuContinuanceIdPass(ctx context.Context, req *EntryTranAuContinuanceIdPassReq) (*EntryTranAuContinuanceIdPassResponse, error) {

	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/EntryTranAuContinuance.idPass", data, false)
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
	ret := &EntryTranAuContinuanceIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.AccessPass = retVals.Get("AccessPass")
	return ret, nil

}

func (c *Client) ExecTranAuContinuanceIdPass(ctx context.Context, req *ExecTranAuContinuanceIdPassReq) (*ExecTranAuContinuanceIdPassResponse, error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTranAuContinuance.idPass", data, false)
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
	ret := &ExecTranAuContinuanceIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.StartLimitDate = retVals.Get("StartLimitDate")
	ret.StartURL = retVals.Get("StartURL")
	ret.Token = retVals.Get("Token")
	return ret, nil
}
