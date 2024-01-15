package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetReplicaReq struct {
	Partitions *[]ResetReplicaReqPartitions `json:"partitions,omitempty"`
}

func (o ResetReplicaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetReplicaReq struct{}"
	}

	return strings.Join([]string{"ResetReplicaReq", string(data)}, " ")
}
