package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPoolsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	HealthmonitorId *string `json:"healthmonitor_id,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	MemberAddress *string `json:"member_address,omitempty"`

	MemberDeviceId *string `json:"member_device_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
