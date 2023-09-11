package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityGroupsResponse struct {
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsResponse", string(data)}, " ")
}
