package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmResponse", string(data)}, " ")
}
