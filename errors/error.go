package errors

import (
	"errors"
	"fmt"
)

var (
	InvalidConfigError = errors.New("Invalid config error")
)

type (
	ErrorGMOPG struct{
		ErrCode string
		ErrInfo string
	}
)

func (e *ErrorGMOPG) Error() string {
	return fmt.Sprintf("gmo error:%s %s", e.ErrCode, e.ErrInfo)
}

func NewErrorGMOPG(errCode, errInfo string) error {
	return &ErrorGMOPG{ErrCode: errCode, ErrInfo: errInfo}
}