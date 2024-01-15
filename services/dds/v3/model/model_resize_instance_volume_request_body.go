package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceVolumeRequestBody struct {
	Volume *ResizeInstanceVolumeOption `json:"volume"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ResizeInstanceVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeRequestBody", string(data)}, " ")
}
