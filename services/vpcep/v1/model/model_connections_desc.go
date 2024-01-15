package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ConnectionsDesc struct {
	Id string `json:"id"`

	Description string `json:"description"`
}

func (o ConnectionsDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionsDesc struct{}"
	}

	return strings.Join([]string{"ConnectionsDesc", string(data)}, " ")
}
