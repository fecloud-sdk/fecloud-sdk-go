package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchAddServerNicsRequest Request Object
type BatchAddServerNicsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *BatchAddServerNicsRequestBody `json:"body,omitempty"`
}

func (o BatchAddServerNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerNicsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddServerNicsRequest", string(data)}, " ")
}
