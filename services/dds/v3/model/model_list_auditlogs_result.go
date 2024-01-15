package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAuditlogsResult struct {
	NodeId string `json:"node_id"`

	Id string `json:"id"`

	Name string `json:"name"`

	Size int64 `json:"size"`

	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`
}

func (o ListAuditlogsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogsResult struct{}"
	}

	return strings.Join([]string{"ListAuditlogsResult", string(data)}, " ")
}
