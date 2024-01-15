package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkJarJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Job            *FlinkSqlJobUpdateTime `json:"job,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateFlinkJarJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJarJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJarJobResponse", string(data)}, " ")
}
