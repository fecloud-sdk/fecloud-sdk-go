package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchLimitSpeedReq struct {
	SpeedLimits []LimitSpeedReq `json:"speed_limits"`
}

func (o BatchLimitSpeedReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchLimitSpeedReq struct{}"
	}

	return strings.Join([]string{"BatchLimitSpeedReq", string(data)}, " ")
}
