package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttackPhase struct {
}

func (o AttackPhase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackPhase struct{}"
	}

	return strings.Join([]string{"AttackPhase", string(data)}, " ")
}
