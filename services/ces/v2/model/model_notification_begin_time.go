package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NotificationBeginTime struct {
}

func (o NotificationBeginTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationBeginTime struct{}"
	}

	return strings.Join([]string{"NotificationBeginTime", string(data)}, " ")
}
