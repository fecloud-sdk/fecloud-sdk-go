package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkSqlJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Job            *FlinkSqlJobUpdateTime `json:"job,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateFlinkSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobResponse", string(data)}, " ")
}
