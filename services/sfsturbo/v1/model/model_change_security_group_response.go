package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeSecurityGroupResponse struct {
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupResponse", string(data)}, " ")
}
