package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PolicyGroupName struct {
}

func (o PolicyGroupName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupName struct{}"
	}

	return strings.Join([]string{"PolicyGroupName", string(data)}, " ")
}
