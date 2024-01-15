package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlJobStatusRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowSqlJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobStatusRequest", string(data)}, " ")
}
