package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceVolumeRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ResizeInstanceVolumeRequestBody `json:"body,omitempty"`
}

func (o ResizeInstanceVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeRequest", string(data)}, " ")
}
