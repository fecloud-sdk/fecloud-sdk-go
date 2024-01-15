package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateSinkTaskQuotaReq struct {
	SinkMaxTasks int32 `json:"sink_max_tasks"`
}

func (o UpdateSinkTaskQuotaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSinkTaskQuotaReq struct{}"
	}

	return strings.Join([]string{"UpdateSinkTaskQuotaReq", string(data)}, " ")
}
