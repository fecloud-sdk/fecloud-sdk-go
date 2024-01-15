package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Severity struct {
}

func (o Severity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Severity struct{}"
	}

	return strings.Join([]string{"Severity", string(data)}, " ")
}
