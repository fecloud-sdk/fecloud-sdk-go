package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o SecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"SecurityGroupRequest", string(data)}, " ")
}
