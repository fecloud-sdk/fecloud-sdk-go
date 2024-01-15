package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSparkJobsRequest struct {
	ClusterName *string `json:"cluster_name,omitempty"`

	End *int64 `json:"end,omitempty"`

	From *int32 `json:"from,omitempty"`

	JobName *string `json:"job-name,omitempty"`

	JobId *string `json:"job-id,omitempty"`

	Order *string `json:"order,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Start *int64 `json:"start,omitempty"`

	State *string `json:"state,omitempty"`
}

func (o ListSparkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSparkJobsRequest", string(data)}, " ")
}
