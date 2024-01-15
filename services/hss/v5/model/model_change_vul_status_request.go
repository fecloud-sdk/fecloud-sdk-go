package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeVulStatusRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChangeVulStatusRequestInfo `json:"body,omitempty"`
}

func (o ChangeVulStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeVulStatusRequest", string(data)}, " ")
}
