package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StopProtectionRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CloseProtectionInfoRequestInfo `json:"body,omitempty"`
}

func (o StopProtectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopProtectionRequest struct{}"
	}

	return strings.Join([]string{"StopProtectionRequest", string(data)}, " ")
}
