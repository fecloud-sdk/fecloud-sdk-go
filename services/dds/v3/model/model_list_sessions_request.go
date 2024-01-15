package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSessionsRequest struct {
	NodeId string `json:"node_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PlanSummary *string `json:"plan_summary,omitempty"`

	Type *string `json:"type,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	CostTime *int32 `json:"cost_time,omitempty"`
}

func (o ListSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsRequest struct{}"
	}

	return strings.Join([]string{"ListSessionsRequest", string(data)}, " ")
}
