package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EnhancedConnectionInfo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	AvailableQueueInfo *[]EnhancedConnectionResourceInfo `json:"available_queue_info,omitempty"`

	ElasticResourcePools *[]EnhancedConnectionResourceInfo `json:"elastic_resource_pools,omitempty"`

	DestVpcId *string `json:"dest_vpc_id,omitempty"`

	DestNetworkId *string `json:"dest_network_id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Hosts *[]EnhancedConnectionHost `json:"hosts,omitempty"`

	IsPrivis *bool `json:"isPrivis,omitempty"`
}

func (o EnhancedConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionInfo struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionInfo", string(data)}, " ")
}
