package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateQueueRequestBody struct {
	QueueName string `json:"queue_name"`

	QueueType *string `json:"queue_type,omitempty"`

	Description *string `json:"description,omitempty"`

	CuCount int32 `json:"cu_count"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Platform *string `json:"platform,omitempty"`

	ResourceMode *int32 `json:"resource_mode,omitempty"`

	Labels *[]interface{} `json:"labels,omitempty"`

	Feature *string `json:"feature,omitempty"`

	Tags *[]TmsTag `json:"tags,omitempty"`

	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`
}

func (o CreateQueueRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequestBody struct{}"
	}

	return strings.Join([]string{"CreateQueueRequestBody", string(data)}, " ")
}
