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
	}
)

func New(endpoint, shopID, shopPass string) (*Config, error) {
	c := &Config{
		EndPoint: endpoint,
		ShopID:   shopID,
		ShopPass: shopPass,
	}
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return nil, err
	}
	return c, err
}
