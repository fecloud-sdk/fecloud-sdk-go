package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NotificationEndTime struct {
}

func (o NotificationEndTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationEndTime struct{}"
	}

	return strings.Join([]string{"NotificationEndTime", string(data)}, " ")
}
