package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaAddSecurityGroupOption struct {
	Name string `json:"name"`
}

func (o NovaAddSecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAddSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"NovaAddSecurityGroupOption", string(data)}, " ")
}
