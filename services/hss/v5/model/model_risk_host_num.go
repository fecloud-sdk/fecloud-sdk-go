package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RiskHostNum struct {
}

func (o RiskHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHostNum struct{}"
	}

	return strings.Join([]string{"RiskHostNum", string(data)}, " ")
}
