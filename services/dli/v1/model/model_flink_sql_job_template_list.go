package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkSqlJobTemplateList struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Templates *[]FlinkSqlJobTemplateInfo `json:"templates,omitempty"`
}

func (o FlinkSqlJobTemplateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobTemplateList struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobTemplateList", string(data)}, " ")
}
