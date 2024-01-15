package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7ruleReq struct {
	CompareType *string `json:"compare_type,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Invert *bool `json:"invert,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o UpdateL7ruleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7ruleReq struct{}"
	}

	return strings.Join([]string{"UpdateL7ruleReq", string(data)}, " ")
}
