package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteJobsRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *JobBatchDelete `json:"body,omitempty"`
}

func (o BatchDeleteJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsRequest", string(data)}, " ")
}
