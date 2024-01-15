package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ComponentAmb struct {
	ComponentId *string `json:"componentId,omitempty"`

	ComponentName *string `json:"componentName,omitempty"`

	ComponentVersion *string `json:"componentVersion,omitempty"`

	ComponentDesc *string `json:"componentDesc,omitempty"`
}

func (o ComponentAmb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentAmb struct{}"
	}

	return strings.Join([]string{"ComponentAmb", string(data)}, " ")
}
