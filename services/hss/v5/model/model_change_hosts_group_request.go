package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeHostsGroupRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChangeHostsGroupRequestInfo `json:"body,omitempty"`
}

func (o ChangeHostsGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostsGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeHostsGroupRequest", string(data)}, " ")
}
