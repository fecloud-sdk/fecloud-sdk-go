package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PrecheckFailSubJobVo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	CheckResult *string `json:"check_result,omitempty"`
}

func (o PrecheckFailSubJobVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckFailSubJobVo struct{}"
	}

	return strings.Join([]string{"PrecheckFailSubJobVo", string(data)}, " ")
}
