package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListL7policiesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ListenerId *string `json:"listener_id,omitempty"`

	Action *string `json:"action,omitempty"`

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	RedirectUrl *string `json:"redirect_url,omitempty"`

	Position *int32 `json:"position,omitempty"`

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DisplayAllRules *bool `json:"display_all_rules,omitempty"`
}

func (o ListL7policiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7policiesRequest struct{}"
	}

	return strings.Join([]string{"ListL7policiesRequest", string(data)}, " ")
}
