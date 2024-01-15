package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociateQueueToEnhancedConnectionRequestBody struct {
	Queues *[]string `json:"queues,omitempty"`

	ElasticResourcePools *[]string `json:"elastic_resource_pools,omitempty"`
}

func (o AssociateQueueToEnhancedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateQueueToEnhancedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateQueueToEnhancedConnectionRequestBody", string(data)}, " ")
}
