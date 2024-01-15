package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkSqlJobRequest struct {
	JobId int64 `json:"job_id"`

	Body *UpdateFlinkSqlJobRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobRequest", string(data)}, " ")
}
