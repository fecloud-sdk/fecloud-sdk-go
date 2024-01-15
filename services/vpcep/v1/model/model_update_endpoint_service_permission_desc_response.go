package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointServicePermissionDescResponse struct {
	Permissions    *[]EpsPermission `json:"permissions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateEndpointServicePermissionDescResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServicePermissionDescResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServicePermissionDescResponse", string(data)}, " ")
}
