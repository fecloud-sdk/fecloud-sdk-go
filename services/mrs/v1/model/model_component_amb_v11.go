package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ComponentAmbV11 struct {
	ComponentName string `json:"component_name"`
}

func (o ComponentAmbV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentAmbV11 struct{}"
	}

	return strings.Join([]string{"ComponentAmbV11", string(data)}, " ")
}
