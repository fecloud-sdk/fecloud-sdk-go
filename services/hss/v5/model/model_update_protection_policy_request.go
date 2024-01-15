package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateProtectionPolicyRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateProtectionPolicyInfoRequestInfo `json:"body,omitempty"`
}

func (o UpdateProtectionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateProtectionPolicyRequest", string(data)}, " ")
}
