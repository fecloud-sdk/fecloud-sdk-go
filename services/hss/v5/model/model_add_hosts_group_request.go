package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddHostsGroupRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *AddHostsGroupRequestInfo `json:"body,omitempty"`
}

func (o AddHostsGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHostsGroupRequest struct{}"
	}

	return strings.Join([]string{"AddHostsGroupRequest", string(data)}, " ")
}
