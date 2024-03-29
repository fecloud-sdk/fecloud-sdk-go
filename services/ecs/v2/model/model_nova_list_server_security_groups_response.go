package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaListServerSecurityGroupsResponse struct {
	SecurityGroups *[]NovaSecurityGroup `json:"security_groups,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o NovaListServerSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListServerSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"NovaListServerSecurityGroupsResponse", string(data)}, " ")
}
