package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AlarmName struct {
}

func (o AlarmName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmName struct{}"
	}

	return strings.Join([]string{"AlarmName", string(data)}, " ")
}
