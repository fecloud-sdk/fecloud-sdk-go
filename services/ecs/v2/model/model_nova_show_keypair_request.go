package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaShowKeypairRequest Request Object
type NovaShowKeypairRequest struct {

	// 密钥名称信息。
	KeypairName string `json:"keypair_name"`

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`
}

func (o NovaShowKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaShowKeypairRequest", string(data)}, " ")
}
