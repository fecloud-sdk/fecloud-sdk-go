package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NotificationEnabled struct {
}

func (o NotificationEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationEnabled struct{}"
	}

	return strings.Join([]string{"NotificationEnabled", string(data)}, " ")
}
