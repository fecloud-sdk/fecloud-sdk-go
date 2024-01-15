package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ComponentConfig struct {
	ComponentName string `json:"component_name"`

	Configs *[]Config `json:"configs,omitempty"`
}

func (o ComponentConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentConfig struct{}"
	}

	return strings.Join([]string{"ComponentConfig", string(data)}, " ")
}
