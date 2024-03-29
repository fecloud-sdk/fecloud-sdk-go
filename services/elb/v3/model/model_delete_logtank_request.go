package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLogtankRequest struct {
	LogtankId string `json:"logtank_id"`
}

func (o DeleteLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogtankRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogtankRequest", string(data)}, " ")
}
