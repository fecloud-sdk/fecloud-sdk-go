package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ServerGroupMember struct {
	InstanceUuid string `json:"instance_uuid"`
}

func (o ServerGroupMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerGroupMember struct{}"
	}

	return strings.Join([]string{"ServerGroupMember", string(data)}, " ")
}
