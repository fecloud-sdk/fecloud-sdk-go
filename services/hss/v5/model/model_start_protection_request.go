package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StartProtectionRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ProtectionInfoRequestInfo `json:"body,omitempty"`
}

func (o StartProtectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartProtectionRequest struct{}"
	}

	return strings.Join([]string{"StartProtectionRequest", string(data)}, " ")
}
