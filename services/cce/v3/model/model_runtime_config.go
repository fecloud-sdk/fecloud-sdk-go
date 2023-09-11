package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RuntimeConfig struct {
	LvType string `json:"lvType"`
}

func (o RuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeConfig struct{}"
	}

	return strings.Join([]string{"RuntimeConfig", string(data)}, " ")
}
