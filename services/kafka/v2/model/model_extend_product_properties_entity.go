package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExtendProductPropertiesEntity struct {
	MaxPartitionPerBroker *string `json:"max_partition_per_broker,omitempty"`

	MaxBroker *string `json:"max_broker,omitempty"`

	MaxStoragePerNode *string `json:"max_storage_per_node,omitempty"`

	MaxConsumerPerBroker *string `json:"max_consumer_per_broker,omitempty"`

	MinBroker *string `json:"min_broker,omitempty"`

	MaxBandwidthPerBroker *string `json:"max_bandwidth_per_broker,omitempty"`

	MinStoragePerNode *string `json:"min_storage_per_node,omitempty"`

	MaxTpsPerBroker *string `json:"max_tps_per_broker,omitempty"`

	ProductAlias *string `json:"product_alias,omitempty"`
}

func (o ExtendProductPropertiesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendProductPropertiesEntity struct{}"
	}

	return strings.Join([]string{"ExtendProductPropertiesEntity", string(data)}, " ")
}
