package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExpandInstanceNodeRequestBody struct {
	Num int32 `json:"num"`

	SubnetId *string `json:"subnet_id,omitempty"`

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ExpandInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeRequestBody", string(data)}, " ")
}
