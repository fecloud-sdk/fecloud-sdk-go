package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSlowLogsRequest struct {
	InstanceId string `json:"instance_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	NodeId *string `json:"node_id,omitempty"`

	Type *string `json:"type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogsRequest", string(data)}, " ")
}
