package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkJobInfoList struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Jobs *[]FlinkJobInfo `json:"jobs,omitempty"`
}

func (o FlinkJobInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobInfoList struct{}"
	}

	return strings.Join([]string{"FlinkJobInfoList", string(data)}, " ")
}
