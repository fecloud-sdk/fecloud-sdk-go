package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceUsersResponse struct {
	Users          *[]ShowInstanceUsersEntity `json:"users,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowInstanceUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersResponse", string(data)}, " ")
}
