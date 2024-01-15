package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceOption struct {
	TargetSpecCode string `json:"target_spec_code"`
}

func (o ResizeInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceOption struct{}"
	}

	return strings.Join([]string{"ResizeInstanceOption", string(data)}, " ")
}
