package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTime struct {
}

func (o UpdateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTime struct{}"
	}

	return strings.Join([]string{"UpdateTime", string(data)}, " ")
}
