package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeEventRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ContainerName *string `json:"container_name,omitempty"`

	ContainerId *string `json:"container_id,omitempty"`

	Body *ChangeEventRequestInfo `json:"body,omitempty"`
}

func (o ChangeEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEventRequest struct{}"
	}

	return strings.Join([]string{"ChangeEventRequest", string(data)}, " ")
}
