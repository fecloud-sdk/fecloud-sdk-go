package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowUpgradeDurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowUpgradeDurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeDurationRequest struct{}"
	}

	return strings.Join([]string{"ShowUpgradeDurationRequest", string(data)}, " ")
}
