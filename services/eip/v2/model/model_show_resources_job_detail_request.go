package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowResourcesJobDetailRequest Request Object
type ShowResourcesJobDetailRequest struct {

	// 批量操作返回的JOB ID
	JobId string `json:"job_id"`
}

func (o ShowResourcesJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowResourcesJobDetailRequest", string(data)}, " ")
}
