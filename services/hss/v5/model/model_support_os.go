package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SupportOs struct {
}

func (o SupportOs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportOs struct{}"
	}

	return strings.Join([]string{"SupportOs", string(data)}, " ")
}
