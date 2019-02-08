package config

import (
	"gopkg.in/go-playground/validator.v9"
)

type (
	Config struct {
		EndPoint string `validate:"required"`
		Version  string
		ShopID   string `validate:"required"`
		ShopPass string `validate:"required"`
		SiteID   string `validate:"required"`
		SitePass string `validate:"required"`
	}
)

func New(endpoint, shopID, shopPass,siteID, sitePass string) (*Config, error) {
	c := &Config{
		EndPoint: endpoint,
		ShopID:   shopID,
		ShopPass: shopPass,
		SiteID:   siteID,
		SitePass: sitePass,
	}
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return nil, err
	}
	return c, err
}
