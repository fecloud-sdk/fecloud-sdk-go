package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowServerAutoRecoveryRequest Request Object
type ShowServerAutoRecoveryRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowServerAutoRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerAutoRecoveryRequest struct{}"
	}

	return strings.Join([]string{"ShowServerAutoRecoveryRequest", string(data)}, " ")
}
