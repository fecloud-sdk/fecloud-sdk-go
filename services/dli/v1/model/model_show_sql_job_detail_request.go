package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlJobDetailRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowSqlJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobDetailRequest", string(data)}, " ")
}
