package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteFlinkSqlJobTemplateRequest struct {
	TemplateId int64 `json:"template_id"`
}

func (o DeleteFlinkSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlinkSqlJobTemplateRequest", string(data)}, " ")
}
