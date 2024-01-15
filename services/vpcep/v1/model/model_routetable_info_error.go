package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RoutetableInfoError struct {
	BindFailed *[]RoutetableInfoErrorDetial `json:"bind_failed,omitempty"`

	UnbindFailed *[]RoutetableInfoErrorDetial `json:"unbind_failed,omitempty"`
}

func (o RoutetableInfoError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableInfoError struct{}"
	}

	return strings.Join([]string{"RoutetableInfoError", string(data)}, " ")
}
