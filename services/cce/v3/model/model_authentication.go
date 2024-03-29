package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Authentication struct {
	Mode *string `json:"mode,omitempty"`

	AuthenticatingProxy *AuthenticatingProxy `json:"authenticatingProxy,omitempty"`
}

func (o Authentication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Authentication struct{}"
	}

	return strings.Join([]string{"Authentication", string(data)}, " ")
}
