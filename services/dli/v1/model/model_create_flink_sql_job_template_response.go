package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkSqlJobTemplateResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Template       *FlinkSqlJobTemplate `json:"template,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateFlinkSqlJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobTemplateResponse", string(data)}, " ")
}
