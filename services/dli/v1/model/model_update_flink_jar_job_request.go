package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateFlinkJarJobRequest struct {
	JobId int64 `json:"job_id"`

	Body *UpdateFlinkJarJobRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkJarJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJarJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJarJobRequest", string(data)}, " ")
}
