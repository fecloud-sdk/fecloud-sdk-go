package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateFlinkSqlJobRequest struct {
	Body *CreateFlinkSqlJobRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobRequest", string(data)}, " ")
}
