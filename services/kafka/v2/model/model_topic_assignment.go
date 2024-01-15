package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TopicAssignment struct {
	Partition *int32 `json:"partition,omitempty"`

	PartitionBrokers *[]int32 `json:"partition_brokers,omitempty"`
}

func (o TopicAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicAssignment struct{}"
	}

	return strings.Join([]string{"TopicAssignment", string(data)}, " ")
}
