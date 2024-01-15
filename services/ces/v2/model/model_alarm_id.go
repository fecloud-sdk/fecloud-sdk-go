package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AlarmId struct {
}

func (o AlarmId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmId struct{}"
	}

	return strings.Join([]string{"AlarmId", string(data)}, " ")
}
