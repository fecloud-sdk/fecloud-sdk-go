package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StructProcessResp struct {
	CreateTime *string `json:"create_time,omitempty"`

	Result *[]StructProcessVo `json:"result,omitempty"`
}

func (o StructProcessResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructProcessResp struct{}"
	}

	return strings.Join([]string{"StructProcessResp", string(data)}, " ")
}
