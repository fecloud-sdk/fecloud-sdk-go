package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchoverReplicaSetRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o SwitchoverReplicaSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverReplicaSetRequest struct{}"
	}

	return strings.Join([]string{"SwitchoverReplicaSetRequest", string(data)}, " ")
}
