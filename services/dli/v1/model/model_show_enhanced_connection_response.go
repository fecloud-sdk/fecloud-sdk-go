package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEnhancedConnectionResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	AvailableQueueInfo *[]EnhancedConnectionResourceInfo `json:"available_queue_info,omitempty"`

	ElasticResourcePools *[]EnhancedConnectionResourceInfo `json:"elastic_resource_pools,omitempty"`

	DestVpcId *string `json:"dest_vpc_id,omitempty"`

	DestNetworkId *string `json:"dest_network_id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Hosts          *[]EnhancedConnectionHost `json:"hosts,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowEnhancedConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionResponse", string(data)}, " ")
}
