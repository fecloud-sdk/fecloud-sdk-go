package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListL7rulesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	L7policyId string `json:"l7policy_id"`

	Id *string `json:"id,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Type *string `json:"type,omitempty"`

	CompareType *string `json:"compare_type,omitempty"`

	Invert *bool `json:"invert,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
}

func (o ListL7rulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7rulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7rulesRequest", string(data)}, " ")
}
