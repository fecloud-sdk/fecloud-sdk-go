package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowClusterRespCluster struct {
	Controller *string `json:"controller,omitempty"`

	Brokers *[]ShowClusterRespClusterBrokers `json:"brokers,omitempty"`

	TopicsCount *int32 `json:"topics_count,omitempty"`

	PartitionsCount *int32 `json:"partitions_count,omitempty"`

	OnlinePartitionsCount *int32 `json:"online_partitions_count,omitempty"`

	ReplicasCount *int32 `json:"replicas_count,omitempty"`

	IsrReplicasCount *int32 `json:"isr_replicas_count,omitempty"`

	ConsumersCount *int32 `json:"consumers_count,omitempty"`
}

func (o ShowClusterRespCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRespCluster struct{}"
	}

	return strings.Join([]string{"ShowClusterRespCluster", string(data)}, " ")
}
