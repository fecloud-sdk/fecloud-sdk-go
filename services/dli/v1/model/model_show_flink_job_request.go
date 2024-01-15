package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowFlinkJobRequest struct {
	JobId int64 `json:"job_id"`
}

func (o ShowFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobRequest", string(data)}, " ")
}
