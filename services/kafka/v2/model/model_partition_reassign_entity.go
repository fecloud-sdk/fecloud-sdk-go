package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PartitionReassignEntity struct {
	Topic string `json:"topic"`

	Brokers *[]int32 `json:"brokers,omitempty"`

	ReplicationFactor *int32 `json:"replication_factor,omitempty"`

	Assignment *[]TopicAssignment `json:"assignment,omitempty"`
}

func (o PartitionReassignEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionReassignEntity struct{}"
	}

	return strings.Join([]string{"PartitionReassignEntity", string(data)}, " ")
}
