package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowOffSiteBackupPolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowOffSiteBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOffSiteBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowOffSiteBackupPolicyRequest", string(data)}, " ")
}
