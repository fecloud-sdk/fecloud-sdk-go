package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionResponse", string(data)}, " ")
}
