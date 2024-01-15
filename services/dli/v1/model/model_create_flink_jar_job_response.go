package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkJarJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Job            *FlinkSqlJobStatusInfo `json:"job,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateFlinkJarJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarJobResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarJobResponse", string(data)}, " ")
}
