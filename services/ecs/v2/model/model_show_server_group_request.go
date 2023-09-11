package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowServerGroupRequest Request Object
type ShowServerGroupRequest struct {

	// 弹性云服务器组UUID。
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowServerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowServerGroupRequest", string(data)}, " ")
}
