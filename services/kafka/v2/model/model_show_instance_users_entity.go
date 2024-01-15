package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceUsersEntity struct {
	UserName *string `json:"user_name,omitempty"`

	UserDesc *string `json:"user_desc,omitempty"`

	Role *string `json:"role,omitempty"`

	DefaultApp *bool `json:"default_app,omitempty"`

	CreatedTime *int64 `json:"created_time,omitempty"`
}

func (o ShowInstanceUsersEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersEntity struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersEntity", string(data)}, " ")
}
