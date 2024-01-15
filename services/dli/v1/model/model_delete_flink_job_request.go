package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteFlinkJobRequest struct {
	JobId int64 `json:"job_id"`
}

func (o DeleteFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlinkJobRequest", string(data)}, " ")
}
