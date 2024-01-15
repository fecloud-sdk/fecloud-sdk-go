package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AliasTarget struct {
	ResourceType *string `json:"resource_type,omitempty"`

	ResourceDomainName *string `json:"resource_domain_name,omitempty"`
}

func (o AliasTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AliasTarget struct{}"
	}

	return strings.Join([]string{"AliasTarget", string(data)}, " ")
}
