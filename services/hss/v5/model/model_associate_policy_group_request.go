package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociatePolicyGroupRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *AssociatePolicyGroupRequestInfo `json:"body,omitempty"`
}

func (o AssociatePolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"AssociatePolicyGroupRequest", string(data)}, " ")
}
