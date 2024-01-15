package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EpsAddPermissionRequest struct {
	Permission string `json:"permission"`

	Description string `json:"description"`
}

func (o EpsAddPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsAddPermissionRequest struct{}"
	}

	return strings.Join([]string{"EpsAddPermissionRequest", string(data)}, " ")
}
