package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Notification struct {
	Type string `json:"type"`

	NotificationList []string `json:"notification_list"`
}

func (o Notification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notification struct{}"
	}

	return strings.Join([]string{"Notification", string(data)}, " ")
}
