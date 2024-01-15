package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBigkeyScanTaskDetailsRequest struct {
	InstanceId string `json:"instance_id"`

	BigkeyId string `json:"bigkey_id"`
}

func (o ShowBigkeyScanTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigkeyScanTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowBigkeyScanTaskDetailsRequest", string(data)}, " ")
}
