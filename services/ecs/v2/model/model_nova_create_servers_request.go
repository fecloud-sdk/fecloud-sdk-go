package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaCreateServersRequest Request Object
type NovaCreateServersRequest struct {

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`

	Body *NovaCreateServersRequestBody `json:"body,omitempty"`
}

func (o NovaCreateServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersRequest struct{}"
	}

	return strings.Join([]string{"NovaCreateServersRequest", string(data)}, " ")
}
