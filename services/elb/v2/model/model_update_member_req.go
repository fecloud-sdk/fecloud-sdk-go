package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateMemberReq struct {
	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberReq struct{}"
	}

	return strings.Join([]string{"UpdateMemberReq", string(data)}, " ")
}
