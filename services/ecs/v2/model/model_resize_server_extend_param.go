package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeServerExtendParam struct {
	IsAutoPay *string `json:"isAutoPay,omitempty"`
}

func (o ResizeServerExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeServerExtendParam struct{}"
	}

	return strings.Join([]string{"ResizeServerExtendParam", string(data)}, " ")
}
