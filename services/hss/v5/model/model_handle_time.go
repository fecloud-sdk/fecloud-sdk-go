package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HandleTime struct {
}

func (o HandleTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleTime struct{}"
	}

	return strings.Join([]string{"HandleTime", string(data)}, " ")
}
