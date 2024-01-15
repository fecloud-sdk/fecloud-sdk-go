package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddReadonlyNodeRequestBody struct {
	SpecCode string `json:"spec_code"`

	Num int32 `json:"num"`

	Delay *int32 `json:"delay,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o AddReadonlyNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddReadonlyNodeRequestBody struct{}"
	}

	return strings.Join([]string{"AddReadonlyNodeRequestBody", string(data)}, " ")
}
