package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelSqlJobRequest struct {
	JobId string `json:"job_id"`
}

func (o CancelSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSqlJobRequest struct{}"
	}

	return strings.Join([]string{"CancelSqlJobRequest", string(data)}, " ")
}
