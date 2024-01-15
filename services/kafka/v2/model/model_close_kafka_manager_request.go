package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CloseKafkaManagerRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o CloseKafkaManagerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseKafkaManagerRequest struct{}"
	}

	return strings.Join([]string{"CloseKafkaManagerRequest", string(data)}, " ")
}
