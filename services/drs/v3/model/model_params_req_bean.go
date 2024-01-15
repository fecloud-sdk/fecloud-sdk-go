package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ParamsReqBean struct {
	Key string `json:"key"`

	TargetValue string `json:"target_value"`
}

func (o ParamsReqBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamsReqBean struct{}"
	}

	return strings.Join([]string{"ParamsReqBean", string(data)}, " ")
}
