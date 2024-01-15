package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlinkSqlJobTemplatesResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	TemplateList   *FlinkSqlJobTemplateList `json:"template_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFlinkSqlJobTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkSqlJobTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkSqlJobTemplatesResponse", string(data)}, " ")
}
