package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AttachServerVolumeRequest Request Object
type AttachServerVolumeRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *AttachServerVolumeRequestBody `json:"body,omitempty"`
}

func (o AttachServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeRequest", string(data)}, " ")
}
