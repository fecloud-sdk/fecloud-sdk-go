package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RoutetableInfoErrorDetial struct {
	Id *string `json:"id,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o RoutetableInfoErrorDetial) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableInfoErrorDetial struct{}"
	}

	return strings.Join([]string{"RoutetableInfoErrorDetial", string(data)}, " ")
}
