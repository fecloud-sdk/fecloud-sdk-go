package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationAppliedHistoryRequest struct {
	ConfigId string `json:"config_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowConfigurationAppliedHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAppliedHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAppliedHistoryRequest", string(data)}, " ")
}
