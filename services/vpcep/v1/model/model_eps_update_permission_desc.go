package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EpsUpdatePermissionDesc struct {
	Description string `json:"description"`
}

func (o EpsUpdatePermissionDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsUpdatePermissionDesc struct{}"
	}

	return strings.Join([]string{"EpsUpdatePermissionDesc", string(data)}, " ")
}
