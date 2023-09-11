package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteServerGroupMemberRequestBody This is a auto create Body Object
type DeleteServerGroupMemberRequestBody struct {
	RemoveMember *ServerGroupMember `json:"remove_member"`
}

func (o DeleteServerGroupMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupMemberRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupMemberRequestBody", string(data)}, " ")
}
