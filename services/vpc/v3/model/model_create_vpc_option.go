package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcOption struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Cidr string `json:"cidr"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcOption struct{}"
	}

	return strings.Join([]string{"CreateVpcOption", string(data)}, " ")
}
