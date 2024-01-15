package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SupportVersion struct {
}

func (o SupportVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportVersion struct{}"
	}

	return strings.Join([]string{"SupportVersion", string(data)}, " ")
}
