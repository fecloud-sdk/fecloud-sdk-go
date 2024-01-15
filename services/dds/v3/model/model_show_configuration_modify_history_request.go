package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationModifyHistoryRequest struct {
	ConfigId string `json:"config_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowConfigurationModifyHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationModifyHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationModifyHistoryRequest", string(data)}, " ")
}
