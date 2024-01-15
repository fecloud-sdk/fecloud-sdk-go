package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetRaspSwitchRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SetRaspSwitchRequestInfo `json:"body,omitempty"`
}

func (o SetRaspSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRaspSwitchRequest struct{}"
	}

	return strings.Join([]string{"SetRaspSwitchRequest", string(data)}, " ")
}
