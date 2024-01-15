package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowQueueResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	QueueId *int64 `json:"queue_id,omitempty"`

	QueueName *string `json:"queueName,omitempty"`

	Description *string `json:"description,omitempty"`

	Owner *string `json:"owner,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	QueueType *string `json:"queueType,omitempty"`

	CuCount *int32 `json:"cuCount,omitempty"`

	ChargingMode *int32 `json:"chargingMode,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	ResourceMode *int32 `json:"resource_mode,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`

	CuSpec *int32 `json:"cu_spec,omitempty"`

	CuScaleOutLimit *int32 `json:"cu_scale_out_limit,omitempty"`

	CuScaleInLimit *int32 `json:"cu_scale_in_limit,omitempty"`

	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o ShowQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueResponse struct{}"
	}

	return strings.Join([]string{"ShowQueueResponse", string(data)}, " ")
}
