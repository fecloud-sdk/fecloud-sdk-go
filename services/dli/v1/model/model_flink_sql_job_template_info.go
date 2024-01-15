package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkSqlJobTemplateInfo struct {
	TemplateId *int32 `json:"template_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Desc *string `json:"desc,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	SqlBody *string `json:"sql_body,omitempty"`

	JobType *string `json:"job_type,omitempty"`
}

func (o FlinkSqlJobTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobTemplateInfo struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobTemplateInfo", string(data)}, " ")
}
