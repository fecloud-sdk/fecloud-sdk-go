package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkSqlJobTemplateResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFlinkSqlJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobTemplateResponse", string(data)}, " ")
}
