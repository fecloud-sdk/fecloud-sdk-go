package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SlowlogResult struct {
	NodeName string `json:"node_name"`

	QuerySample string `json:"query_sample"`

	Type string `json:"type"`

	Time string `json:"time"`

	LockTime string `json:"lock_time"`

	RowsSent string `json:"rows_sent"`

	RowsExamined string `json:"rows_examined"`

	Database string `json:"database"`

	StartTime string `json:"start_time"`
}

func (o SlowlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogResult struct{}"
	}

	return strings.Join([]string{"SlowlogResult", string(data)}, " ")
}
