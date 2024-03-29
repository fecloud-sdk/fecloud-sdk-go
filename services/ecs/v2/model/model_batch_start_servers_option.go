package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchStartServersOption struct {
	Servers []ServerId `json:"servers"`
}

func (o BatchStartServersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServersOption struct{}"
	}

	return strings.Join([]string{"BatchStartServersOption", string(data)}, " ")
}
