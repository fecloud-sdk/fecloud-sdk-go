package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttackTag struct {
}

func (o AttackTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackTag struct{}"
	}

	return strings.Join([]string{"AttackTag", string(data)}, " ")
}
