package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ServerNicSecurityGroup struct {
	Id string `json:"id"`
}

func (o ServerNicSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNicSecurityGroup struct{}"
	}

	return strings.Join([]string{"ServerNicSecurityGroup", string(data)}, " ")
}
