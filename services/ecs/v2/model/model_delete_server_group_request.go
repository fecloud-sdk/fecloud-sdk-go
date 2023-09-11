package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteServerGroupRequest Request Object
type DeleteServerGroupRequest struct {

	// 弹性云服务器组UUID。
	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteServerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupRequest", string(data)}, " ")
}
