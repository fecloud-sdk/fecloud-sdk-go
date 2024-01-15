package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GroupName struct {
}

func (o GroupName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupName struct{}"
	}

	return strings.Join([]string{"GroupName", string(data)}, " ")
}
