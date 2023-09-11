package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaCreateKeypairRequest Request Object
type NovaCreateKeypairRequest struct {

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`

	Body *NovaCreateKeypairRequestBody `json:"body,omitempty"`
}

func (o NovaCreateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairRequest", string(data)}, " ")
}
