package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProxyReadonlyInstances struct {
	Id string `json:"id"`

	Weight int32 `json:"weight"`
}

func (o ProxyReadonlyInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyReadonlyInstances struct{}"
	}

	return strings.Join([]string{"ProxyReadonlyInstances", string(data)}, " ")
}
