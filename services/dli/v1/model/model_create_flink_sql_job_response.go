package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkSqlJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Job            *FlinkSqlJobStatusInfo `json:"job,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateFlinkSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobResponse", string(data)}, " ")
}
