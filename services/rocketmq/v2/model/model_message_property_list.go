package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MessagePropertyList struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o MessagePropertyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessagePropertyList struct{}"
	}

	return strings.Join([]string{"MessagePropertyList", string(data)}, " ")
}
