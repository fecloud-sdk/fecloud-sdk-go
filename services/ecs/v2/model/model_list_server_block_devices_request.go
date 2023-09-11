package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListServerBlockDevicesRequest Request Object
type ListServerBlockDevicesRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ListServerBlockDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerBlockDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListServerBlockDevicesRequest", string(data)}, " ")
}
