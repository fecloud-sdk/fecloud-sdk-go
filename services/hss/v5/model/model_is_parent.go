package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type IsParent struct {
}

func (o IsParent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsParent struct{}"
	}

	return strings.Join([]string{"IsParent", string(data)}, " ")
}
