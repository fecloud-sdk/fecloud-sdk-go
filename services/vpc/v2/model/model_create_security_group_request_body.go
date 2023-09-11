package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRequestBody struct {
	SecurityGroup *CreateSecurityGroupOption `json:"security_group"`
}

func (o CreateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRequestBody", string(data)}, " ")
}
