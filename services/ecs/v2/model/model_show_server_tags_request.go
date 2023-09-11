package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowServerTagsRequest Request Object
type ShowServerTagsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowServerTagsRequest", string(data)}, " ")
}