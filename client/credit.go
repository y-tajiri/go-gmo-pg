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
)

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
