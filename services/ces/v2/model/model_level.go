package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Level struct {
}

func (o Level) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Level struct{}"
	}

	return strings.Join([]string{"Level", string(data)}, " ")
}
