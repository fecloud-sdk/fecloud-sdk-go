package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListGroupReplicationInfoRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListGroupReplicationInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupReplicationInfoRequest struct{}"
	}

	return strings.Join([]string{"ListGroupReplicationInfoRequest", string(data)}, " ")
}
