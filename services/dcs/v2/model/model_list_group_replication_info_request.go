package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListGroupReplicationInfoRequest Request Object
type ListGroupReplicationInfoRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListGroupReplicationInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupReplicationInfoRequest struct{}"
	}

	return strings.Join([]string{"ListGroupReplicationInfoRequest", string(data)}, " ")
}
