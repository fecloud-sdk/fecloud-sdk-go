package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ResetServerPasswordRequest Request Object
type ResetServerPasswordRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ResetServerPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetServerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordRequest", string(data)}, " ")
}
