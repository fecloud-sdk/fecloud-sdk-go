package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaRemoveSecurityGroupOption struct {
	Name string `json:"name"`
}

func (o NovaRemoveSecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaRemoveSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"NovaRemoveSecurityGroupOption", string(data)}, " ")
}
