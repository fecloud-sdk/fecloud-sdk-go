package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkSqlJobTemplateRequest struct {
	Body *CreateFlinkSqlJobTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobTemplateRequest", string(data)}, " ")
}
