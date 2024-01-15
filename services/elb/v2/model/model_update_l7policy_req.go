package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7policyReq struct {
	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`
}

func (o UpdateL7policyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policyReq struct{}"
	}

	return strings.Join([]string{"UpdateL7policyReq", string(data)}, " ")
}
