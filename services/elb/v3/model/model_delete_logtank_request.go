package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteLogtankRequest Request Object
type DeleteLogtankRequest struct {

	// 云日志ID。
	LogtankId string `json:"logtank_id"`
}

func (o DeleteLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogtankRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogtankRequest", string(data)}, " ")
}
