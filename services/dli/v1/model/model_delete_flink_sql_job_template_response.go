package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteFlinkSqlJobTemplateResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Template       *TemplateId `json:"template,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o DeleteFlinkSqlJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkSqlJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlinkSqlJobTemplateResponse", string(data)}, " ")
}
