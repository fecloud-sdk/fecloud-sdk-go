package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GroupId struct {
}

func (o GroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupId struct{}"
	}

	return strings.Join([]string{"GroupId", string(data)}, " ")
}
