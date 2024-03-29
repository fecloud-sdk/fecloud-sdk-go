package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceRestartRequsetBody struct {
	Restart *interface{} `json:"restart"`
}

func (o InstanceRestartRequsetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRestartRequsetBody struct{}"
	}

	return strings.Join([]string{"InstanceRestartRequsetBody", string(data)}, " ")
}
