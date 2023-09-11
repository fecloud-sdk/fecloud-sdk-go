package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteServerPasswordRequest Request Object
type DeleteServerPasswordRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o DeleteServerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerPasswordRequest", string(data)}, " ")
}
