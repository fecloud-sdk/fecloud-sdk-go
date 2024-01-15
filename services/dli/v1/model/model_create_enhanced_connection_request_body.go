package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateEnhancedConnectionRequestBody struct {
	Name string `json:"name"`

	DestVpcId string `json:"dest_vpc_id"`

	DestNetworkId string `json:"dest_network_id"`

	Queues *[]string `json:"queues,omitempty"`

	Hosts *[]EnhancedConnectionHost `json:"hosts,omitempty"`

	RoutetableId *string `json:"routetable_id,omitempty"`

	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o CreateEnhancedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionRequestBody", string(data)}, " ")
}
