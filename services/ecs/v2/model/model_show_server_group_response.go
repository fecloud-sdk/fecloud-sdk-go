package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowServerGroupResponse Response Object
type ShowServerGroupResponse struct {
	ServerGroup    *ShowServerGroupResult `json:"server_group,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowServerGroupResponse", string(data)}, " ")
}
