package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AlarmDescription struct {
}

func (o AlarmDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDescription struct{}"
	}

	return strings.Join([]string{"AlarmDescription", string(data)}, " ")
}
