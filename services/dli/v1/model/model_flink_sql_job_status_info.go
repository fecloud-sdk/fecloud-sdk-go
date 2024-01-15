package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlinkSqlJobStatusInfo struct {
	JobId *int64 `json:"job_id,omitempty"`

	StatusName *string `json:"status_name,omitempty"`

	StatusDesc *string `json:"status_desc,omitempty"`
}

func (o FlinkSqlJobStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobStatusInfo struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobStatusInfo", string(data)}, " ")
}
