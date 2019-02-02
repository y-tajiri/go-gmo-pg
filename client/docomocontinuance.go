package client

import (
	"context"
	"github.com/y-tajiri/go-gmo-pg/errors"
	"io/ioutil"
	"net/url"
)


type (
	EntryTranDocomoContinuanceIdPassResponse struct {
		AccessID   string
		AccessPass string
	}

	EntryTranDocomoContinuanceIdPassReq struct {
		OrderID     string
		Amount      int
		Tax         int
		FirstAmount int
		FirstTax    int
	}

	ExecTranDocomoContinuanceIdPassReq struct {
		AccessID         string
		AccessPass       string
		OrderID          string
		ClientField1     string
		ClientField2     string
		ClientField3     string
		FirstMonthFreeFlag int
		ConfirmBaseDate int
		RetURL           string
		PaymentTermSec   int
	}

	ExecTranDocomoContinuanceIdPassResponse struct {
		AccessID       string
		Token          string
		StartURL       string
		StartLimitDate string
	}
)


func (c *Client) EntryTranDocomoContinuanceIdPass(ctx context.Context, req *EntryTranDocomoContinuanceIdPassReq) (*EntryTranDocomoContinuanceIdPassResponse, error) {

	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/EntryTranDocomoContinuance.idPass", data)
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
	ret := &EntryTranDocomoContinuanceIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.AccessPass = retVals.Get("AccessPass")
	return ret, nil

}

func (c *Client) ExecTranDocomoContinuanceIdPass(ctx context.Context, req *ExecTranDocomoContinuanceIdPassReq) (*ExecTranDocomoContinuanceIdPassResponse, error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTranDocomoContinuance.idPass", data)
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
	ret := &ExecTranDocomoContinuanceIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.StartLimitDate = retVals.Get("StartLimitDate")
	ret.StartURL = retVals.Get("StartURL")
	ret.Token = retVals.Get("Token")
	return ret, nil
}
