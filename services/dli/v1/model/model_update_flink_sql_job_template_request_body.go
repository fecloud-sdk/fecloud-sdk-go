package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkSqlJobTemplateRequestBody struct {
	Name *string `json:"name,omitempty"`

	Desc *string `json:"desc,omitempty"`

	SqlBody *string `json:"sql_body,omitempty"`
}

func (o UpdateFlinkSqlJobTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobTemplateRequestBody", string(data)}, " ")
}
