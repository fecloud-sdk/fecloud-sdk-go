package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PolicyGroupId struct {
}

func (o PolicyGroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupId struct{}"
	}

	return strings.Join([]string{"PolicyGroupId", string(data)}, " ")
}
