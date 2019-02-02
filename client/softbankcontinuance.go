package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"
	"strconv"

	"github.com/y-tajiri/go-gmo-pg/errors"
)

type (
	EntryTranSbContinuanceIdPassResponse struct {
		AccessID string
		AccessPass string
	}
	ExecTranSbContinuanceIdPassReq struct {
		AccessID       string
		AccessPass     string
		OrderID        string
		ClientField1   string
		ClientField2   string
		ClientField3   string
		RetURL         string
		PaymentTermSec int
		ChargeDay      int
		FirstMonthFreeFlag int
	}

	ExecTranSbContinuanceIdPassResponse struct{
		AccessID string
		Token string
		StartURL string
		StartLimitDate string
	}
)

func (c *Client) initRequestData(req interface{}) (data url.Values) {
	data = url.Values{}
	t := reflect.TypeOf(req)
	fmt.Printf("sss %+v\n", t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	val := reflect.ValueOf(req)
	for i := 0; i < t.NumField(); i++ {
		switch(t.Field(i).Type.Kind()){
		case reflect.String:
			data.Set(t.Field(i).Name, val.Elem().Field(i).String())
			break
		case reflect.Int:
			data.Set(
				t.Field(i).Name,
				strconv.Itoa(int(val.Elem().Field(i).Int())),
			)
			break
		}
	}
	fmt.Printf("sss %+v\n", val)
	return data
}
func (c *Client) EntryTranSbContinuanceIdPass(ctx context.Context, orderID string, amount, tax int)  (*EntryTranSbContinuanceIdPassResponse, error) {

	data := url.Values{}
	data.Set("OrderID", orderID)
	data.Set("Amount", strconv.Itoa(amount))
	data.Set("Tax", strconv.Itoa(tax))
	resp, err := c.Post(ctx, "/payment/EntryTranSbContinuance.idPass" , data)
	if err != nil {
		return nil, err
	}
	b,_ := ioutil.ReadAll(resp.Body)
	retVals, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	if retVals.Get("ErrCode") != "" {
		return nil, errors.NewErrorGMOPG(retVals.Get("ErrCode"), retVals.Get("ErrInfo"))
	}
	ret := &EntryTranSbContinuanceIdPassResponse{}
	ret.AccessID = retVals["AccessID"][0]
	ret.AccessPass = retVals["AccessPass"][0]
	return ret,nil

}

func (c *Client) ExecTranSbContinuanceIdPass(ctx context.Context, req *ExecTranSbContinuanceIdPassReq)(*ExecTranSbContinuanceIdPassResponse, error){
	data := c.initRequestData(req)
	resp, err := c.Post(ctx, "/payment/ExecTranSbContinuance.idPass" , data)
	if err != nil {
		return nil, err
	}
	b,_ := ioutil.ReadAll(resp.Body)
	retVals, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	if retVals.Get("ErrCode")  != "" {
		return nil, errors.NewErrorGMOPG(retVals.Get("ErrCode"), retVals.Get("ErrInfo"))
	}
	ret := &ExecTranSbContinuanceIdPassResponse{}
	ret.AccessID = retVals.Get("AccessID")
	ret.StartLimitDate = retVals.Get("StartLimitDate")
	ret.StartURL = retVals.Get("StartURL")
	ret.Token = retVals.Get("Token")
	return ret,nil
}
