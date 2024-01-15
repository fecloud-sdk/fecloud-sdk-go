package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DisassociateQueueFromEnhancedConnectionRequestBody struct {
	Queues *[]string `json:"queues,omitempty"`

	ElasticResourcePools *[]string `json:"elastic_resource_pools,omitempty"`
}

func (o DisassociateQueueFromEnhancedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateQueueFromEnhancedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateQueueFromEnhancedConnectionRequestBody", string(data)}, " ")
}
