package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AgentId struct {
}

func (o AgentId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentId struct{}"
	}

	return strings.Join([]string{"AgentId", string(data)}, " ")
}
