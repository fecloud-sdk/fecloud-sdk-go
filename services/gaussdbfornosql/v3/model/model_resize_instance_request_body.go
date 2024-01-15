package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceRequestBody struct {
	Resize *ResizeInstanceOption `json:"resize"`

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ResizeInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequestBody", string(data)}, " ")
}
