package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddEndpointServicePermissionsResponse struct {
	Permissions    *[]EpsPermission `json:"permissions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchAddEndpointServicePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddEndpointServicePermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddEndpointServicePermissionsResponse", string(data)}, " ")
}
