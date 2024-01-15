package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Subscription struct {
	Topic *string `json:"topic,omitempty"`

	Type *string `json:"type,omitempty"`

	Expression *string `json:"expression,omitempty"`
}

func (o Subscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subscription struct{}"
	}

	return strings.Join([]string{"Subscription", string(data)}, " ")
}
