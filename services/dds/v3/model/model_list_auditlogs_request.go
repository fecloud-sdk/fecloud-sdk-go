package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAuditlogsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	NodeId *string `json:"node_id,omitempty"`

	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAuditlogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditlogsRequest", string(data)}, " ")
}
