package client

import (
	"context"
	"github.com/y-tajiri/go-gmo-pg/sjis"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/y-tajiri/go-gmo-pg/config"
)

type (
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}
	Client struct {
		url    *url.URL
		client HTTPClient
		config config.Config
	}
)

type ClientOption func(*Client)

func SetHTTPClient(cli HTTPClient) ClientOption {
	return func(c *Client) {
		c.client = cli
	}
}

func NewClient(config config.Config, opts ...ClientOption) (*Client, error) {

	u, err := url.ParseRequestURI(config.EndPoint)
	if err != nil {
		return nil, err
	}

	c := &Client{
		url:    u,
		client: http.DefaultClient,
		config: config,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func (c *Client) newRequest(ctx context.Context, spath string, data url.Values) (*http.Request, error) {
	u := *c.url
	u.Path = path.Join(c.url.Path, spath)
	data.Add("ShopID", c.config.ShopID)
	data.Add("ShopPass", c.config.ShopPass)
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

// Do sends a request and returns Response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// Post sends a POST request and returns Response
func (c *Client) Post(ctx context.Context, spath string, data url.Values) (*http.Response, error) {
	req, err := c.newRequest(ctx, spath, data)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) initRequestData(req interface{}) (data url.Values) {
	data = url.Values{}
	t := reflect.TypeOf(req)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	val := reflect.ValueOf(req)
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Type.Kind() {
		case reflect.String:
			s, err := sjis.ConvertUtf8ToSjis(val.Elem().Field(i).String())
			if err != nil {
				s = val.Elem().Field(i).String()
			}
			data.Set(t.Field(i).Name, s)
			break
		case reflect.Int:
			data.Set(
				t.Field(i).Name,
				strconv.Itoa(int(val.Elem().Field(i).Int())),
			)
			break
		}
	}
	return data
}
