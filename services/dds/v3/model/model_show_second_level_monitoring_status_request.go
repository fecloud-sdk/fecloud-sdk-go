package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSecondLevelMonitoringStatusRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowSecondLevelMonitoringStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecondLevelMonitoringStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSecondLevelMonitoringStatusRequest", string(data)}, " ")
}
