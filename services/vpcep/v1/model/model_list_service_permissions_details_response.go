package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListServicePermissionsDetailsResponse struct {
	Permissions *[]PermissionObject `json:"permissions,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServicePermissionsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePermissionsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServicePermissionsDetailsResponse", string(data)}, " ")
}
