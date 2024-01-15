package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaServerSecurityGroup struct {
	Name *string `json:"name,omitempty"`
}

func (o NovaServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"NovaServerSecurityGroup", string(data)}, " ")
}
