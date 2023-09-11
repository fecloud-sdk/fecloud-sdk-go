package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ResizePostPaidServerRequest Request Object
type ResizePostPaidServerRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ResizePostPaidServerRequestBody `json:"body,omitempty"`
}

func (o ResizePostPaidServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerRequest struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerRequest", string(data)}, " ")
}