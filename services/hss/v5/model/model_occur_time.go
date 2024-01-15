package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OccurTime struct {
}

func (o OccurTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OccurTime struct{}"
	}

	return strings.Join([]string{"OccurTime", string(data)}, " ")
}
