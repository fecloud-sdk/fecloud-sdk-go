package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SlowLogDetail struct {
	NodeName string `json:"node_name"`

	NodeId string `json:"node_id"`

	WholeMessage string `json:"whole_message"`

	OperateType string `json:"operate_type"`

	CostTime int32 `json:"cost_time"`

	LockTime int32 `json:"lock_time"`

	DocsReturned int32 `json:"docs_returned"`

	DocsScanned int32 `json:"docs_scanned"`

	Database string `json:"database"`

	Collection string `json:"collection"`

	LogTime string `json:"log_time"`

	LineNum string `json:"line_num"`
}

func (o SlowLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogDetail struct{}"
	}

	return strings.Join([]string{"SlowLogDetail", string(data)}, " ")
}
