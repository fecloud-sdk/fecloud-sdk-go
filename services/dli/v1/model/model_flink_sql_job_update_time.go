package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkSqlJobUpdateTime struct {
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o FlinkSqlJobUpdateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobUpdateTime struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobUpdateTime", string(data)}, " ")
}
