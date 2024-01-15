package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkSqlJobTemplate struct {
	TemplateId *int64 `json:"template_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Desc *string `json:"desc,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	JobType *string `json:"job_type,omitempty"`
}

func (o FlinkSqlJobTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobTemplate struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobTemplate", string(data)}, " ")
}
