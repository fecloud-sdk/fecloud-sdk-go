package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHostsGroupRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	GroupId string `json:"group_id"`
}

func (o DeleteHostsGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostsGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostsGroupRequest", string(data)}, " ")
}
