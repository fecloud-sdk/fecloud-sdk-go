package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SlowlogResult struct {
	Time string `json:"time"`

	Database string `json:"database"`

	QuerySample string `json:"query_sample"`

	Type string `json:"type"`

	StartTime string `json:"start_time"`
}

func (o SlowlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogResult struct{}"
	}

	return strings.Join([]string{"SlowlogResult", string(data)}, " ")
}
