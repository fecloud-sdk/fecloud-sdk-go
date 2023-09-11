package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ResizeServerRequest Request Object
type ResizeServerRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ResizeServerRequestBody `json:"body,omitempty"`
}

func (o ResizeServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeServerRequest struct{}"
	}

	return strings.Join([]string{"ResizeServerRequest", string(data)}, " ")
}
