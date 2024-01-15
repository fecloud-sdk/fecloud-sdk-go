package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConnectionStatisticsRequest struct {
	InstanceId string `json:"instance_id"`

	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowConnectionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionStatisticsRequest", string(data)}, " ")
}
