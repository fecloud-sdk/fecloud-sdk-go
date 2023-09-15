package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowLogtankRequest Request Object
type ShowLogtankRequest struct {

	// 云日志ID。
	LogtankId string `json:"logtank_id"`
}

func (o ShowLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankRequest struct{}"
	}

	return strings.Join([]string{"ShowLogtankRequest", string(data)}, " ")
}
