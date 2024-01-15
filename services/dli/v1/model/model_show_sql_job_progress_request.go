package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlJobProgressRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowSqlJobProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobProgressRequest", string(data)}, " ")
}
