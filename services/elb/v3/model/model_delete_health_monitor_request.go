package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteHealthMonitorRequest Request Object
type DeleteHealthMonitorRequest struct {

	// 健康检查ID。
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorRequest", string(data)}, " ")
}
