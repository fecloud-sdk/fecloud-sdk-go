package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteServerGroupMemberResponse Response Object
type DeleteServerGroupMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupMemberResponse", string(data)}, " ")
}
