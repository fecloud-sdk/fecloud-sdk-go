package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowClusterRespClusterBrokers struct {
	Host *string `json:"host,omitempty"`

	Port *int32 `json:"port,omitempty"`

	BrokerId *string `json:"broker_id,omitempty"`

	IsController *bool `json:"is_controller,omitempty"`

	Version *string `json:"version,omitempty"`

	RegisterTime *int64 `json:"register_time,omitempty"`

	IsHealth *bool `json:"is_health,omitempty"`
}

func (o ShowClusterRespClusterBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRespClusterBrokers struct{}"
	}

	return strings.Join([]string{"ShowClusterRespClusterBrokers", string(data)}, " ")
}
