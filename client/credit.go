package client

import (
	"context"
	"github.com/y-tajiri/go-gmo-pg/errors"
	"io/ioutil"
	"net/url"
)

type (
	SaveMemberIdPassResponse struct {
		MemberID string
	}

	SaveCardIdPassReq struct {
		MemberID    string
		SeqMode     int
		DefaultFlag int
		Token       string
	}

	SaveCardIdPassResponse struct {
		CardSeq string
		CardNo  string
		Forward string
	}
	EntryTranIdPassReq struct {
		OrderID     string
		JobCd       string
	}
	EntryTranIdPassResponse struct {
		AccessID   string
		AccessPass string
	}
	ExecTranIdPassReq struct {
		AccessID         string
		AccessPass       string
		OrderID          string
		Method           int
		PayTimes         int
		MemberID         string
		SeqMode          int
		CardSeq          int
	}

)
func (c *Client) EntryTranIdPassCheck(ctx context.Context, req *EntryTranIdPassReq) (*EntryTranIdPassResponse, error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/EntryTran.idPass", data, false)
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
	ret := &EntryTranIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.AccessPass = retVals.Get("AccessPass")
	return ret, nil
}
func (c *Client) SaveMemberIdPass(ctx context.Context, memberID string, memberName string) (*SaveMemberIdPassResponse, error) {

	data := url.Values{}
	data.Set("MemberID", memberID)
	if memberName != "" {
		data.Set("MemberName", memberName)
	}
	resp, err := c.Post(ctx, "/payment/SaveMember.idPass", data, true)
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
	ret := &SaveMemberIdPassResponse{}
	ret.MemberID = retVals.Get("MemberID")
	return ret, nil

}

func (c *Client) SaveCreditIdPass(ctx context.Context, req *SaveCardIdPassReq) (*SaveCardIdPassResponse, error) {

	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/SaveCard.idPass", data, true)
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
	ret := &SaveCardIdPassResponse{}
	ret.CardSeq = retVals.Get("CardSeq")
	ret.CardNo = retVals.Get("CardNo")
	ret.Forward = retVals.Get("Forward")
	return ret, nil

}

func (c *Client) ExecTranIdPass(ctx context.Context, req *ExecTranIdPassReq) (error) {
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTran.idPass", data, false)
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
	return ret, nil
}
