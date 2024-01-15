package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueueInfo struct {
	QueueId *int64 `json:"queue_id,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	Description *string `json:"description,omitempty"`

	Owner *string `json:"owner,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	QueueType *string `json:"queue_type,omitempty"`

	CuCount *int32 `json:"cu_count,omitempty"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	CidrInMgntsubnet *string `json:"cidr_in_mgntsubnet,omitempty"`

	CidrInSubnet *string `json:"cidr_in_subnet,omitempty"`

	ResourceMode *int32 `json:"resource_mode,omitempty"`

	Platform *string `json:"platform,omitempty"`

	IsRestarting *bool `json:"is_restarting,omitempty"`

	Labels *string `json:"labels,omitempty"`

	Feature *string `json:"feature,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`

	CuSpec *int32 `json:"cu_spec,omitempty"`

	CuScaleOutLimit *int32 `json:"cu_scale_out_limit,omitempty"`

	CuScaleInLimit *int32 `json:"cu_scale_in_limit,omitempty"`

	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`
}

func (o QueueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueInfo struct{}"
	}

	return strings.Join([]string{"QueueInfo", string(data)}, " ")
}
