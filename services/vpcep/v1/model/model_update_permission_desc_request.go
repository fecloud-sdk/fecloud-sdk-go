package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePermissionDescRequest struct {
	Permission *EpsUpdatePermissionDesc `json:"permission"`
}

func (o UpdatePermissionDescRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionDescRequest struct{}"
	}

	return strings.Join([]string{"UpdatePermissionDescRequest", string(data)}, " ")
}
