package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AddServerGroupMemberResponse Response Object
type AddServerGroupMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddServerGroupMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerGroupMemberResponse struct{}"
	}

	return strings.Join([]string{"AddServerGroupMemberResponse", string(data)}, " ")
}
