package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRemoveEndpointServicePermissionsResponse struct {
	Permissions    *[]EpsPermission `json:"permissions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchRemoveEndpointServicePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveEndpointServicePermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveEndpointServicePermissionsResponse", string(data)}, " ")
}
