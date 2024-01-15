package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowDiskUsageRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowDiskUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiskUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowDiskUsageRequest", string(data)}, " ")
}
